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
	"net"
	"bufio"
	"fmt"
	"strings"
)

const (
	defaultHTTPPort = 8080
	defaultTCPPort  = 8081
	defaultTCPType  = "tcp"
)

func PrintErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getDirectory() string {
	folder := "."
	flag.Parse()
	if flag.NArg() > 0 {
		folder = flag.Arg(0)
	}

	dir, _ := filepath.Abs(folder)
	return dir
}

func main() {
	runTCPServer()

	PrintErr(API.AddDirectory(getDirectory()))
	PrintErr(API.AddDirectory("/Users"))

	catchSignal(API.Clean)

	router := API.NewRouter()

	log.Printf("Listening TCP on : %d\n", defaultTCPPort)
	log.Printf("Listening HTTP on : %d\n", defaultHTTPPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(defaultHTTPPort), router))

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

func runTCPServer() {
	l, err := net.Listen(defaultTCPType, ":"+strconv.Itoa(defaultTCPPort))
	if err != nil {
		log.Printf("Connecting to TCPServer on : %d\n", defaultTCPPort)
		AddDirectoryToServer()
		os.Exit(0)
	} else {
		go handleTCPConn(l)
	}
}

func handleTCPConn(l net.Listener) {
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		folder, _ := bufio.NewReader(conn).ReadString('\n')
		err = API.AddDirectory(strings.Replace(folder, "\n", "", -1))
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			conn.Write([]byte("Sucess\n"))
		}
	}
}

func AddDirectoryToServer() {
	conn, _ := net.Dial(defaultTCPType, "localhost:"+strconv.Itoa(defaultTCPPort))
	fmt.Fprintln(conn, getDirectory())
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)
}
