package API

import (
	"strings"
	"strconv"
	"path"
	"log"
	"os"
	"errors"
)

var Directories = make(map[string]string, 0)

type Dir struct {
	Name         string `json:"name"`
	AbsDirectory string `json:"abs_directory"`
}

func AddDirectory(folder string) error {
	for _, dir := range Directories {
		if dir == folder {
			return errors.New("Folder already exist")
		}
	}

	_, err := os.Stat(folder)
	if err != nil {
		return err
	}

	dirs := strings.Split(folder, "/")

	name := dirs[len(dirs)-1]

	if _, ok := Directories[name]; !ok {
		Directories[name] = folder
	} else {
		i := 0
		for {
			i++
			newName := name + "_" + strconv.Itoa(i)
			if _, ok := Directories[newName]; !ok {
				Directories[newName] = folder
				break
			}
		}
	}

	log.Printf("Adding Directory : %s\n", folder)
	return nil
}

func getAbsDirectoryByRoot(root string) string {
	for dir, absDirectory := range Directories {
		if root == dir {
			return absDirectory
		}
	}
	return ""
}

func getFullPath(folder string) string {
	firstIndex := strings.Index(folder, "/")
	if firstIndex == -1 {
		firstIndex = len(folder)
	}

	absDirectory := getAbsDirectoryByRoot(folder[:firstIndex])

	return path.Join(absDirectory, folder[firstIndex:])
}
