package pg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/features/db"
	"github.com/Kotyarich/find-your-pet/features/search"
	"github.com/Kotyarich/find-your-pet/models"
	set "github.com/deckarep/golang-set"
)

type LostControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewLostControllerPg(itemsPerPage int, db *sql.DB) *LostControllerPg {
	return &LostControllerPg{
		itemsPerPage: itemsPerPage,
		db:           db,
	}
}

func (lc *LostControllerPg) GetById(ctx context.Context, id int) (*models.Lost, error) {
	closeId := ctx.Value("close_id").(int)
	var lost models.Lost
	err := lc.db.QueryRow("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, "+
		"st_y(location) as longitude, picture_id FROM lost "+
		"WHERE id = $1 AND status_id != $2", id, closeId).
		Scan(&lost.Id, &lost.TypeId, &lost.AuthorId,
			&lost.Sex, &lost.Breed, &lost.Description,
			&lost.StatusId, &lost.Date,
			&lost.Latitude, &lost.Longitude, &lost.PictureId)
	if err != nil {
		return nil, err
	}
	return &lost, nil
}

/*
typeId, authorId int,
	sex, breed, description string,
	statusId int,
	date, place string
*/
func (lc *LostControllerPg) Add(ctx context.Context, params *models.Lost) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	tx := strTx.(*sql.Tx)
	var id int = 0
	// status_id = 1 (Not found). Temporarily
	query := fmt.Sprintf("INSERT INTO lost(type_id, vk_id, sex, "+
		"breed, description, status_id, location) "+
		"VALUES($1, $2, $3, $4, $5, 1, "+
		"st_GeomFromText('point(%f %f)', 4326)) RETURNING id",
		params.Latitude, params.Longitude)

	err := tx.QueryRow(query,
		params.TypeId, params.AuthorId, params.Sex,
		params.Breed, params.Description).Scan(&id)
	return id, err
}

/*
typeId int,
	sex, breed, description string,
	status int,
	date, place string, typeId int,
*/

func (lc *LostControllerPg) Search(ctx context.Context, params *models.Lost) ([]models.Lost, error) {
	ctxParams := ctx.Value("params").(map[string]interface{})
	// Get everything without parameters to search
	if features.CheckEmptyLost(params) {
		rows, err := lc.db.Query("SELECT id, type_id, "+
			"vk_id, sex, "+
			"breed, description, status_id, "+
			"date, st_x(location) as latitude, "+
			"st_y(location) as longitude, picture_id FROM lost "+
			"WHERE status_id != $1",
			ctxParams["close_id"].(int))
		if err != nil {
			return nil, err
		}
		losts, err := db.ConvertRowsToLost(rows)
		rows.Close()
		return losts, err
	}

	tx, err := lc.db.Begin()
	if err != nil {
		return nil, err
	}
	ctxParams["tx"] = tx
	ctx = context.WithValue(context.Background(), "params", ctxParams)
	searchManager := search.NewSearchManager()

	addResultToSearchManager := func(result *[]models.Lost,
		sm *search.SearchManager) {
		// Convert a slice of lost to the slice of interface{}
		// It's needed to convert the slice to the set.
		// Sets allow us to perform an operation to intersect
		// results of queries
		interfaceSlice := features.ConvertSlicesElementsToInterface(*result)
		set := set.NewSetFromSlice(interfaceSlice)
		searchManager.Add(&set)
	}
	if params.TypeId != 0 {
		result, err := lc.SearchByType(ctx, params.TypeId)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Sex != "" {
		result, err := lc.SearchBySex(ctx, params.Sex)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Breed != "" {
		result, err := lc.SearchByBreed(ctx, params.Breed)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	// if params.Location != "" {
	// 	result, err := lc.SearchByLocation(params.Location)
	// 	if err != nil {
	// 		if rollErr := tx.Rollback(); rollErr != nil {
	// 			return nil, rollErr
	// 		}
	// 		return nil, err
	// 	}
	// 	addResultToSearchManager(&result, searchManager)
	// }
	// if params.Description != "" {
	// 	result, err := lc.SearchByDescription(params.Description)
	// 	if err != nil {
	// 		if rollErr := tx.Rollback(); rollErr != nil {
	// 			return nil, rollErr
	// 		}
	// 		return nil, err
	// 	}
	// addResultToSearchManager(&result, searchManager)

	// ****** ? ***************
	// if params.Date != "" {
	// 	result, err := lc.SearchByPlace(params.Place)
	// 	if err != nil {
	// 		if rollErr := tx.Rollback(); rollErr != nil {
	// 			return nil, rollErr
	// 		}
	// 		return nil, err
	// 	}
	// 	addResultToSearchManager(&result, searchManager)
	// }
	// if params.StatusId != 0 {
	// 	result, err := lc.SearchByStatus(params.StatusId)
	// 	if err != nil {
	// 		if rollErr := tx.Rollback(); rollErr != nil {
	// 			return nil, rollErr
	// 		}
	// 		return nil, err
	// 	}
	// 	addResultToSearchManager(&result, searchManager)
	// }
	err = tx.Commit()
	if err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return nil, rollErr
		}
		return nil, err
	}

	// Now we must intersect all the sets stored in
	// the general slice called 'resultSets'

	resultSet := searchManager.GetSet()
	results := features.ConvertInterfaceElementsToLost((*resultSet).ToSlice())
	return results, nil
}

func (lc *LostControllerPg) SearchByType(ctx context.Context, typeId int) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, date, "+
		"st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM lost "+
		"WHERE type_id = $1 AND status_id != $2", typeId, closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (lc *LostControllerPg) SearchBySex(ctx context.Context, sex string) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, date, "+
		"st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM lost "+
		"WHERE LOWER(sex) = $1 AND status_id != $2", strings.ToLower(sex), closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (lc *LostControllerPg) SearchByBreed(ctx context.Context, breed string) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM lost "+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%' "+
		"AND status_id != $2", strings.ToLower(breed), closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

// func (lc *LostControllerPg) SearchByLocation(location string) ([]models.Lost, error) {
// 	rows, err := lc.db.Query("SELECT id, type_id, vk_id, sex, "+
// 		"breed, description, status_id, date, "+
// 		"location, picture_id FROM lost "+
// 		"WHERE location = $1", location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	losts, err := db.ConvertRowsToLost(rows)
// 	return losts, err
// }

// func (lc *LostControllerPg) SearchByDescription(description string) ([]models.Lost, error) {
// 	rows, err := lc.db.Query("SELECT id, type_id, vk_id, sex, "+
// 		"breed, description, status_id, date, "+
// 		"st_x(location) as latitude, st_y(location) as longitude, "+
// 		"picture_id FROM lost "+
// 		"WHERE LOWER(description) LIKE '%' || $1 || '%' "+
// 		"ORDER BY date DESC",
// 		strings.ToLower(description))
// 	if err != nil {
// 		return nil, err
// 	}
// 	losts, err := db.ConvertRowsToLost(rows)
// 	return losts, err
// }

// func (lc *LostControllerPg) SearchByStatus(statusId int) ([]models.Lost, error) {
// 	rows, err := lc.db.Query("SELECT id, type_id, vk_id, sex, "+
// 		"breed, description, status_id, "+
// 		"date, st_x(location) as latitude, st_y(location) as longitude, "+
// 		"picture_id FROM lost "+
// 		"WHERE status_id = $1", statusId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	losts, err := db.ConvertRowsToLost(rows)
// 	return losts, err
// }

// A direction is needed to specify a date (must be less or greater or equal)
func (lc *LostControllerPg) SearchByDate(date, direction string) ([]models.Lost, error) {
	if direction != "<" && direction != ">" && direction != "=" {
		return nil, errs.IncorrectDirection
	}
	sqlQuery := fmt.Sprintf("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM lost "+
		"WHERE date %s $1", direction)
	rows, err := lc.db.Query(sqlQuery, date)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) GetItemsPerPageCount() int {
	return lc.itemsPerPage
}

func (lc *LostControllerPg) GetDbAdapter() *sql.DB {
	return lc.db
}
