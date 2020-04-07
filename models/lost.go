package models

type Lost struct {
	Id          int    `json:"id"`
	TypeId      int    `json:"type_id"`
	AuthorId    int    `json:"author_id"`
	Sex         string `json:"sex"`
	Breed       string `json:"breed"`
	Description string `json:"description"`
	StatusId    int    `json:"status_id"`
	Date        string `json:"date"`
	Place       string `json:"place"`
}
