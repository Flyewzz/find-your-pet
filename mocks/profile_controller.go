package mocks

import (
	"context"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/models"
)

type MockProfileController struct {
	Losts  []models.Lost
	Founds []models.Found
	Error  error
}

func NewMockProfileController(lost []models.Lost, found []models.Found, err error) *MockProfileController {
	return &MockProfileController{
		Losts:  lost,
		Founds: found,
		Error:  err,
	}
}

func (pc *MockProfileController) GetLost(ctx context.Context, userId int) ([]models.Lost, error) {
	if pc.Error != nil {
		return []models.Lost{}, pc.Error
	}
	var selected []models.Lost
	for _, lost := range pc.Losts {
		if lost.AuthorId == userId {
			selected = append(selected, lost)
		}
	}
	return selected, nil
}
func (pc *MockProfileController) SetLostOpening(ctx context.Context, lostId int, statusId int) error {
	if pc.Error != nil {
		return pc.Error
	}
	for i, lost := range pc.Losts {
		if lost.Id == lostId {
			lost.StatusId = statusId
			pc.Losts[i] = lost
			return nil
		}
	}
	return errs.TheFoundNotFound
}

func (pc *MockProfileController) GetFound(ctx context.Context, userId int) ([]models.Found, error) {
	if pc.Error != nil {
		return []models.Found{}, pc.Error
	}
	var selected []models.Found
	for _, found := range pc.Founds {
		if found.AuthorId == userId {
			selected = append(selected, found)
		}
	}
	return selected, nil
}
func (pc *MockProfileController) SetFoundOpening(ctx context.Context, foundId int, statusId int) error {
	if pc.Error != nil {
		return pc.Error
	}
	for i, found := range pc.Founds {
		if found.Id == foundId {
			found.StatusId = statusId
			pc.Founds[i] = found
			return nil
		}
	}
	return errs.TheFoundNotFound
}

func (pc *MockProfileController) GetItemsPerPageCount() int {
	return 0
}
