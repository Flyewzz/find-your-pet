package handlers

import "github.com/Kotyarich/find-your-pet/interfaces"

type HandlerData struct {
	LostController interfaces.LostController
}

func NewHandlerData(lc interfaces.LostController) *HandlerData {
	return &HandlerData{
		LostController: lc,
	}
}
