package db

import (
	"database/sql"

	"github.com/Kotyarich/find-your-pet/models"
)

func ConvertRowsToLost(rows *sql.Rows) ([]models.Lost, error) {
	var err error
	var losts []models.Lost
	for rows.Next() {
		var lost models.Lost
		err = rows.Scan(&lost.Id, &lost.TypeId, &lost.AuthorId,
			&lost.Sex, &lost.Breed, &lost.Description,
			&lost.StatusId, &lost.Date,
			&lost.Latitude, &lost.Longitude, &lost.PictureId)
		if err != nil {
			continue // !
		}
		losts = append(losts, lost)
	}
	return losts, err
}
