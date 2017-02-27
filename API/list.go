package API

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Directory struct {
	Name  string `json:"name"`
	Files []File `json:"files"`
}

type File struct {
	Name  string `json:"name"`
	IsDir bool `json:"is_dir"`
}


func list(folder string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if folder == "" {
		json.NewEncoder(w).Encode(Directories)
		return
	}

	directory := getFullPath(folder)

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		http.Error(w, "Cannot find Direcotory", http.StatusNotFound)
		return
	}

	ret := Directory{folder, make([]File, 0)}

	for _, file := range files {
		ret.Files = append(ret.Files, File{file.Name(), file.IsDir()})
	}
	json.NewEncoder(w).Encode(ret)
}
