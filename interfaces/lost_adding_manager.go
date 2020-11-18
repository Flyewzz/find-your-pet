package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostAddingManager interface {
	Add(ctx context.Context, params *models.Lost,
		lostIdCh chan<- int,
		fileCh <-chan *models.File, errCh chan<- error)
	Remove(id int) error
}
