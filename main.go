package main

import (
	"database/sql"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
)

const (
	assetDirName = "assets"
	templDirName = "templates"
	assetPrefix = "/" + assetDirName + "/"
)

var (
	myProcess *os.Process
	db *sql.DB

	driver   string
	dsn      string
	listen   string

	dataDir  string
	assetDir string
	templDir string
)

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
	parseFlags()
	prepare()
	start()
}

func prepare() {
	var err error
	if db, err = sql.Open(driver, dsn); err != nil {
		log.Fatal(err)
	}
	prepareDB()
	assetDir = filepath.Join(dataDir, assetDirName)
	templDir = filepath.Join(dataDir, templDirName)
	http.Handle(assetPrefix, http.StripPrefix(assetPrefix,
		http.FileServer(http.Dir(assetDir))))
	http.HandleFunc("/stop", stopHandler)
}

func start() {
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&driver, "db", "mysql", "Database driver");
	flag.StringVar(&dsn, "dsn", "warehz:warehz@localhost/warehz",
		"Data source name");
	flag.StringVar(&listen, "listen", "localhost:3003",
		"Address and port to listen");
	flag.StringVar(&dataDir, "data", wd,
		"The directory containing data files");
	flag.Parse();
}
