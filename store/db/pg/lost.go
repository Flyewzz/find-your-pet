package pg

import (
	"database/sql"

	"github.com/Kotyarich/find-your-pet/models"
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
		"WHERE id = $1", id).Scan(&lost)
	if err != nil {
		return nil, err
	}
	return &lost, nil
}

/*
typeId, authorId int,
	sex, breed, description string,
	status int,
	date, place string
*/
func (lc *LostControllerPg) Add(params *models.Lost) (int, error) {
	var id int = 0
	// status_id = 1 (Not found). Temporary
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
	date, place stringtypeId int,
	sex, breed, description string,
	status int,
	date, place string
*/
func (lc *LostControllerPg) Search(params *models.Lost) ([]models.Lost, error) {

	rows, err := lc.db.Query("SELECT id, type_id, author_id, sex, " +
		"breed, description, status_id, date, place FROM lost")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var losts []models.Lost
	for rows.Next() {
		var lost models.Lost
		err = rows.Scan(&lost.Id, &lost.TypeId, &lost.AuthorId,
			&lost.Sex, &lost.Breed, &lost.Description,
			&lost.StatusId, &lost.Date, &lost.Place)
		if err != nil {
			continue // !
		}
		losts = append(losts, lost)
	}
	return losts, err
}

func (lc *LostControllerPg) GetItemsPerPageCount() int {
	return lc.itemsPerPage
}

func (lc *LostControllerPg) GetDbAdapter() *sql.DB {
	return lc.db
}
