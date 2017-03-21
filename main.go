package main

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"

	"gopkg.in/mgo.v2"
)

const (
	assetDirName = "assets"
	templDirName = "templates"
	assetPrefix = "/" + assetDirName + "/"
)

var (
	myProcess *os.Process
	dbSes *mgo.Session

	dbURL    string
	listen   string

	dataDir  string
	assetDir string
	templDir string

	templates *template.Template
)

func stopHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "stop\n")
	if err := myProcess.Signal(syscall.SIGTERM); err != nil {
		log.Print(err)
	}
}

func giveUp (w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, err.Error())
}

func personsHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templates.ExecuteTemplate(w, "persons", nil)
}

func main() {
	pid := os.Getpid()
	var err error
	if myProcess, err = os.FindProcess(pid); err != nil {
		log.Fatal(err)
	}
	parseFlags()
	prepare()
	defer dbSes.Close()
	start()
}

func prepare() {
	var err error
	if dbSes, err = mgo.Dial(dbURL); err != nil {
		log.Fatal(err)
	}
	prepareDB()
	assetDir = filepath.Join(dataDir, assetDirName)
	templDir = filepath.Join(dataDir, templDirName)
	http.Handle(assetPrefix, http.StripPrefix(assetPrefix,
		http.FileServer(http.Dir(assetDir))))
	http.HandleFunc("/stop", stopHandler)
	http.HandleFunc("/persons", personsHandler)

	templates = template.New(templDirName)
	templates.Delims("{{%", "%}}")
	_, err = templates.ParseFiles(
		filepath.Join(templDir, "header.html"),
		filepath.Join(templDir, "footer.html"),
		filepath.Join(templDir, "persons.html"))
	if err != nil {
		log.Fatal(err)
	}
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
	flag.StringVar(&dbURL, "db", "localhost", "MongoDB database URL");
	flag.StringVar(&listen, "listen", "localhost:3003",
		"Address and port to listen");
	flag.StringVar(&dataDir, "data", wd,
		"The directory containing data files");
	flag.Parse();
}
