package pg

import (
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

func (lc *LostControllerPg) GetById(id int) (*models.Lost, error) {
	var lost models.Lost
	err := lc.db.QueryRow("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE id = $1", id).Scan(&lost.Id, &lost.TypeId, &lost.AuthorId,
		&lost.Sex, &lost.Breed, &lost.Description,
		&lost.StatusId, &lost.Date, &lost.Place)
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
func (lc *LostControllerPg) Add(params *models.Lost) (int, error) {
	var id int = 0
	// status_id = 1 (Not found). Temporarily
	err := lc.db.QueryRow("INSERT INTO lost(type_id, author_id, sex, "+
		"breed, description, status_id, place) "+
		"VALUES($1, $2, $3, $4, $5, 1, $6) RETURNING id",
		params.TypeId, params.AuthorId, params.Sex,
		params.Breed, params.Description, params.Place).Scan(&id)
	return id, err
}

/*
typeId int,
	sex, breed, description string,
	status int,
	date, place string, typeId int,
*/

func (lc *LostControllerPg) Search(params *models.Lost) ([]models.Lost, error) {
	tx, err := lc.db.Begin()
	if err != nil {
		return nil, err
	}
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
		result, err := lc.SearchByType(params.TypeId)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Sex != "" {
		result, err := lc.SearchBySex(params.Sex)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Breed != "" {
		result, err := lc.SearchByBreed(params.Breed)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Place != "" {
		result, err := lc.SearchByPlace(params.Place)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Description != "" {
		result, err := lc.SearchByDescription(params.Description)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.Date != "" {
		result, err := lc.SearchByPlace(params.Place)
		if err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return nil, rollErr
			}
			return nil, err
		}
		addResultToSearchManager(&result, searchManager)
	}
	if params.StatusId != 0 {
		result, err := lc.SearchByStatus(params.StatusId)
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
		return nil, err
	}

	// Now we must intersect all the sets stored in
	// the general slice called 'resultSets'

	resultSet := searchManager.GetSet()
	results := features.ConvertInterfaceElementsToLost((*resultSet).ToSlice())
	return results, nil
}

func (lc *LostControllerPg) SearchByType(typeId int) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE type_id = $1", typeId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) SearchBySex(sex string) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE LOWER(sex) = $1", strings.ToLower(sex))
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) SearchByPlace(place string) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE LOWER(place) LIKE '%' || $1 || '%'", strings.ToLower(place))
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) SearchByDescription(description string) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE LOWER(description) LIKE '%' || $1 || '%' "+
		"ORDER BY date DESC",
		strings.ToLower(description))
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) SearchByBreed(breed string) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE LOWER(breed) LIKE '%' || $1 || '%'", strings.ToLower(breed))
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

func (lc *LostControllerPg) SearchByStatus(statusId int) ([]models.Lost, error) {
	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
		"WHERE status_id = $1", statusId)
	if err != nil {
		return nil, err
	}
	losts, err := db.ConvertRowsToLost(rows)
	return losts, err
}

// A direction is needed to specify a date (must be less or greater or equal)
func (lc *LostControllerPg) SearchByDate(date, direction string) ([]models.Lost, error) {
	if direction != "<" && direction != ">" && direction != "=" {
		return nil, errs.IncorrectDirection
	}
	sqlQuery := fmt.Sprintf("SELECT id, type_id, author_id, sex, "+
		"breed, description, status_id, date, place FROM lost "+
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
