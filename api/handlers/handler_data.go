package handlers

import (
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/managers"
)

type HandlerData struct {
	LostController     interfaces.LostController
	LostFileController interfaces.LostFileController
	LostAddingManager  *managers.LostAddingManager
	DebugMode          bool
}

func NewHandlerData(lc interfaces.LostController,
	fc interfaces.LostFileController,
	lam *managers.LostAddingManager, isDebug bool) *HandlerData {
	return &HandlerData{
		LostController:     lc,
		LostFileController: fc,
		LostAddingManager:  lam,
		DebugMode:          isDebug,
	}
}
