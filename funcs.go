package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func (s server) parseList() {
	file, err := ioutil.ReadFile(*flagMembers)
	if err != nil {
		log.Fatal("Error while loading list of webring members: ", err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		s.ring = append(s.ring, ring{handle: fields[0], url: fields[1]})
	}
}

// Link returns an HTML, HTTPS link of a given URI
func link(l string) string {
	return "<a href='https://" + l + "'>" + l + "</a>"
}

// Modify returns true if the index and ring list have been modified since last
// read
func (s server) modify() bool {
	members, err := os.Stat(*flagMembers)
	if err != nil {
		log.Fatalln(err)
	}
	index, err := os.Stat(*flagIndex)
	if err != nil {
		log.Fatalln(err)
	}

	curRingModTime := members.ModTime().Unix()
	curIndexModTime := index.ModTime().Unix()

	if s.ringModTime == 0 {
		s.ringModTime = curRingModTime
	} else if s.ringModTime < curRingModTime {
		return true
	}

	if s.indexModTime == 0 {
		s.indexModTime = curIndexModTime
	} else if s.indexModTime < curIndexModTime {
		return true
	}

	return false
}
