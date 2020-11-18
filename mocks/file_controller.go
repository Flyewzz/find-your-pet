package mocks

import (
	"context"

	"github.com/Kotyarich/find-your-pet/models"
)

type FileController struct {
	currentId int
}

func NewFileController() *FileController {
	return &FileController{
		currentId: 1,
	}
}

func (fc *FileController) GetById(id int) (*models.File, error) {
	return nil, nil
}

func (fc *FileController) AddToLost(ctx context.Context, file *models.File, lostId int) (int, error) {
	id := fc.currentId
	fc.currentId++
	return id, nil
}

func (fc *FileController) AddToFound(ctx context.Context, file *models.File, foundId int) (int, error) {
	return 0, nil
}

func (fc *FileController) Remove(ctx context.Context, fileId int) error {
	return nil
}
