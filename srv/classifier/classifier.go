package classifier

import (
	"context"
	"log"
	"strings"
	"time"

	pb "github.com/Kotyarich/find-your-pet/pkg/classifier"
	"google.golang.org/grpc"
)

type BreedClassifier struct {
	address          string
	connTimeout      int
	recognizeTimeout int
}

func NewBreedClassifier(address string, connTimeout, recognizeTimeout int) *BreedClassifier {
	return &BreedClassifier{
		address:          address,
		connTimeout:      connTimeout,
		recognizeTimeout: recognizeTimeout,
	}
}

func (bc *BreedClassifier) GetBreeds(image string) ([]string, error) {
	// Set up a connection to the server.
	log.Println("Attempting to connect to the breed classifier server...")
	conn, err := grpc.Dial(bc.address, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(time.Duration(bc.connTimeout)*time.Second))
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Printf("Timeout error: %v", err)
		} else {
			log.Printf("did not connect: %v", err)
		}
		return nil, err
	}
	log.Println("*** A breed client is OK ***")
	defer conn.Close()
	client := pb.NewBreedClassifierServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(bc.recognizeTimeout)*time.Second)
	defer cancel()
	breed, err := client.RecognizeBreed(ctx, &pb.Image{Path: image})
	if err != nil {
		log.Printf("Could not get a breed: %v", err)
		return nil, err
	}
	breeds := breed.GetName()
	log.Printf("Breeds: %s\n", strings.Join(breeds, " "))

	return breeds, nil
}
