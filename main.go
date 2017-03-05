package main

import (
	"flag"
	"path/filepath"
	"./API"
	"os"
	"os/signal"
	"syscall"
	"log"
	"net/http"
	"strconv"
)

const (
	defaultPort = 8080
)

func getDirectory() string {
	folder := "."
	if flag.NArg() > 0 {
		folder = flag.Arg(0)
	}

	dir, _ := filepath.Abs(folder)
	return dir
}

func main() {
	API.AddDirectory(getDirectory())
	API.AddDirectory("/Users")

	catchSignal(API.Clean)

	router := API.NewRouter()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(defaultPort), router))

}

func catchSignal(fn func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fn()
		os.Exit(1)
	}()
}
