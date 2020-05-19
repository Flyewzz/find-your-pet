package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostController interface {
	GetById(ctx context.Context, id int) (*models.Lost, error)
	Add(ctx context.Context, params *models.Lost) (int, error)
	Search(ctx context.Context, params *models.Lost, query string, page int) ([]models.Lost, bool, error)
	// Returns picture_id (int) and error. If picture is null, returns 0
	RemoveById(ctx context.Context, id int) (int, error)

	GetPageCapacity() int
}
