package main

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/X-Ander/tmpmysql"
)

func TestMain(t *testing.T) {
	mysql, err := tmpmysql.NewServer()
	if err != nil {
		t.Fatal(err)
	}
	defer mysql.Destroy()
	dsn = mysql.DSN
	
	dataDir, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	listen = "localhost:3003"

	go start()

	time.Sleep(time.Second)

	_, err = http.Get("http://" + listen + "/stop")
	if err != nil {
		t.Error(err)
	}
}
