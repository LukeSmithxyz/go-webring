package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	flag "github.com/spf13/pflag"
)

type ring struct {
	handle string
	url    string
}

// Pre-define all of our flags
var (
	flagPort  *int    = flag.IntP("port", "p", 9285, "Port go-webring binds to")
	flagList  *string = flag.StringP("list", "l", "list.txt", "Path to list of webring members")
	flagIndex *string = flag.StringP("index", "i", "index.html", "Path to home page template")
	flagCert  *string = flag.StringP("cert", "c", "cert.crt", "Path to certificate")
	flagKey   *string = flag.StringP("key", "k", "cert.key", "Path to private certificate key")
)

func main() {
	flag.Parse()
	ring := parseList()
	fmt.Println(ring)
}

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func parseList() []ring {
	file, err := ioutil.ReadFile(*flagList)
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

// createPage create the root webpage that will be served. It takes a []ring
// argument to create the list of members that will be placed in the target HTML
// element.
func createPage(r []ring) {
}

func next() {
}

func previous() {
}

func random() {
}
