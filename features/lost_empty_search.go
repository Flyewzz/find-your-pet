package features

import "github.com/Kotyarich/find-your-pet/models"

func CheckEmptyLost(lost *models.Lost, query string) bool {
	if lost.Breed == "" &&
		lost.Description == "" &&
		lost.Latitude == 0 && lost.Longitude == 0 &&
		lost.Sex == "" && lost.StatusId == 0 &&
		lost.TypeId == 0 &&
		query == "" {
		return true
	}
	return false
}
