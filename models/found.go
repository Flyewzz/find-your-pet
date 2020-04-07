package models

type Found struct {
	Id          int    `json:"id"`
	TypeId      int    `json:"type_id"`
	AuthorId    int    `json:"author_id"`
	Sex         string `json:"sex"`
	Breed       string `json:"breed"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Date        string `json:"date"`
	Place       string `json:"place"`
}
