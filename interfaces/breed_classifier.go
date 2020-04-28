package interfaces

type BreedClassifier interface {
	GetBreeds(image string) ([]string, error)
}
