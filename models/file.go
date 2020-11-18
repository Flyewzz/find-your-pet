package models

type File struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"-"`
	IsZoosearch bool   `json:"is_zoosearch"`
}
