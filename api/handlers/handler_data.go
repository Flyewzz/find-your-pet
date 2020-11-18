package handlers

import (
	"github.com/Kotyarich/find-your-pet/interfaces"
)

type HandlerData struct {
	LostController      interfaces.LostController
	FileController      interfaces.FileController
	LostAddingManager   interfaces.LostAddingManager
	FoundController     interfaces.FoundController
	FoundAddingManager  interfaces.FoundAddingManager
	ProfileController   interfaces.ProfileController
	breedClassifier     interfaces.BreedClassifier
	FileStoreController interfaces.FileStoreController
	FileMaxSize         int64
	DebugMode           bool
}

func NewHandlerData(lc interfaces.LostController,
	fc interfaces.FileController,
	lam interfaces.LostAddingManager, fnd interfaces.FoundController,
	fam interfaces.FoundAddingManager, pc interfaces.ProfileController,
	bc interfaces.BreedClassifier, fsc interfaces.FileStoreController,
	fileMaxSize int64, isDebug bool) *HandlerData {

	return &HandlerData{
		LostController:      lc,
		FileController:      fc,
		LostAddingManager:   lam,
		FoundController:     fnd,
		FoundAddingManager:  fam,
		ProfileController:   pc,
		breedClassifier:     bc,
		FileStoreController: fsc,
		FileMaxSize:         fileMaxSize,
		DebugMode:           isDebug,
	}
}
