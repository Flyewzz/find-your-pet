package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostController interface {
	GetById(id int) (*models.Lost, error)
	Add(ctx context.Context, params *models.Lost) (int, error)
	Search(params *models.Lost) ([]models.Lost, error)

	GetItemsPerPageCount() int
}
