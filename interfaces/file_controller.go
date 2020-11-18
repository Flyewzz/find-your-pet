package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type FileController interface {
	GetById(id int) (*models.File, error)
	AddToLost(ctx context.Context, file *models.File, lostId int) (int, error)
	AddToFound(ctx context.Context, file *models.File, foundId int) (int, error)
	Remove(ctx context.Context, fileId int) error
}
