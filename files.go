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

// func createFile(filePath, mimeType string, size int) *File {
// 	f := File{
// 		FilePath: filePath,
// 		MimeType: mimeType,
// 		Size:     size,
// 	}

// 	return &f
// }

// func (f File) String() string {
// 	return fmt.Sprintf("{ file_path: %s, mime_type: .%s, size: %d }", f.filePath, f.mimeType, f.size)
// }

func readDir(path string) {
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	images := []database.File{}

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

		// image := createFile(fileInfo.Name(), mimeType, int(fileInfo.Size()))
		images = append(images, *image)
	}

	for _, image := range images {
		fmt.Println(image)
	}
}

func getFileMimetype(fileName string, mimeTypes []string) (string, bool) {
	r, _ := regexp.Compile("([a-z]+)$")

	mimeType := r.FindString(fileName)

	return mimeType, slices.Contains(mimeTypes, mimeType)
}

func OpenFileExplorer() {
	dir, err := zenity.SelectFile(zenity.Directory())
	if err != nil {
		log.Fatal(err)
	}

	readDir(dir)

	fmt.Println("Selected directory:", dir)
}
