package managers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/models"
)

type LostAddingManager struct {
	db                    *sql.DB
	lostController        interfaces.LostController
	FileController        interfaces.FileController
	baseLostDirectoryPath string
}

func NewLostAddingManager(db *sql.DB, lc interfaces.LostController,
	lfc interfaces.FileController, lostPath string) *LostAddingManager {
	return &LostAddingManager{
		db:                    db,
		lostController:        lc,
		FileController:        lfc,
		baseLostDirectoryPath: lostPath,
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
		// A client may don't send a picture
		if file == nil {
			break
		}
		_, err = lam.FileController.AddToLost(ctx, file, lostId)
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

func (lam *LostAddingManager) Remove(id int) error {
	tx, err := lam.db.Begin()
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), "tx", tx)
	hasPicture, err := lam.lostController.RemoveById(ctx, id)
	if err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return errors.New(fmt.Sprintf("err : %v\n rollback err: %v\n", err, errRoll))
		}
		return err
	}
	if hasPicture {
		lostDirectoryPath := strconv.Itoa(id)
		fullDirectoryPath := filepath.Join(lam.baseLostDirectoryPath,
			lostDirectoryPath)
		os.RemoveAll(fullDirectoryPath)
	}
	return nil
}
