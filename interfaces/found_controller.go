package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type FoundController interface {
	GetById(ctx context.Context, id int) (*models.Found, error)
	Add(ctx context.Context, params *models.Found) (int, error)
	Search(ctx context.Context, params *models.Found) ([]models.Found, error)

	GetItemsPerPageCount() int
}
