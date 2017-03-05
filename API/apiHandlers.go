package API

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"time"
)

func listRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret := listRoot()

	json.NewEncoder(w).Encode(ret)
	return
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := list(mux.Vars(r)["path"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(ret)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	ret, err := download(mux.Vars(r)["path"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !ret.isZip {
		w.Header().Set("Content-Disposition", "attachment; filename="+ret.Name)
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, ret.Path)
		return
	}

	zip, err := getZip(ret.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+zip.Name)
	http.ServeContent(w, r, zip.Name, time.Now(), zip.File)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = upload(mux.Vars(r)["path"], handler.Filename, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.WriteHeader(200)
}

func zipHandler(w http.ResponseWriter, r *http.Request) {
	zipFile, err := getZip(mux.Vars(r)["zipID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+zipFile.Name)
	http.ServeContent(w, r, zipFile.Name, time.Now(), zipFile.File)
}