package http_breed_classifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BreedClassifier struct {
	url string
}

func NewBreedClassifier(url string) *BreedClassifier {
	return &BreedClassifier{
		url: url,
	}
}

func (bc *BreedClassifier) GetBreeds(image string) ([]string, error) {
	// Set up a connection to the server.
	log.Println("Attempting to connect to the breed classifier server...")
	fmt.Println("URL:>", bc.url)

	type picture struct {
		Picture string `json:"picture"`
	}
	pictureStruct := picture{
		Picture: image,
	}
	// var jsonStr = []byte(fmt.Sprintf(`{"picture":"%s"}`, )) /
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(pictureStruct)
	res, err := http.Post(bc.url, "application/json; charset=utf-8", buf)
	if err != nil {
		return nil, err
	}
	var breeds []string
	d := json.NewDecoder(res.Body)
	err = d.Decode(&breeds)
	if err != nil {
		return nil, err
	}
	return breeds, err
}
