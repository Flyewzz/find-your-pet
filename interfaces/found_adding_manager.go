package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type FoundAddingManager interface {
	Add(ctx context.Context, params *models.Found,
		foundIdCh chan<- int,
		fileCh <-chan *models.File, errCh chan<- error)
	Remove(id int) error
}
