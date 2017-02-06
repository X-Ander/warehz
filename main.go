package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"syscall"
)

var myProcess *os.Process

func jsHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	io.WriteString(w, "// here shall be a script\n")

}

func stopHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "stop\n")
	if err := myProcess.Signal(syscall.SIGTERM); err != nil {
		log.Print(err)
	}
}

func main() {
	pid := os.Getpid()
	var err error
	if myProcess, err = os.FindProcess(pid); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/script.js", jsHandler)
	http.HandleFunc("/stop", stopHandler)
	if err := http.ListenAndServe("localhost:3003", nil); err != nil {
		log.Fatal(err)
	}
}

// vim: syntax=go
