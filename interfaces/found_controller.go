package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type FoundController interface {
	GetById(ctx context.Context, id int) (*models.Found, error)
	Add(ctx context.Context, params *models.Found) (int, error)
	Search(ctx context.Context, params *models.Found, query string, page int) ([]models.Found, bool, error)
	// Returns picture_id (int) and error. If picture is null, returns 0
	RemoveById(ctx context.Context, id int) (int, error)

	GetPageCapacity() int
}
