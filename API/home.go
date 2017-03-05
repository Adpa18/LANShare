package API

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"log"
	"path/filepath"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]

	var dirs Directory
	if path == "" {
		dirs = listRoot()
	} else {
		var err error
		dirs, err = list(path)
		if err != nil {
			log.Println(err)
		}
	}

	tmplFuncMap := make(template.FuncMap)
	tmplFuncMap["Split"] = strings.Split
	tmplFuncMap["Join"] = filepath.Join

	tmpl, err := template.New("home").Funcs(tmplFuncMap).ParseFiles("WEB/index.html", "WEB/templates/dropzone.html", "WEB/templates/contextMenu.html")
	if err != nil {
		log.Println(err)
	}

	err = tmpl.Execute(w, dirs)
	if err != nil {
		log.Println(err)
	}
}

type Path struct {
	Name string
	URL  string
}

func (directory Directory) SplitPath(path string) []Path {
	dirs := strings.Split(path, "/")
	dirsPath := make([]Path, len(dirs))
	link := "/"
	for index, dir := range dirs {
		link = filepath.Join(link, dir)
		dirsPath[index] = Path{dir, link}
	}
	return dirsPath
}

func (directory Directory) GetParent() string {
	index := strings.LastIndex(directory.Name, "/")
	if index == -1 {
		return "/"
	}
	return "/" + directory.Name[:index]
}

func (file File) GetURL(root string) string {
	if root == "Root" {
		return file.Name
	}
	return "/" + filepath.Join(root, file.Name)
}

func (file File) GetClass() string {
	if file.IsDir {
		return "item folder"
	}
	return "item file"
}
