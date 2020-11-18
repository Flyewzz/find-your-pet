package mocks

import (
	"context"

	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/models"
)

type MockLostAddingManager struct {
	LostController interfaces.LostController
}

func (lam MockLostAddingManager) Add(ctx context.Context, params *models.Lost,
	lostIdCh chan<- int,
	fileCh <-chan *models.File, errCh chan<- error) {

}

func (lam MockLostAddingManager) Remove(id int) error {
	_, err := lam.LostController.RemoveById(context.Background(), id)
	return err
}
