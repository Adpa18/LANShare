package API

import (
	"io/ioutil"
	"errors"
	"os"
)

type Directory struct {
	Name  string `json:"name"`
	Files []File `json:"files"`
}

type File struct {
	Name  string `json:"name"`
	IsDir bool `json:"is_dir"`
}

func listRoot() Directory {
	ret := Directory{"Root", make([]File, len(Directories))}

	i := 0
	for key := range Directories {
		ret.Files[i] = File{key, true}
		i++
	}
	return ret
}

func list(path string) (Directory, error) {
	directory := getFullPath(path)

	info, err := os.Stat(directory)
	if err != nil || !info.IsDir() {
		return Directory{}, errors.New(directory + " : Cannot find Direcotory")
	}

	files, err := ioutil.ReadDir(directory)

	ret := Directory{path, make([]File, 0)}

	for _, file := range files {
		ret.Files = append(ret.Files, File{file.Name(), file.IsDir()})
	}
	return ret, nil
}
