package mocks

import (
	"context"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/models"
)

type MockLostController struct {
	Losts []models.Lost
}

func NewMockLostController(losts []models.Lost) *MockLostController {
	return &MockLostController{
		Losts: losts,
	}
}

func (lc *MockLostController) GetById(ctx context.Context, id int) (*models.Lost, error) {
	for _, lost := range lc.Losts {
		if lost.Id == id {
			return &lost, nil
		}
	}
	return nil, errs.LostNotFound
}

func (lc *MockLostController) Add(ctx context.Context, params *models.Lost) (int, error) {
	return 0, nil
}

func (lc *MockLostController) Search(ctx context.Context, params *models.Lost, query string, page int) ([]models.Lost, bool, error) {
	return []models.Lost{}, false, nil
}

// Returns picture_id (int) and error. If picture is null, returns 0
func (lc *MockLostController) RemoveById(ctx context.Context, id int) (int, error) {
	_, err := lc.GetById(context.Background(), id)
	if err != nil {
		return 0, err
	}
	for index, lost := range lc.Losts {
		if lost.Id == id {
			pictureId := lost.PictureId
			if index == len(lc.Losts)-1 {
				lc.Losts = lc.Losts[:id]
			} else {
				lc.Losts = append(lc.Losts[:id], lc.Losts[(id+1):]...)
			}
			return pictureId, nil
		}
	}
	return 0, errs.LostNotFound
}

func (lc *MockLostController) GetPageCapacity() int {
	return 0
}
