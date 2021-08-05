package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	pb "github.com/zippoxer/pricer/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"

	// concurrency limits how many requests can run at the same time.
	concurrency = 3
)

var (
	tokens   = flag.String("tokens", "BTC,ETH,USDT", "Comma-separated list of tokens to check.")
	interval = flag.Int("interval", 3, "Price check interval in seconds.")
)

type PriceResponse struct {
	Token string
	Price float64
}

type Client struct {
	pricer      pb.PricerClient
	retryTicker *time.Ticker
}

func (c *Client) Run(ctx context.Context) error {
	// We spawn a controlled amount of goroutines (per `concurrency`) to make requests, instead of a goroutine per token.
	// This assures we're not reaching the point of diminishing returns if there are many tokens to request.
	var wg sync.WaitGroup
	requests := make(chan string)
	responses := make(chan PriceResponse)

	// Process price requests.
	c.retryTicker = time.NewTicker(3 * time.Second)
	defer c.retryTicker.Stop()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go c.doRequests(ctx, wg, requests, responses)
	}

	// Log price responses.
	go c.logPrices(ctx, responses)

	// Push price requests to the workers at the given interval.
	c.pushRequests(ctx, requests)

	// Wait for workers to finish.
	wg.Wait()
	close(responses)

	return nil
}

func (c *Client) pushRequests(ctx context.Context, priceRequests chan<- string) error {
	tokenSlice := strings.Split(*tokens, ",")
	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	defer ticker.Stop()
	for {
		for _, token := range tokenSlice {
			priceRequests <- token
		}

		select {
		case <-ticker.C: // ⌛ Tick.
		case <-ctx.Done(): // ❌ Context has been canceled.
			close(priceRequests)
			return nil
		}
	}
}

func (c *Client) doRequests(ctx context.Context, wg sync.WaitGroup, priceRequests <-chan string, priceResponses chan<- PriceResponse) {
	defer wg.Done()
	for token := range priceRequests {
		price, err := c.getPrice(ctx, c.pricer, token)
		if err != nil {
			log.Fatalf("getPrice(%s): %v", token, err)
		}
		priceResponses <- PriceResponse{
			Token: token,
			Price: price,
		}

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (c *Client) getPrice(ctx context.Context, pricer pb.PricerClient, token string) (float64, error) {
	var price float64
	for {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		r, err := pricer.Price(ctx, &pb.PriceRequest{Token: token})
		if err != nil {
			// TODO: check if err is a network error and only keep retrying if it is.
			log.Printf("getPrice: %v (retrying...)", err)
			<-c.retryTicker.C
			continue
		}
		price = r.Price
		break
	}
	return price, nil
}

func (c *Client) logPrices(ctx context.Context, priceResponses <-chan PriceResponse) {
	prices := make(map[string]float64)
	for resp := range priceResponses {
		changeStr := ""
		oldPrice, ok := prices[resp.Token]
		if ok {
			changeStr = fmt.Sprintf(", %.2f, %.2f%%", oldPrice, (resp.Price-oldPrice)/oldPrice*100)
		}
		log.Printf("%s: %.2f%s", resp.Token, resp.Price, changeStr)
		prices[resp.Token] = resp.Price

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Start the Client.
	ctx, cancelFunc := context.WithCancel(context.Background())
	_ = cancelFunc // Use to gracefully stop the Client.

	client := &Client{
		pricer: pb.NewPricerClient(conn),
	}
	err = client.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Stopped gracefully.")
}
