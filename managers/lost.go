package managers

import (
	"context"
	"database/sql"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/models"
)

type LostAddingManager struct {
	db                 *sql.DB
	lostController     interfaces.LostController
	lostFileController interfaces.LostFileController
}

func NewLostAddingManager(db *sql.DB, lc interfaces.LostController,
	lfc interfaces.LostFileController) *LostAddingManager {
	return &LostAddingManager{
		db:                 db,
		lostController:     lc,
		lostFileController: lfc,
	}
}

func (lam *LostAddingManager) Add(ctx context.Context, params *models.Lost,
	lostIdCh chan<- int,
	fileCh <-chan *models.File, errCh chan<- error) {
	tx, err := lam.db.Begin()
	if err != nil {
		errCh <- err
		return
	}
	ctx = context.WithValue(ctx, "tx", tx)
	lostId, err := lam.lostController.Add(ctx, params)
	if err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			errCh <- errRoll
		} else {
			errCh <- err
		}
		return
	}
	lostIdCh <- lostId
	select {
	case file := <-fileCh:
		_, err = lam.lostFileController.Add(ctx, file, lostId)
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
