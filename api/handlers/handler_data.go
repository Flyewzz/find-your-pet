package handlers

import (
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/managers"
)

type HandlerData struct {
	LostController     interfaces.LostController
	LostFileController interfaces.LostFileController
	LostAddingManager  *managers.LostAddingManager
}

func NewHandlerData(lc interfaces.LostController,
	fc interfaces.LostFileController,
	lam *managers.LostAddingManager) *HandlerData {
	return &HandlerData{
		LostController:     lc,
		LostFileController: fc,
		LostAddingManager:  lam,
	}
}
