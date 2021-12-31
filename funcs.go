package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// parseIndex parses the index template and returns a template struct.
func (m *model) parseIndex() {
	tmpl, err := template.ParseFiles(*flagIndex)
	if err != nil {
		log.Fatal(err)
	}
	m.index = tmpl
}

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func (m *model) parseList() {
	file, err := ioutil.ReadFile(*flagMembers)
	if err != nil {
		log.Fatal("Error while loading list of webring members: ", err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines[:len(lines)-1] {
		fields := strings.Fields(line)
		m.ring = append(m.ring, ring{handle: fields[0], url: fields[1]})
	}
}

// Link returns an HTML, HTTPS link of a given URI
func link(l string) string {
	return "<a href='https://" + l + "'>" + l + "</a>"
}

// Modify takes arguments "index" or "ring" and returns true if either have been
// modified since last read
func (m *model) modify(a string) bool {
	if a == "ring" {
		members, err := os.Stat(*flagMembers)
		if err != nil {
			log.Fatalln(err)
		}
		curRingModTime := members.ModTime().Unix()
		if m.ringModTime == 0 {
			m.ringModTime = curRingModTime
		} else if m.ringModTime < curRingModTime {
			return true
		}
		return false
	} else if a == "index" {
		index, err := os.Stat(*flagIndex)
		if err != nil {
			log.Fatalln(err)
		}
		curIndexModTime := index.ModTime().Unix()
		if m.indexModTime == 0 {
			m.indexModTime = curIndexModTime
		} else if m.indexModTime < curIndexModTime {
			return true
		}
		return false
	} else {
		log.Fatalln("Please call modify() with argument of either \"index\" or \"ring\"")
	}
	return true
}
