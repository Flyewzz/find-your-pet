package models

type MaterialData struct {
	MaterialId int    `json:"material_id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	TypeName   string `json:"type"`
	UserEmail  string `json:"user_email"`
}
