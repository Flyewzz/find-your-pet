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
	queryLost    string
	queryFound   string
}

func NewProfileControllerPg(pages int, db *sql.DB, queryLost, queryFound string) *ProfileControllerPg {
	return &ProfileControllerPg{
		itemsPerPage: pages,
		db:           db,
		queryLost:    queryLost,
		queryFound:   queryFound,
	}
}

func (pc *ProfileControllerPg) GetLost(ctx context.Context, userId int) ([]models.Lost, error) {
	closeId := ctx.Value("close_id")
	rows, err := pc.db.Query(
		pc.queryLost+
			"WHERE vk_id = $1 AND status_id != $2 ORDER BY date DESC", userId, closeId)
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

func (pc *ProfileControllerPg) GetFound(ctx context.Context, userId int) ([]models.Found, error) {
	closeId := ctx.Value("close_id")
	rows, err := pc.db.Query(
		pc.queryFound+
			"WHERE vk_id = $1 AND status_id != $2 ORDER BY date DESC", userId, closeId)
	if err != nil {
		return nil, err
	}
	found, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return found, err
}

func (pc *ProfileControllerPg) SetFoundOpening(
	ctx context.Context, foundId int, opened bool) error {
	var err error = nil
	query := "UPDATE found SET status_id = $1 " +
		"WHERE id = $2"
	var result sql.Result
	params := ctx.Value("params").(features.StatusIdParams)
	if opened {
		// To open the announcement
		result, err = pc.db.Exec(query, params.OpenId, foundId)
	} else {
		// To close the announcement
		result, err = pc.db.Exec(query, params.CloseId, foundId)
	}
	if err == nil {
		countAffected, errAff := result.RowsAffected()
		if errAff != nil {
			return errAff
		}
		// If id doesn't exist
		if countAffected == 0 {
			return errs.TheFoundNotFound
		}
	}
	return err
}

func (pc *ProfileControllerPg) GetItemsPerPageCount() int {
	return pc.itemsPerPage
}
