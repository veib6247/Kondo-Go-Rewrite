package main

import (
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

	documentsExtensions := []string{".txt"}

	for _, file := range files {
		// read and filter files only
		if !file.IsDir() {

			// move each file to folders based on file extension
			if slices.Contains(documentsExtensions, filepath.Ext(file.Name())) {
				fmt.Printf("%s is in the list\n", file.Name())

			}
		}

	}
}
