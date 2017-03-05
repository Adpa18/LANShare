package API

import (
	"os"
	"errors"
)

type Download struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	isZip bool
}

func download(path string) (Download, error) {
	if path == "" {
		return Download{}, errors.New("No directory selected")
	}
	fullPath := getFullPath(path)

	info, err := os.Stat(fullPath)
	if err != nil {
		return Download{}, errors.New("Cannot find the ressource")
	}

	if !info.IsDir() {
		return Download{info.Name(), fullPath, false}, nil
	}
	return Download{path, genZip(path, fullPath, info.Name()), true}, nil
}