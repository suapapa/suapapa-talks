// +build !appengine

package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

func launchBrowser(url string) {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		// XXX:
		args = []string{"open", url}
	case "windows":
		// XXX:
		args = []string{"cmd", "/c", "start", url}
	default:
		args = []string{"google-chrome",
			"--user-data-dir=/tmp/present",
			fmt.Sprintf("--app=%s", url)}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Wait()
}
