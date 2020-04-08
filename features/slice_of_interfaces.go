package features

import (
	"github.com/Kotyarich/find-your-pet/models"
)

func ConvertSlicesElementsToInterface(slice []models.Lost) []interface{} {
	result := make([]interface{}, len(slice))
	for i := range slice {
		result[i] = slice[i]
	}
	return result
}

func ConvertInterfaceElementsToLost(slice []interface{}) []models.Lost {
	result := make([]models.Lost, len(slice))
	for i := range slice {
		element := slice[i].(models.Lost)
		result[i] = element
	}
	return result
}
