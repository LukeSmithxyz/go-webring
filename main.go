// SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"html/template"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

type ring struct {
	handle string
	url    string
}

type model struct {
	ring         []ring
	index        *template.Template
	ringModTime  int64
	indexModTime int64
}

// Pre-define all of our flags
var (
	flagListen  *string = flag.StringP("listen", "l", "127.0.0.1:2857", "Host and port go-webring will listen on")
	flagMembers *string = flag.StringP("members", "m", "list.txt", "Path to list of webring members")
	flagIndex   *string = flag.StringP("index", "i", "index.html", "Path to home page template")
)

func main() {
	m := model{}
	m.init()

	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    *flagListen,
		Handler: mux,
	}

	mux.HandleFunc("/", m.root)
	mux.HandleFunc("/next", m.next)
	mux.HandleFunc("/n", m.next)
	mux.HandleFunc("/previous", m.previous)
	mux.HandleFunc("/prev", m.previous)
	mux.HandleFunc("/p", m.previous)
	mux.HandleFunc("/random", m.random)
	mux.HandleFunc("/rand", m.random)
	mux.HandleFunc("/r", m.random)

	if err := httpServer.ListenAndServe(); err == http.ErrServerClosed {
		log.Println("Web server closed")
	} else {
		log.Fatalln(err)
	}
}

func (m *model) init() {
	flag.Parse()
	log.Println("Listening on", *flagListen)
	log.Println("Looking for members in", *flagMembers)
	m.parseList()
	log.Println("Found", len(m.ring), "members")
	log.Println("Building homepage with", *flagIndex)
	m.parseIndex()
}
