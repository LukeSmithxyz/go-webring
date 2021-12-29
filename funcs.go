package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func parseList() *[]ring {
	file, err := ioutil.ReadFile(*flagMembers)
	if err != nil {
		log.Fatal("Error while loading list of webring members: ", err)
	}
	lines := strings.Split(string(file), "\n")
	var r []ring
	for _, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		r = append(r, ring{handle: fields[0], url: fields[1]})
	}
	return &r
}

// Link returns an HTML, HTTPS link of a given URI
func link(l string) string {
	return "<a href='https://" + l + "'>" + l + "</a>"
}

// Modify returns true if the index and ring list have been modified since last
// read
func modify() bool {
	members, err := os.Stat(*flagMembers)
	if err != nil {
		log.Fatalln(err)
	}
	index, err := os.Stat(*flagIndex)
	if err != nil {
		log.Fatalln(err)
	}

	curRTime := members.ModTime().Unix()
	curITime := index.ModTime().Unix()

	if *rModTime == 0 {
		*rModTime = curRTime
	} else if *rModTime < curRTime {
		return true
	}

	if *indexModTime == 0 {
		*indexModTime = curITime
	} else if *indexModTime < curITime {
		return true
	}

	return false
}
