package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// parseIndex parses the index template and returns a template struct.
func (s server) parseIndex() {
	tmpl, err := template.ParseFiles(*flagIndex)
	if err != nil {
		log.Fatal(err)
	}
	s.index = tmpl
}

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

// Modify takes arguments "index" or "ring" and returns true if either have been
// modified since last read
func (s server) modify(a string) bool {
	if a == "ring" {
		members, err := os.Stat(*flagMembers)
		if err != nil {
			log.Fatalln(err)
		}
		curRingModTime := members.ModTime().Unix()
		if s.ringModTime == 0 {
			s.ringModTime = curRingModTime
		} else if s.ringModTime < curRingModTime {
			return true
		}
		return false
	} else if a == "index" {
		index, err := os.Stat(*flagIndex)
		if err != nil {
			log.Fatalln(err)
		}
		curIndexModTime := index.ModTime().Unix()
		if s.indexModTime == 0 {
			s.indexModTime = curIndexModTime
		} else if s.indexModTime < curIndexModTime {
			return true
		}
		return false
	} else {
		log.Fatalln("Please call modify() with argument of either \"index\" or \"ring\"")
	}
	return true
}
