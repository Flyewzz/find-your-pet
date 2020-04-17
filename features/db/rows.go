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

func ConvertRowsToFound(rows *sql.Rows) ([]models.Found, error) {
	var err error
	var founds []models.Found
	for rows.Next() {
		var found models.Found
		err = rows.Scan(&found.Id, &found.TypeId, &found.AuthorId,
			&found.Sex, &found.Breed, &found.Description,
			&found.StatusId, &found.Date,
			&found.Latitude, &found.Longitude, &found.PictureId)
		if err != nil {
			continue // !
		}
		founds = append(founds, found)
	}
	return founds, err
}
