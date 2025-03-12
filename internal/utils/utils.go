package utils

import (
	"errors"
	"log"
	"os"
)

// util to create folders
func CreateFolders() {
	folderNames := []string{
		"Documents",
		"Images",
		"Compressed",
		"Installers",
		"Videos",
		"Audio",
		"Others",
	}

	for _, folderName := range folderNames {
		if err := os.Mkdir(folderName, os.ModePerm); err != nil {
			log.Println(err)
		}
	}
}

// util to check before moving file to avoid overwrite
func IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
