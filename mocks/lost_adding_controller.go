package mocks

import (
	"context"
	"errors"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostAddingManager struct {
}

func (lam *LostAddingManager) Add(ctx context.Context, params *models.Lost,
	lostIdCh chan<- int,
	fileCh <-chan *models.File, errCh chan<- error) {
	//!PASS
}

func (lam *LostAddingManager) Remove(id int) error {
	return errors.New("Directory is unavailable")
}
