package API

import (
	"net/http"
	"os"
	"encoding/json"
	"github.com/pierrre/archivefile/zip"
	"io/ioutil"
)

type Download struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func download(folder string, w http.ResponseWriter, r *http.Request) {
	if folder == "" {
		http.Error(w, "No directory selected", http.StatusForbidden)
		return
	}
	path := getFullPath(folder)

	info, err := os.Stat(path)
	if err != nil {
		http.Error(w, "Cannot find the ressource", http.StatusNotFound)
		return
	}

	if !info.IsDir() {
		w.Header().Set("Content-Disposition", "attachment; filename="+info.Name())
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, path)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	filename := getZip(folder, path, info.Name())

	json.NewEncoder(w).Encode(Download{folder, filename})
}

func getZip(folder, path, name string) string {
	for token, val := range zips {
		if val.From == folder {
			return token
		}
	}

	token, err := GenerateRandomString(64)

	tmpFile, err := ioutil.TempFile(os.TempDir(), token + ".zip")

	err = zip.Archive(path, tmpFile, nil)
	if err != nil {
		panic(err)
	}

	zips[token] = ZipFile{tmpFile, name + ".zip", folder, token}
	return token
}