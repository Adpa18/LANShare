package API

import (
	"os"
	"errors"
	"io/ioutil"
	"github.com/pierrre/archivefile/zip"
)

var zips = make(map[string]ZipFile)

type ZipFile struct {
	File  *os.File
	Name  string
	From  string
	Token string
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

func genZip(path, fullPath, name string) string {
	for token, zipFile := range zips {
		if zipFile.From == path {
			return token
		}
	}

	token, err := GenerateRandomString(64)

	tmpFile, err := ioutil.TempFile(os.TempDir(), token+".zip")

	err = zip.Archive(fullPath, tmpFile, nil)
	if err != nil {
		panic(err)
	}

	zips[token] = ZipFile{tmpFile, name + ".zip", path, token}
	return token
}

func getZip(zipID string) (ZipFile, error) {
	if zipFile, ok := zips[zipID]; ok {
		return zipFile, nil
	}
	return ZipFile{}, errors.New("Cannot find zip file")
}