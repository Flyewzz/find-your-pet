package interfaces

import (
	"github.com/Kotyarich/find-your-pet/models"
)

type ProfileController interface {
	GetLost(userId int) ([]models.Lost, error)
	SetLostOpening(lostId int, opened bool) error

	GetItemsPerPageCount() int
}
