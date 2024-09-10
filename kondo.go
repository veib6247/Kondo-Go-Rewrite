package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The current working directory is: '%s'\n", cwd)

	// create folders
	folderNames := []string{"Documents", "Images", "Compressed", "Installers", "Videos", "Audio", "Others"}
	for _, folderName := range folderNames {
		if err := os.Mkdir(folderName, os.ModePerm); err != nil {
			log.Println(err)
		}
	}

	// scan dir for files
	dir, err := os.Open(cwd)
	if err != nil {
		log.Fatal(err)
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	documentsExtensions := []string{
		".txt",
		".pdf",
		".csv",
		".doc",
		".docx",
		".xls",
		".xlsx",
		".ppt",
		".pptx",
	}

	imageExtensions := []string{
		".png",
		".PNG",
		".jpg",
		".JPG",
		".gif",
		".GIF",
		".WEBP",
		".webp",
		".svg",
	}

	audioExtensions := []string{
		".mp3",
		".flac",
		".wav",
	}

	videoExtensions := []string{
		".mp4",
		".m4v",
		".mkv",
		".avi",
		".wmv",
		".mov",
		".webm",
	}

	compressedExtensions := []string{
		".zip",
		".rar",
		".7z",
	}

	installerExtensions := []string{
		".exe",
		".msi",
	}

	targetFolder := "Others"

	for _, file := range files {

		// filter files only
		if !file.IsDir() {

			ext := filepath.Ext(file.Name())

			if slices.Contains(documentsExtensions, ext) {
				targetFolder = "Documents"
			}

			if slices.Contains(imageExtensions, ext) {
				targetFolder = "Images"
			}

			if slices.Contains(audioExtensions, ext) {
				targetFolder = "Audio"
			}

			if slices.Contains(videoExtensions, ext) {
				targetFolder = "Videos"
			}

			if slices.Contains(compressedExtensions, ext) {
				targetFolder = "Compressed"
			}

			if slices.Contains(installerExtensions, ext) {
				targetFolder = "Installers"
			}

			// move each file to folders based on file extension
			currentDir := filepath.Join(cwd, file.Name())
			newDir := filepath.Join(cwd, targetFolder, file.Name())

			// don't move self
			if file.Name() != "kondo.exe" {
				// only move if file does not exist in newDir yet!
				if !isFileExists(newDir) {
					if err := os.Rename(currentDir, newDir); err != nil {
						log.Fatal(err)
					}
				}
			}

		}

	}
}

// util to check before moving file to avoid overwrite
func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
