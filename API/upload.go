package API

import (
	"os"
	"io"
	"path/filepath"
	"log"
	"errors"
)

func upload(path string, filename string, src io.Reader) error {
	if path == "" {
		return errors.New("No directory selected")
	}
	fullPath := getFullPath(path)

	_, err := os.Stat(fullPath)
	if err != nil {
		return errors.New("Cannot find the folder")
	}

	outFile, err := os.OpenFile(filepath.Join(fullPath, filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer outFile.Close()

	log.Printf("Uploading %s in %s\n", filename, fullPath)

	io.Copy(outFile, src)
	return nil
}