package pg

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/models"
)

type FileControllerPg struct {
	db *sql.DB
}

func NewFileControllerPg(db *sql.DB) *FileControllerPg {
	return &FileControllerPg{
		db: db,
	}
}

func (fc *FileControllerPg) GetById(id int) (*models.File, error) {
	row := fc.db.QueryRow("SELECT file_id, name, path from files "+
		"WHERE file_id = $1", id)
	var f models.File
	err := row.Scan(&f.Id, &f.Name, &f.Path)
	return &f, err
}

func (fc *FileControllerPg) AddToLost(ctx context.Context, file *models.File,
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

func (fc *FileControllerPg) AddToFound(ctx context.Context, file *models.File,
	foundId int) (int, error) {
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
	_, err = tx.Exec("UPDATE found SET picture_id = $1 "+
		"WHERE id = $2", fileId, foundId)
	return fileId, err
}
