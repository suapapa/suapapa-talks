// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

import (
	"flag"
	/* "go/build" */
	"log"
	"net/http"
	"strings"

	"./go_present"
)

var basePath string

func main() {
	httpListen := flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	flag.StringVar(&basePath, "base", "", "base path for slide template and static resources")
	flag.BoolVar(&present.PlayEnabled, "play", true, "enable playground (permit execution of arbitrary user code)")
	flag.Parse()

	if basePath == "" {
		basePath = "./"
	}

	if present.PlayEnabled {
		HandleSocket("/socket")
	}
	http.Handle("/static/", http.FileServer(http.Dir(basePath)))

	go func() {
		if ok := waitServer("http://127.0.0.1:3999"); ok == false {
			log.Printf("Timeout on waiting server")
			return
		}

		launchBrowser("http://127.0.0.1:3999/slide/harder_way_to_arduino_with_python/harder_way_to_arduino_with_python.slide")
	}()

	if !strings.HasPrefix(*httpListen, "127.0.0.1") &&
		!strings.HasPrefix(*httpListen, "localhost") &&
		present.PlayEnabled {
		log.Print(localhostWarning)
	}

	log.Printf("Open your web browser and visit http://%s/", *httpListen)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}

const localhostWarning = `
WARNING!  WARNING!  WARNING!

The present server appears to be listening on an address that is not localhost.
Anyone with access to this address and port will have access to this machine as
the user running present.

To avoid this message, listen on localhost or run with -play=false.

If you don't understand this message, hit Control-C to terminate this process.

WARNING!  WARNING!  WARNING!
`
