package classifier

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

type BreedClassifier struct {
}

func NewBreedClassifier() *BreedClassifier {
	return &BreedClassifier{}
}

func (bc *BreedClassifier) GetBreeds(image string) ([]string, error) {
	// Set up a connection to the server.
	log.Println("Attempting to connect to the breed classifier server...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Printf("Timeout error: %v", err)
		}
		log.Printf("did not connect: %v", err)
		return nil, err
	}
	log.Println("*** Breed client is OK ***")
	defer conn.Close()
	client := pb.NewBreedClassifierServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	breed, err := client.RecognizeBreed(ctx, &pb.Image{Path: image})
	if err != nil {
		log.Printf("could not get a breed: %v", err)
		return nil, err
	}
	log.Printf("Breed: %s\n", breed.GetName())

	return []string{breed.GetName()}, nil
}
