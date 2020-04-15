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

type FoundControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewFoundControllerPg(itemsPerPage int, db *sql.DB) *FoundControllerPg {
	return &FoundControllerPg{
		itemsPerPage: itemsPerPage,
		db:           db,
	}
}

func (fc *FoundControllerPg) GetById(ctx context.Context, id int) (*models.Found, error) {
	closeId := ctx.Value("close_id").(int)
	var found models.Found
	err := fc.db.QueryRow("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, "+
		"st_y(location) as longitude, picture_id FROM found "+
		"WHERE id = $1 AND status_id != $2", id, closeId).
		Scan(&found.Id, &found.TypeId, &found.AuthorId,
			&found.Sex, &found.Breed, &found.Description,
			&found.StatusId, &found.Date,
			&found.Latitude, &found.Longitude, &found.PictureId)
	if err != nil {
		return nil, err
	}
	return &found, nil
}

/*
typeId, authorId int,
	sex, breed, description string,
	statusId int,
	date, place string
*/
func (fc *FoundControllerPg) Add(ctx context.Context, params *models.Found) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	tx := strTx.(*sql.Tx)
	var id int = 0
	// status_id = 1 (Not found). Temporarily
	query := fmt.Sprintf("INSERT INTO found(type_id, vk_id, sex, "+
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

func (fc *FoundControllerPg) Search(ctx context.Context, params *models.Found) ([]models.Found, error) {
	ctxParams := ctx.Value("params").(map[string]interface{})
	// Get everything without parameters to search
	if features.CheckEmptyFound(params) {
		rows, err := fc.db.Query("SELECT id, type_id, "+
			"vk_id, sex, "+
			"breed, description, status_id, "+
			"date, st_x(location) as latitude, "+
			"st_y(location) as longitude, picture_id FROM found "+
			"WHERE status_id != $1",
			ctxParams["close_id"].(int))
		if err != nil {
			return nil, err
		}
		found, err := db.ConvertRowsToFound(rows)
		rows.Close()
		return found, err
	}

	tx, err := fc.db.Begin()
	if err != nil {
		return nil, err
	}
	ctxParams["tx"] = tx
	ctx = context.WithValue(context.Background(), "params", ctxParams)
	searchManager := search.NewSearchManager()

	addResultToSearchManager := func(result *[]models.Found,
		sm *search.SearchManager) {
		// Convert a slice of found to the slice of interface{}
		// It's needed to convert the slice to the set.
		// Sets allow us to perform an operation to intersect
		// results of queries
		interfaceSlice := features.ConvertFoundElementsToInterface(*result)
		set := set.NewSetFromSlice(interfaceSlice)
		searchManager.Add(&set)
	}
	if params.TypeId != 0 {
		result, err := fc.SearchByType(ctx, params.TypeId)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Sex != "" {
		result, err := fc.SearchBySex(ctx, params.Sex)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Breed != "" {
		result, err := fc.SearchByBreed(ctx, params.Breed)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
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
	results := features.ConvertInterfaceElementsToFound((*resultSet).ToSlice())
	return results, nil
}

func (fc *FoundControllerPg) SearchByType(ctx context.Context, typeId int) ([]models.Found, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, date, "+
		"st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM found "+
		"WHERE type_id = $1 AND status_id != $2", typeId, closeId)
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return founds, err
}

func (fc *FoundControllerPg) SearchBySex(ctx context.Context, sex string) ([]models.Found, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, date, "+
		"st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM found "+
		"WHERE LOWER(sex) = $1 AND status_id != $2", strings.ToLower(sex), closeId)
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return founds, err
}

func (fc *FoundControllerPg) SearchByBreed(ctx context.Context, breed string) ([]models.Found, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	closeId := params["close_id"].(int)
	rows, err := tx.Query("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM found "+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%' "+
		"AND status_id != $2", strings.ToLower(breed), closeId)
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return founds, err
}

// A direction is needed to specify a date (must be less or greater or equal)
func (fc *FoundControllerPg) SearchByDate(date, direction string) ([]models.Found, error) {
	if direction != "<" && direction != ">" && direction != "=" {
		return nil, errs.IncorrectDirection
	}
	sqlQuery := fmt.Sprintf("SELECT id, type_id, vk_id, sex, "+
		"breed, description, status_id, "+
		"date, st_x(location) as latitude, st_y(location) as longitude, "+
		"picture_id FROM found "+
		"WHERE date %s $1", direction)
	rows, err := fc.db.Query(sqlQuery, date)
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	return founds, err
}

func (fc *FoundControllerPg) GetItemsPerPageCount() int {
	return fc.itemsPerPage
}

func (fc *FoundControllerPg) GetDbAdapter() *sql.DB {
	return fc.db
}
