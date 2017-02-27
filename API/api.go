package API

import (
	"log"
	"net/http"
	"strconv"
	"os"
	"time"
	"net/url"
)

var zips = make(map[string]ZipFile)

type ZipFile struct {
	File  *os.File
	Name  string
	From  string
	Token string
}

const (
	ZIPType = "ZIP"
	DownloadType = "Download"
	UploadType = "Upload"
	ListType = "List"
)

func Run(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		folder := r.URL.Path[1:]

		switch getState(r.Method, r.URL.Query(), folder) {
		case ZIPType:
			zip := zips[folder]
			w.Header().Set("Content-Disposition", "attachment; filename="+zip.Name)
			http.ServeContent(w, r, zip.Name, time.Now(), zip.File)
			break
		case DownloadType:
			download(folder, w, r)
			break
		case UploadType:
			upload(folder, w, r)
			break
		case ListType:
			list(folder, w, r)
			break
		default:
			w.WriteHeader(404)
			break
		}
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func getState(method string, query url.Values, folder string) string {
	if method == "GET" {
		if _, ok := zips[folder]; ok {
			return ZIPType
		}
		if query.Get("download") == "true" {
			return DownloadType
		}
		return ListType
	} else if method == "POST" {
		if query.Get("upload") == "true" {
			return UploadType
		}
	}
	return ""
}

func Clean() {
	for _, zip := range zips {
		zip.Remove()
	}
}

func (zip ZipFile) Remove() {
	zip.File.Close()
	os.Remove(zip.File.Name())
	delete(zips, zip.Token)
}
