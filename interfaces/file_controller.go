package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type LostFileController interface {
	GetById(id int) (*models.File, error)
	Add(ctx context.Context, file *models.File, lostId int) (int, error)
}
