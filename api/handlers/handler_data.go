package handlers

import (
	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/managers"
)

type HandlerData struct {
	LostController      interfaces.LostController
	FileController      interfaces.FileController
	LostAddingManager   *managers.LostAddingManager
	FoundController     interfaces.FoundController
	FoundAddingManager  *managers.FoundAddingManager
	ProfileController   interfaces.ProfileController
	breedClassifier     interfaces.BreedClassifier
	FileStoreController interfaces.FileStoreController
	FileMaxSize         int64
	DebugMode           bool
}

func NewHandlerData(lc interfaces.LostController,
	fc interfaces.FileController,
	lam *managers.LostAddingManager, fnd interfaces.FoundController,
	fam *managers.FoundAddingManager, pc interfaces.ProfileController,
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
