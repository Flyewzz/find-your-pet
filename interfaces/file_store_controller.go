package interfaces

import "mime/multipart"

type FileStoreController interface {
	// recordType (lost or found)
	// Returns the path of the saved file and error
	Save(file *multipart.File, id int, recordType string, fileName string) (string, error)
}
