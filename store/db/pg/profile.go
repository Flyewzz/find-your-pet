package pg

import (
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features/db"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/spf13/viper"
)

type ProfileControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewProfileControllerPg(pages int, db *sql.DB) *ProfileControllerPg {
	return &ProfileControllerPg{
		itemsPerPage: pages,
		db:           db,
	}
}

func (pc *ProfileControllerPg) GetLost(userId int) ([]models.Lost, error) {
	rows, err := pc.db.Query(
		"SELECT id, type_id, vk_id, sex, "+
			"breed, description, status_id, "+
			"date, st_x(location) as latitude, "+
			"st_y(location) as longitude, picture_id FROM lost "+
			"WHERE vk_id = $1", userId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (pc *ProfileControllerPg) SetLostOpening(lostId int, opened bool) error {
	var err error = nil
	query := "UPDATE lost SET status_id = $1 " +
		"WHERE id = $2"
	var result sql.Result
	if opened {
		// To open the announcement
		result, err = pc.db.Exec(query, viper.GetInt("lost.open_id"), lostId)
	} else {
		// To close the announcement
		result, err = pc.db.Exec(query, viper.GetInt("lost.close_id"), lostId)
	}
	if err == nil {
		countAffected, errAff := result.RowsAffected()
		if errAff != nil {
			return errAff
		}
		// If id doesn't exist
		if countAffected == 0 {
			return errs.LostNotFound
		}
	}
	return err
}

func (pc *ProfileControllerPg) GetItemsPerPageCount() int {
	return pc.itemsPerPage
}
