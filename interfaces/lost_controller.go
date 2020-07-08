package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostController interface {
	GetById(ctx context.Context, id int) (*models.Lost, error)
	Add(ctx context.Context, params *models.Lost) (int, error)
	GetAll() ([]models.Lost, error)
	Search(ctx context.Context, params *models.Lost, query string, page int) ([]models.Lost, bool, error)

	// Returns picture_id (int) and error. If picture is null, returns 0
	RemoveById(ctx context.Context, id int) (int, error)

	// Finds similar pets. Returns their ids and error
	GetSimilars(found *models.Found) ([]models.Similar, error)

	GetPageCapacity() int
}
