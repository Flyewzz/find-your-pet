package models

type Lost struct {
	Id          int     `json:"id"`
	TypeId      int     `json:"type_id"`
	AuthorId    int     `json:"vk_id"`
	Sex         string  `json:"sex"`
	Breed       string  `json:"breed"`
	Description string  `json:"description"`
	StatusId    int     `json:"status_id"`
	Date        string  `json:"date"`
	Location    string  `json:"location,omitempty"`
	PictureId   int     `json:"picture_id,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}
