package interfaces

import (
	"github.com/Kotyarich/find-your-pet/models"
)

type LostController interface {
	GetById(id int) (*models.Lost, error)
	Add(params *models.Lost) (int, error)
	Search(params *models.Lost) ([]models.Lost, error)

	GetItemsPerPageCount() int
}
