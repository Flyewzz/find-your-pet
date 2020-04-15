package pg

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/features/db"
	"github.com/Kotyarich/find-your-pet/models"
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

func (pc *ProfileControllerPg) GetLost(ctx context.Context, userId int) ([]models.Lost, error) {
	closeId := ctx.Value("close_id")
	rows, err := pc.db.Query(
		"SELECT id, type_id, vk_id, sex, "+
			"breed, description, status_id, "+
			"date, st_x(location) as latitude, "+
			"st_y(location) as longitude, picture_id FROM lost "+
			"WHERE vk_id = $1 AND status_id != $2", userId, closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (pc *ProfileControllerPg) SetLostOpening(
	ctx context.Context, lostId int, opened bool) error {
	var err error = nil
	query := "UPDATE lost SET status_id = $1 " +
		"WHERE id = $2"
	var result sql.Result
	params := ctx.Value("params").(features.StatusIdParams)
	if opened {
		// To open the announcement
		result, err = pc.db.Exec(query, params.OpenId, lostId)
	} else {
		// To close the announcement
		result, err = pc.db.Exec(query, params.CloseId, lostId)
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
