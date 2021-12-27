package main

import (
	"io/ioutil"
	"log"
	"strings"
)

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

// Link returns an HTML, HTTPS link of a given URI
func link(l string) string {
	return "<a href='https://" + l + "'>" + l + "</a>"
}
