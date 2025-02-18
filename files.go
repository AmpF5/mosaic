package main

import (
	"fmt"
	"log"
	"mosaic/database"
	"os"
	"regexp"
	"slices"

	"github.com/ncruces/zenity"
)

var mimeTypes = []string{"jpg", "png", "jpeg", "heic"}

func readDir(path string, app *App) {
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	createdFiles := []database.File{}

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return
		}

		if fileInfo.IsDir() {
			continue
		}

		mimeType, founded := getFileMimetype(fileInfo.Name(), mimeTypes)

		if !founded {
			continue
		}

		file, err := app.db.CreateFile(app.ctx, database.CreateFileParams{
			FilePath: fileInfo.Name(),
			MimeType: mimeType,
			Size:     int64(fileInfo.Size()),
		})

		if err != nil {
			log.Fatalf("Error during saving file %s, %s", fileInfo.Name(), err)
		}

		createdFiles = append(createdFiles, file)
	}

	for _, image := range createdFiles {
		fmt.Println(image)
	}
}

func getFileMimetype(fileName string, mimeTypes []string) (string, bool) {
	r, _ := regexp.Compile("([a-z]+)$")

	mimeType := r.FindString(fileName)

	return mimeType, slices.Contains(mimeTypes, mimeType)
}

func OpenFileExplorer(app *App) {
	dir, err := zenity.SelectFile(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}

	readDir(dir, app)

	fmt.Println("Selected directory:", dir)
}
