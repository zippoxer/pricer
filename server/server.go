package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	pb "github.com/zippoxer/pricer/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	dbPath = flag.String("db", "pricer.db", "Path to the local database.")
)

type server struct {
	pb.UnimplementedPricerServer
	store *Store
}

func (s *server) Price(ctx context.Context, in *pb.PriceRequest) (*pb.PriceReply, error) {
	price, err := s.store.Get(in.GetToken())
	if err != nil {
		return nil, err
	}
	return &pb.PriceReply{Price: price}, nil
}

func main() {
	flag.Parse()

	// Open the Store.
	store, err := OpenStore(*dbPath)
	if err != nil {
		log.Fatalf("bolt.Open: %v", err)
	}
	defer store.Close()

	// Start the Fetcher.
	ctx, cancelFunc := context.WithCancel(context.Background())
	_ = cancelFunc // Use to gracefully stop the Fetcher.
	fetcher := Fetcher{store: store, client: &http.Client{}}
	go func() {
		err := fetcher.Run(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Start the PricerServer.
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPricerServer(s, &server{store: store})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
