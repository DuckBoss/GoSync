package util

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func DirectoryContainsFile(dir string, fileName string) bool {
	filePath := filepath.Join(dir, fileName)
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func CopyFile(sourceFilePath string, destinationFilePath string) bool {
	sourceFileData, err := os.Open(sourceFilePath)
	if err != nil {
		log.Panic(err)
		return false
	}
	defer sourceFileData.Close()
	destFileData, err := os.Create(destinationFilePath)
	if err != nil {
		log.Panic(err)
		return false
	}
	defer destFileData.Close()
	_, err = io.Copy(destFileData, sourceFileData)
	if err != nil {
		log.Panic(err)
		return false
	}
	return true
}
