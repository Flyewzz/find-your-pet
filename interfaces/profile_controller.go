package interfaces

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type ProfileController interface {
	GetLost(ctx context.Context, userId int) ([]models.Lost, error)
	SetLostOpening(ctx context.Context, lostId int, opened bool) error

	GetFound(ctx context.Context, userId int) ([]models.Found, error)
	SetFoundOpening(ctx context.Context, foundId int, opened bool) error

	GetItemsPerPageCount() int
}
