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

type FoundAddingManager struct {
	db                    *sql.DB
	FoundController       interfaces.FoundController
	FileController        interfaces.FileController
	baseLostDirectoryPath string
}

func NewFoundAddingManager(db *sql.DB, fc interfaces.FoundController,
	lfc interfaces.FileController, foundPath string) *FoundAddingManager {
	return &FoundAddingManager{
		db:                    db,
		FoundController:       fc,
		FileController:        lfc,
		baseLostDirectoryPath: foundPath,
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
		// A client may don't send a picture
		if file == nil {
			break
		}
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

func (fam *FoundAddingManager) Remove(id int) error {
	tx, err := fam.db.Begin()
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), "tx", tx)
	pictureId, err := fam.FoundController.RemoveById(ctx, id)
	if err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return errors.New(fmt.Sprintf("err : %v\n rollback err: %v\n", err, errRoll))
		}
		return err
	}
	if pictureId != 0 {
		err = fam.FileController.Remove(ctx, pictureId)
		if err != nil {
			if errRoll := tx.Rollback(); errRoll != nil {
				return errors.New(fmt.Sprintf("err : %v\n rollback err: %v\n", err, errRoll))
			}
			return err
		}
		lostDirectoryPath := strconv.Itoa(id)
		fullDirectoryPath := filepath.Join(fam.baseLostDirectoryPath,
			lostDirectoryPath)
		err = os.RemoveAll(fullDirectoryPath)
		if err != nil {
			if errRoll := tx.Rollback(); errRoll != nil {
				return errors.New(fmt.Sprintf("err : %v\n rollback err: %v\n", err, errRoll))
			}
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return errors.New(fmt.Sprintf("err : %v\n rollback err: %v\n", err, errRoll))
		}
		return err
	}
	return nil
}
