package API

import (
	"net/http"
	"os"
	"io"
	"log"
	"path/filepath"
)

func upload(folder string, w http.ResponseWriter, r *http.Request) {
	if folder == "" {
		http.Error(w, "No directory selected", http.StatusForbidden)
		return
	}
	path := getFullPath(folder)

	_, err := os.Stat(path)
	if err != nil {
		http.Error(w, "Cannot find the folder", http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile(filepath.Join(path, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	w.WriteHeader(200)
}
