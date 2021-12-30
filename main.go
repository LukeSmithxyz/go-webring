package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

type ring struct {
	handle string
	url    string
}

type server struct {
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
	// These are not (and might never be) implemented, we'll see
	// flagCert  *string = flag.StringP("cert", "c", "cert.crt", "Path to certificate")
	// flagKey   *string = flag.StringP("key", "k", "cert.key", "Path to private certificate key")
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    *flagListen,
		Handler: mux,
	}

	log.Fatalln(server.ListenAndServe())

	mux.HandleFunc("/", s.root)
	mux.HandleFunc("/next", s.next)
	mux.HandleFunc("/previous", s.previous)
	mux.HandleFunc("/random", s.random)
}

func (s server) init() {
	flag.Parse()
	fmt.Println("Listening on", *flagListen)
	fmt.Println("Looking for members in", *flagMembers)
	s.parseList()
	fmt.Println("Building homepage with", *flagIndex)
	s.parseIndex()
}
