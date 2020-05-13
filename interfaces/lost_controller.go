package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostController interface {
	GetById(ctx context.Context, id int) (*models.Lost, error)
	Add(ctx context.Context, params *models.Lost) (int, error)
	Search(ctx context.Context, params *models.Lost, query string, page int) ([]models.Lost, bool, error)

	GetPageCapacity() int
}
