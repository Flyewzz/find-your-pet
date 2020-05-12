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
	itemsPerPage        int
	db                  *sql.DB
	searchRequiredQuery string
}

func NewLostControllerPg(itemsPerPage int, db *sql.DB, query string) *LostControllerPg {
	return &LostControllerPg{
		itemsPerPage:        itemsPerPage,
		db:                  db,
		searchRequiredQuery: query,
	}
}

func (lc *LostControllerPg) GetById(ctx context.Context, id int) (*models.Lost, error) {
	closeId := ctx.Value("close_id").(int)
	var lost models.Lost
	var pictureId sql.NullInt32
	err := lc.db.QueryRow(lc.searchRequiredQuery+
		"WHERE id = $1 AND status_id != $2", id, closeId).
		Scan(&lost.Id, &lost.TypeId, &lost.AuthorId,
			&lost.Sex, &lost.Breed, &lost.Description,
			&lost.StatusId, &lost.Date,
			&lost.Latitude, &lost.Longitude, &pictureId, &lost.Address)
	if err != nil {
		return nil, err
	}
	// If the user added a picture
	if pictureId.Valid {
		lost.PictureId = int(pictureId.Int32)
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
		"breed, description, status_id, location, address) "+
		"VALUES($1, $2, $3, $4, $5, 1, "+
		"st_GeomFromText('point(%f %f)', 4326), $6) RETURNING id",
		params.Latitude, params.Longitude)

	err := tx.QueryRow(query,
		params.TypeId, params.AuthorId, params.Sex,
		params.Breed, params.Description, params.Address).Scan(&id)
	return id, err
}

/*
typeId int,
	sex, breed, description string,
	status int,
	date, place string, typeId int,
*/

func (lc *LostControllerPg) Search(ctx context.Context, params *models.Lost, query string) ([]models.Lost, error) {
	ctxParams := ctx.Value("params").(map[string]interface{})
	// Get everything without parameters to search
	if features.CheckEmptyLost(params, query) {
		rows, err := lc.db.Query(lc.searchRequiredQuery+
			"WHERE status_id != $1 ORDER BY date DESC",
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
	ctxParams["query"] = lc.searchRequiredQuery
	ctx = context.WithValue(context.Background(), "params", ctxParams)
	searchManager := search.NewSearchManager()

	addResultToSearchManager := func(result *[]models.Lost,
		sm *search.SearchManager) {
		// Convert a slice of lost to the slice of interface{}
		// It's needed to convert the slice to the set.
		// Sets allow us to perform an operation to intersect
		// results of queries
		interfaceSlice := features.ConvertLostElementsToInterface(*result)
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
	if query != "" {
		result, err := lc.SearchByTextQuery(ctx, query)
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
	results := features.ConvertInterfaceElementsToLost((*resultSet).ToSlice())
	return results, nil
}

func (lc *LostControllerPg) SearchByType(ctx context.Context, typeId int) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	closeId := params["close_id"].(int)
	rows, err := tx.Query(searchRequiredQuery+
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
	searchRequiredQuery := params["query"].(string)
	closeId := params["close_id"].(int)
	rows, err := tx.Query(searchRequiredQuery+
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
	searchRequiredQuery := params["query"].(string)
	closeId := params["close_id"].(int)
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%' "+
		"AND status_id != $2", strings.ToLower(breed), closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

// A direction is needed to specify a date (must be less or greater or equal)
func (lc *LostControllerPg) SearchByDate(ctx context.Context, date, direction string) ([]models.Lost, error) {
	if direction != "<" && direction != ">" && direction != "=" {
		return nil, errs.IncorrectDirection
	}
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	closeId := params["close_id"].(int)
	sqlQuery := fmt.Sprintf(searchRequiredQuery+
		"WHERE date %s $1 AND status_id != $2", direction)
	rows, err := tx.Query(sqlQuery, date, closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (lc *LostControllerPg) SearchByTextQuery(ctx context.Context, query string) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	closeId := params["close_id"].(int)
	sqlQuery := searchRequiredQuery + `WHERE textsearchable_index_col @@ to_tsquery('russian', $1) 
	AND status_id != $2`
	rows, err := tx.Query(sqlQuery, query, closeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (lc *LostControllerPg) GetItemsPerPageCount() int {
	return lc.itemsPerPage
}

func (lc *LostControllerPg) GetDbAdapter() *sql.DB {
	return lc.db
}
