package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type ProfileController interface {
	GetLost(ctx context.Context, userId int) ([]models.Lost, error)
	SetLostOpening(ctx context.Context, lostId int, statusId int) error

	GetFound(ctx context.Context, userId int) ([]models.Found, error)
	SetFoundOpening(ctx context.Context, foundId int, statusId int) error

	GetItemsPerPageCount() int
}
