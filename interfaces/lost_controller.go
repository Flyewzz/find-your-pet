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

/*

	Id          int    `json:"id"`
	TypeId      int    `json:"type_id"`
	AuthorId    int    `json:"author_id"`
	Sex         string `json:"sex"`
	Breed       string `json:"breed"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Date        string `json:"date"`
	Place       string `json:"place"`

*/
