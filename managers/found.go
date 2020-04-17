package managers

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/models"
)

type FoundAddingManager struct {
	db              *sql.DB
	FoundController interfaces.FoundController
	FileController  interfaces.FileController
}

func NewFoundAddingManager(db *sql.DB, fc interfaces.FoundController,
	lfc interfaces.FileController) *FoundAddingManager {
	return &FoundAddingManager{
		db:              db,
		FoundController: fc,
		FileController:  lfc,
	}
}

func (fam *FoundAddingManager) Add(ctx context.Context, params *models.Found,
	foundIdCh chan<- int,
	fileCh <-chan *models.File, errCh chan<- error) {
	tx, err := fam.db.Begin()
	if err != nil {
		errCh <- err
		return
	}
	ctx = context.WithValue(ctx, "tx", tx)
	foundId, err := fam.FoundController.Add(ctx, params)
	if err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			errCh <- errRoll
		} else {
			errCh <- err
		}
		return
	}
	foundIdCh <- foundId
	select {
	case file := <-fileCh:
		_, err = fam.FileController.AddToFound(ctx, file, foundId)
		if err != nil {
			if errRoll := tx.Rollback(); errRoll != nil {
				errCh <- errRoll
			} else {
				errCh <- errs.FileError
			}
			return
		}
	case <-ctx.Done():
		if errRoll := tx.Rollback(); errRoll != nil {
			errCh <- errRoll
		} else {
			errCh <- errs.FileOperationInterrupted
		}
		return
	}
	err = tx.Commit()
	errCh <- err
}
