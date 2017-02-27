package main

import (
	"flag"
	"path/filepath"
	"./API"
	"os"
	"os/signal"
	"syscall"
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

	catchSignal(API.Clean)

	API.Run(defaultPort)

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