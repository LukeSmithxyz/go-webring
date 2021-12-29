package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

type ring struct {
	handle string
	url    string
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

// Declare global variables for the list, index, and modification times for each
var (
	r            *[]ring
	index        *string
	rModTime     *int64
	indexModTime *int64
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/next", next)
	mux.HandleFunc("/previous", previous)
	mux.HandleFunc("/random", random)

	server := &http.Server{
		Addr:    *flagListen,
		Handler: mux,
	}
	log.Fatalln(server.ListenAndServe())
}

func init() {
	flag.Parse()
	fmt.Println("Listening on", *flagListen)
	fmt.Println("Looking for members in", *flagMembers)
	r = parseList()
	fmt.Println("Building homepage with", *flagIndex)
}
