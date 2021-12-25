package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

type ring struct {
	name string
	url  string
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
	list()
}

// List parses the list of members and appends data to the `ring` struct
func list() {
	bytes, _ := os.ReadFile(os.Args[1])
	file := strings.Split(string(bytes), "\n")
	for _, line := range file {
		fmt.Println(line)
	}
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
