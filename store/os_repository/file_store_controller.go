package os_repository

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type FileStoreController struct {
	// recordType (lost or found)
	BaseLostDirectoryPath  string
	BaseFoundDirectoryPath string
}

func NewFileStoreController(lostPath, foundPath string) *FileStoreController {
	return &FileStoreController{
		BaseLostDirectoryPath:  lostPath,
		BaseFoundDirectoryPath: foundPath,
	}
}

func (fsc *FileStoreController) Save(file *multipart.File, id int, recordType string, fileName string) (string, error) {
	idStr := strconv.Itoa(id)
	var fullDirectoryPath string
	if recordType == "lost" {
		fullDirectoryPath = filepath.Join(fsc.BaseLostDirectoryPath,
			idStr)

	} else if recordType == "found" {
		fullDirectoryPath = filepath.Join(fsc.BaseFoundDirectoryPath,
			idStr)
	} else {
		return "", errors.New("Unknown type of recordType")
	}
	err := os.MkdirAll(fullDirectoryPath,
		os.ModePerm)
	if err != nil {
		return fullDirectoryPath, err
	}
	//Create a name with an extension for the file
	dst, err := os.Create(filepath.Join(
		fullDirectoryPath,
		fileName))
	if err != nil {
		return fullDirectoryPath, err
	}
	_, err = io.Copy(dst, *file)
	if err != nil {
		return fullDirectoryPath, err
	}
	return fullDirectoryPath, nil

}
