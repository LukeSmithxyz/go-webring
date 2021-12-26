package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
	// These are not yet implemented
	// flagCert  *string = flag.StringP("cert", "c", "cert.crt", "Path to certificate")
	// flagKey   *string = flag.StringP("key", "k", "cert.key", "Path to private certificate key")
)

func main() {
	flag.Parse()
	fmt.Println("Listening on", *flagListen)
	fmt.Println("Looking for members in", *flagMembers)
	fmt.Println("Building homepage with", *flagIndex)

	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/next", next)
	mux.HandleFunc("/previous", previous)
	mux.HandleFunc("/random", random)

	server := &http.Server{
		Addr:    *flagListen,
		Handler: mux,
	}
	server.ListenAndServe()
}

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func parseList() []ring {
	file, err := ioutil.ReadFile(*flagMembers)
	if err != nil {
		log.Fatal("Error while loading list of webring members:", err)
	}
	lines := strings.Split(string(file), "\n")
	var r []ring
	for _, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		r = append(r, ring{handle: fields[0], url: fields[1]})
	}
	return r
}

// createRoot creates the root webpage that will be served. It takes a []ring
// parameter, generates a table based on it, then inserts that into the template
// webpage.
func createRoot(r []ring) {
}

// Serves the webpage created by createRoot()
func root(writer http.ResponseWriter, request *http.Request) {
}

// Redirects the visitor to the next member, wrapping around the list if the
// next would be out-of-bounds
func next(writer http.ResponseWriter, request *http.Request) {
}

// Redirects the visitor to the previous member, wrapping around the list of the
// next would be out-of-bounds
func previous(writer http.ResponseWriter, request *http.Request) {
}

// Redirects the visitor to a random member
func random(writer http.ResponseWriter, request *http.Request) {
}
