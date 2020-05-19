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
	pageCapacity        int
	db                  *sql.DB
	searchRequiredQuery string
}

func NewFoundControllerPg(pageCapacity int, db *sql.DB, query string) *FoundControllerPg {
	return &FoundControllerPg{
		pageCapacity:        pageCapacity,
		db:                  db,
		searchRequiredQuery: query,
	}
}

func (fc *FoundControllerPg) GetById(ctx context.Context, id int) (*models.Found, error) {
	var found models.Found
	var pictureId sql.NullInt32
	err := fc.db.QueryRow(fc.searchRequiredQuery+
		"WHERE id = $1", id).
		Scan(&found.Id, &found.TypeId, &found.AuthorId,
			&found.Sex, &found.Breed, &found.Description,
			&found.StatusId, &found.Date,
			&found.Latitude, &found.Longitude, &pictureId, &found.Address)
	if err != nil {
		return nil, err
	}
	// Only if picture_id is not null
	if pictureId.Valid {
		found.PictureId = int(pictureId.Int32)
	}
	return &found, nil
}

func (fc *FoundControllerPg) Add(ctx context.Context, params *models.Found) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	tx := strTx.(*sql.Tx)
	var id int = 0
	// status_id = 1 (Not found). Temporarily
	query := fmt.Sprintf("INSERT INTO found(type_id, vk_id, sex, "+
		"breed, description, status_id, location, address) "+
		"VALUES($1, $2, $3, $4, $5, 1, "+
		"st_GeomFromText('point(%f %f)', 4326), $6) RETURNING id",
		params.Latitude, params.Longitude)

	err := tx.QueryRow(query,
		params.TypeId, params.AuthorId, params.Sex,
		params.Breed, params.Description, params.Address).Scan(&id)
	return id, err
}

func (fc *FoundControllerPg) Search(ctx context.Context, params *models.Found, query string, page int) ([]models.Found, bool, error) {
	ctxParams := ctx.Value("params").(map[string]interface{})

	// Get everything without parameters to search
	if features.CheckEmptyFound(params, query) {
		rows, err := fc.db.Query(fc.searchRequiredQuery+
			"ORDER BY date DESC LIMIT $1 OFFSET $2",
			fc.pageCapacity+1, (page-1)*fc.pageCapacity)
		if err != nil {
			return nil, false, err
		}
		found, err := db.ConvertRowsToFound(rows)
		rows.Close()
		// The next page is exist if the database returns
		hasMore := (len(found) == fc.pageCapacity+1)
		if hasMore {
			// Cut the last element to make a count of records = page capacity
			found = found[:len(found)-1]
		}
		return found, hasMore, err
	}

	tx, err := fc.db.Begin()
	if err != nil {
		return nil, false, err
	}
	ctxParams["tx"] = tx
	ctxParams["query"] = fc.searchRequiredQuery
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
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Sex != "" {
		result, err := fc.SearchBySex(ctx, params.Sex)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Breed != "" {
		result, err := fc.SearchByBreed(ctx, params.Breed)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if query != "" {
		result, err := fc.SearchByTextQuery(ctx, query)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	err = tx.Commit()
	if err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return nil, false, rollErr
		}
		return nil, false, err
	}

	// Now we must intersect all the sets stored in
	// the general slice called 'resultSets'

	resultSet := searchManager.GetSet()
	results := features.ConvertInterfaceElementsToFound((*resultSet).ToSlice())
	countOfElements := len(results)
	startIndex := (page - 1) * fc.pageCapacity

	if startIndex >= countOfElements {
		return nil, false, errs.IncorrectPageNumber
	}

	endIndex := (startIndex + fc.pageCapacity) - 1
	var hasMore bool
	// if a page is incomplete
	if endIndex >= countOfElements {
		endIndex = countOfElements - 1
		// An incomplete page is the last page
		hasMore = false
		return results[startIndex:], hasMore, nil
	}
	// Check for exist of the next page
	hasMore = (countOfElements > (page * fc.pageCapacity))
	// Get a page of results
	return results[startIndex:(endIndex + 1)], hasMore, nil
}

func (fc *FoundControllerPg) SearchByType(ctx context.Context, typeId int) ([]models.Found, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	rows, err := tx.Query(searchRequiredQuery+"WHERE type_id = $1", typeId)
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
	searchRequiredQuery := params["query"].(string)
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE LOWER(sex) = $1", strings.ToLower(sex))
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
	searchRequiredQuery := params["query"].(string)
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%' ", strings.ToLower(breed))
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return founds, err
}

func (fc *FoundControllerPg) SearchByTextQuery(ctx context.Context, query string) ([]models.Found, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	sqlQuery := searchRequiredQuery + `WHERE textsearchable_index_col @@ to_tsquery('russian', $1)`
	rows, err := tx.Query(sqlQuery, query)
	if err != nil {
		return nil, err
	}
	founds, err := db.ConvertRowsToFound(rows)
	rows.Close()
	return founds, err
}

func (fc *FoundControllerPg) GetPageCapacity() int {
	return fc.pageCapacity
}

func (fc *FoundControllerPg) GetDbAdapter() *sql.DB {
	return fc.db
}

func (lc *FoundControllerPg) RemoveById(ctx context.Context, id int) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	tx := strTx.(*sql.Tx)

	var pictureId sql.NullInt32
	err := tx.QueryRow("SELECT picture_id FROM found WHERE id = $1", id).Scan(&pictureId)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec("DELETE FROM found WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	// If a picture exists
	if pictureId.Valid {
		return int(pictureId.Int32), nil
	}
	// If the picture is null
	return 0, nil
}
