package mocks

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/afero"
)

type FileStoreController struct {
	// recordType (lost or found)
	BaseLostDirectoryPath  string
	BaseFoundDirectoryPath string
	Storage                afero.Fs
}

func NewFileStoreController(lostPath, foundPath string, storage afero.Fs) *FileStoreController {
	return &FileStoreController{
		BaseLostDirectoryPath:  lostPath,
		BaseFoundDirectoryPath: foundPath,
		Storage:                storage,
	}
}

func (fsc *FileStoreController) Save(file *multipart.File, id int, recordType string, fileName string) error {
	idStr := strconv.Itoa(id)
	var fullDirectoryPath string
	if recordType == "lost" {
		fullDirectoryPath = filepath.Join(fsc.BaseLostDirectoryPath,
			idStr)

	} else if recordType == "found" {
		fullDirectoryPath = filepath.Join(fsc.BaseFoundDirectoryPath,
			idStr)
	} else {
		return errors.New("Unknown type of recordType")
	}
	err := fsc.Storage.MkdirAll(fullDirectoryPath,
		os.ModePerm)
	if err != nil {
		return err
	}
	//Create a name with an extension for the file
	dst, err := fsc.Storage.Create(filepath.Join(
		fullDirectoryPath,
		fileName))
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, *file)
	if err != nil {
		return err
	}
	return nil

}
