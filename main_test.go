package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/X-Ander/tmpmysql"
)

func TestWarehz(t *testing.T) {
	mysql, err := tmpmysql.NewServer()

	defer func() {
		if mysql != nil {
			if err := mysql.Destroy(); err != nil { t.Error(err) }
		}
	}()

	if err != nil { t.Fatal(err) }

	driver = "mysql"
	dsn = mysql.DSN
	dataDir, err = os.Getwd()
	if err != nil { t.Fatal(err) }
	prepare()

	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	res, err := http.Get(server.URL + "/assets/script.js")
	if err != nil { t.Fatal(err) }

	str, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil { t.Fatal(err) }

	if string(str) != "// here shall be a script\n" {
		t.Errorf("Bad script, got '%s'", str)
	}
}
