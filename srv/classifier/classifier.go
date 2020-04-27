package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Kotyarich/find-your-pet/pkg/classifier"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	log.Println("Attempting to connect to the breed classifier server...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Fatalf("Timeout error: %v", err)
		}
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("*** Breed client is OK ***")
	defer conn.Close()
	client := pb.NewBreedClassifierServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	breed, err := client.RecognizeBreed(ctx, &pb.Image{Path: "/Users/alpha/Downloads/11093267_2.jpg"})
	if err != nil {
		log.Fatalf("could not get a breed: %v", err)
	}
	log.Printf("Breed: %s\n", breed.GetName())
}
