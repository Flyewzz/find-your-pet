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
	pageCapacity        int
	db                  *sql.DB
	searchRequiredQuery string
}

func NewLostControllerPg(pageCapacity int, db *sql.DB,
	query string) *LostControllerPg {
	return &LostControllerPg{
		pageCapacity:        pageCapacity,
		db:                  db,
		searchRequiredQuery: query,
	}
}

func (lc *LostControllerPg) GetById(ctx context.Context, id int) (*models.Lost, error) {
	var lost models.Lost
	var pictureId sql.NullInt32
	err := lc.db.QueryRow(lc.searchRequiredQuery+
		"WHERE id = $1", id).
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

func (lc *LostControllerPg) Search(ctx context.Context, params *models.Lost, query string, page int) ([]models.Lost, bool, error) {
	ctxParams := ctx.Value("params").(map[string]interface{})
	// Get everything without parameters to search
	if features.CheckEmptyLost(params, query) {
		// lc.pageCapacity + 1 - check for an exist of the next page
		rows, err := lc.db.Query(lc.searchRequiredQuery+
			"ORDER BY date DESC LIMIT $1 OFFSET $2",
			lc.pageCapacity+1, (page-1)*lc.pageCapacity)
		if err != nil {
			return nil, false, err
		}
		defer rows.Close()
		losts, err := db.ConvertRowsToLost(rows)
		if err != nil {
			return nil, false, err
		}
		// The next page is exist if the database returns
		hasMore := (len(losts) == lc.pageCapacity+1)
		if hasMore {
			// Cut the last element to make a count of records = page capacity
			losts = losts[:len(losts)-1]
		}
		return losts, hasMore, err
	}

	tx, err := lc.db.Begin()
	if err != nil {
		return nil, false, err
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
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Sex != "" {
		result, err := lc.SearchBySex(ctx, params.Sex)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Breed != "" {
		result, err := lc.SearchByBreed(ctx, params.Breed)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, false, rollErr
			}
			return nil, false, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if query != "" {
		result, err := lc.SearchByTextQuery(ctx, query)
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
	// the general slice called 'resultSet'

	resultSet := searchManager.GetSet()
	results := features.ConvertInterfaceElementsToLost((*resultSet).ToSlice())
	countOfElements := len(results)
	if countOfElements == 0 {
		return []models.Lost{}, false, nil
	}
	startIndex := (page - 1) * lc.pageCapacity

	if startIndex >= countOfElements {
		return nil, false, errs.IncorrectPageNumber
	}

	endIndex := (startIndex + lc.pageCapacity) - 1
	var hasMore bool
	// if a page is incomplete
	if endIndex >= countOfElements {
		endIndex = countOfElements - 1
		// An incomplete page is the last page
		hasMore = false
		return results[startIndex:], hasMore, nil
	}
	// Check for exist of the next page
	hasMore = (countOfElements > (page * lc.pageCapacity))
	// Get a page of results
	return results[startIndex:(endIndex + 1)], hasMore, nil
}

func (lc *LostControllerPg) SearchByType(ctx context.Context, typeId int) ([]models.Lost, error) {
	params := ctx.Value("params").(map[string]interface{})
	tx := params["tx"].(*sql.Tx)
	searchRequiredQuery := params["query"].(string)
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE type_id = $1", typeId)
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
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE LOWER(sex) = $1", strings.ToLower(sex))
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
	rows, err := tx.Query(searchRequiredQuery+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%' ", strings.ToLower(breed))
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
	sqlQuery := searchRequiredQuery + `WHERE textsearchable_index_col @@ to_tsquery('russian', $1)`
	rows, err := tx.Query(sqlQuery, query)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	rows.Close()
	return losts, err
}

func (lc *LostControllerPg) GetPageCapacity() int {
	return lc.pageCapacity
}

func (lc *LostControllerPg) GetDbAdapter() *sql.DB {
	return lc.db
}

func (lc *LostControllerPg) RemoveById(ctx context.Context, id int) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	tx := strTx.(*sql.Tx)

	var pictureId sql.NullInt32
	err := tx.QueryRow("SELECT picture_id FROM lost WHERE id = $1", id).Scan(&pictureId)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec("DELETE FROM lost WHERE id = $1", id)
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

func (lc *LostControllerPg) GetAll() ([]models.Lost, error) {
	rows, err := lc.db.Query(lc.searchRequiredQuery)
	if err != nil {
		return []models.Lost{}, err
	}
	defer rows.Close()
	lost, err := db.ConvertRowsToLost(rows)
	return lost, nil
}

func (lc *LostControllerPg) GetSimilars(found *models.Found) ([]models.Similar, error) {
	query := `SELECT id, 
		(SELECT floor(ST_DistanceSphere(
					location,
					st_GeomFromText('point(%f %f)', 4326)
			) / 1000))
	FROM lost
	WHERE (lower(breed) = lower($1))
			AND (((EXTRACT(EPOCH FROM current_timestamp) - EXTRACT(EPOCH FROM $2::timestamp)) / 3600 > 24
				AND (EXTRACT(EPOCH FROM current_timestamp) - EXTRACT(EPOCH FROM date)) / 3600 <= 24)
						OR
					(EXTRACT(EPOCH FROM current_timestamp) - EXTRACT(EPOCH FROM $2::timestamp)) / 3600 <= 24)
			AND (
			SELECT ST_DistanceSphere(
					location,
					st_GeomFromText('point(%f %f)', 4326)
			) / 1000 <= 100);`
	rows, err := lc.db.Query(fmt.Sprintf(query, found.Latitude, found.Longitude,
		found.Latitude, found.Longitude),
		found.Breed, found.Date)
	if err != nil {
		return []models.Similar{}, err
	}
	defer rows.Close()
	var similars []models.Similar
	for rows.Next() {
		var similar models.Similar
		err = rows.Scan(&similar.Id, &similar.Distance)
		if err != nil {
			continue
		}
		similars = append(similars, similar)
	}
	return similars, nil
}
