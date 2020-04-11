package pg

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/models"
)

type LostFileControllerPg struct {
	db *sql.DB
}

func NewLostFileControllerPg(db *sql.DB) *LostFileControllerPg {
	return &LostFileControllerPg{
		db: db,
	}
}

func (fc *LostFileControllerPg) GetById(id int) (*models.File, error) {
	row := fc.db.QueryRow("SELECT file_id, name, path from files "+
		"WHERE file_id = $1", id)
	var f models.File
	err := row.Scan(&f.Id, &f.Name, &f.Path)
	return &f, err
}

func (fc *LostFileControllerPg) Add(ctx context.Context, file *models.File,
	lostId int) (int, error) {
	strTx := ctx.Value("tx")
	if strTx == "" {
		return 0, errs.MissedTransaction
	}
	var fileId int
	tx := strTx.(*sql.Tx)
	err := tx.QueryRow("INSERT INTO files (name, path) "+
		"VALUES ($1, $2) RETURNING file_id", file.Name, file.Path).
		Scan(&fileId)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec("UPDATE lost SET picture_id = $1 "+
		"WHERE id = $2", fileId, lostId)
	return fileId, err
}
