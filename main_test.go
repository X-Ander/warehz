package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/X-Ander/tmpmysql"
)

func TestMain(t *testing.T) {
	t.Log("NewServer...")
	mysql, err := tmpmysql.NewServer()
	if err != nil {
		t.Fatalf("NewServer error: %v\n", err)
	}
	defer mysql.Destroy()
	dsn = mysql.DSN
	
	t.Log("Getwd...")
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd error: %v\n", err)
	}
	
	t.Log("Stat...")
	exe := filepath.Join(wd, "warehz")
	_, err = os.Stat(exe)
	if err != nil {
		t.Fatalf("Stat error: %v\n", err)
	}

	listen = "localhost:3003"

	t.Log("Start...")
	cmd := exec.Command(exe,
		"--dsn", dsn,
		"--listen", listen)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		t.Fatalf("Start error: %v\n", err)
	}

	t.Log("Sleep...")
	time.Sleep(3 * time.Second)

	t.Log("http.Get...")
	var res *http.Response
	if res, err = http.Get("http://" + listen + "/ping"); err != nil {
		t.Errorf("http.Get error: %v\n", err)
	}
	pong, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadAll error: %v\n", err)
	}
	if err = res.Body.Close(); err != nil {
		t.Errorf("Close error: %v\n", err)
	}
	if (string(pong) != "pong\n") {
		t.Errorf("pong = \"%s\", expected \"pong\\n\"\n", pong)
	}
	t.Logf("res.Close=%v\n", res.Close)

	t.Log("Sleep...")
	time.Sleep(1 * time.Second)

	t.Log("Signal...")
	if err = cmd.Process.Signal(syscall.SIGTERM); err != nil {
		t.Fatalf("Signal error: %v\n", err)
	}

	t.Log("Wait...")
	if err = cmd.Wait(); err != nil {
		t.Fatalf("Wait error: %v\n", err)
	}
}
