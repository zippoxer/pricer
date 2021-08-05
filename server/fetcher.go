package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	cmcEndpoint = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?start=1&limit=3&convert=USD"
)

// Fetcher fetches prices from CoinMarketCap's API
// every minute and writes them to Store.
type Fetcher struct {
	store  *Store
	client *http.Client
}

func (f *Fetcher) Run(ctx context.Context) error {
	// Ticker is preferrable to Sleep because it keeps
	// counting even while we're making requests. 👍
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		prices, err := f.getPrices()
		if err != nil {
			return err
		}
		for token, price := range prices {
			err = f.store.Put(token, price)
			if err != nil {
				return err
			}
		}
		log.Printf("Fetcher: received prices (%v)", prices)

		select {
		case <-ticker.C: // ⌛ A minute has gone by.
		case <-ctx.Done(): // ❌ Context has been canceled.
			return nil
		}
	}
}

func (f *Fetcher) getPrices() (map[string]float64, error) {
	// Make the HTTP request.
	req, err := http.NewRequest("GET", cmcEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CMC_PRO_API_KEY", "08b53f4d-0a5d-452e-b67f-3af8af0d035c")
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Fetcher.getPrices: invalid status %d", resp.StatusCode)
	}

	// Parse the JSON response.
	var result cmcResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	if result.Status.ErrorCode != 0 {
		return nil, fmt.Errorf("Fetcher.getPrices: %s", result.Status.ErrorMessage)
	}

	// Extract the prices.
	prices := make(map[string]float64)
	for _, data := range result.Data {
		prices[data.Symbol] = data.Quote.USD.Price
	}

	return prices, nil
}

type cmcResponse struct {
	Status struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	}

	Data []struct {
		Symbol string
		Quote  struct {
			USD struct {
				Price float64
			}
		}
	}
}
