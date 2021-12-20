package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

type ring struct {
	name string
	url  string
}

func main() {
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
