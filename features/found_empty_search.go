package features

import "github.com/Kotyarich/find-your-pet/models"

func CheckEmptyFound(found *models.Found, query string) bool {
	if found.Breed == "" &&
		found.Description == "" &&
		found.Latitude == 0 && found.Longitude == 0 &&
		found.Sex == "" && found.StatusId == 0 &&
		found.TypeId == 0 &&
		query == "" {
		return true
	}
	return false
}
