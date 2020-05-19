package pg

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
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
	rows, err := pc.db.Query(
		pc.queryLost+
			"WHERE vk_id = $1 ORDER BY date DESC", userId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (pc *ProfileControllerPg) SetLostOpening(
	ctx context.Context, lostId int, statusId int) error {
	var err error = nil
	query := "UPDATE lost SET status_id = $1 " +
		"WHERE id = $2"
	var result sql.Result
	result, err = pc.db.Exec(query, statusId, lostId)
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
	rows, err := pc.db.Query(
		pc.queryFound+
			"WHERE vk_id = $1 ORDER BY date DESC", userId)
	if err != nil {
		return nil, err
	}
	found, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return found, err
}

func (pc *ProfileControllerPg) SetFoundOpening(
	ctx context.Context, foundId int, statusId int) error {
	var err error = nil
	query := "UPDATE found SET status_id = $1 " +
		"WHERE id = $2"
	var result sql.Result
	result, err = pc.db.Exec(query, statusId, foundId)
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
