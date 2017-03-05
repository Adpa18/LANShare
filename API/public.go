package API

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"
)

func publicHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]

	http.ServeFile(w, r, filepath.Join("WEB", path))
}