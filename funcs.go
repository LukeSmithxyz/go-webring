// SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// parseIndex parses the index template and returns a template struct.
func (m *model) parseIndex() {
	m.index = nil
	tmpl, err := template.ParseFiles(*flagIndex)
	if err != nil {
		log.Fatal(err)
	}
	m.index = tmpl
	tmplStat, err := os.Stat(*flagIndex)
	if err != nil {
		log.Fatalln(err)
	}
	m.indexModTime = tmplStat.ModTime().Unix()
}

// List parses the list of members, appends the data to a slice of type list,
// then returns the slice
func (m *model) parseList() {
	m.ring = nil
	file, err := ioutil.ReadFile(*flagMembers)
	if err != nil {
		log.Fatal("Error while loading list of webring members: ", err)
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines[:len(lines)-1] {
		// fields := strings.Fields(line)
		fields := strings.Split(line, "\t")
		m.ring = append(m.ring, ring{handle: fields[0], url: fields[1]})
	}
	fileStat, err := os.Stat(*flagMembers)
	if err != nil {
		log.Fatalln(err)
	}
	m.ringModTime = fileStat.ModTime().Unix()
}

// Modify takes arguments "index" or "ring" and returns true if either have been
// modified since last read
func (m *model) modify(a string) bool {
	if a == "ring" {
		ringStat, err := os.Stat(*flagMembers)
		if err != nil {
			log.Fatalln(err)
		}
		curRingModTime := ringStat.ModTime().Unix()
		if m.ringModTime < curRingModTime {
			return true
		}
		return false
	} else if a == "index" {
		indexStat, err := os.Stat(*flagIndex)
		if err != nil {
			log.Fatalln(err)
		}
		curIndexModTime := indexStat.ModTime().Unix()
		if m.indexModTime < curIndexModTime {
			return true
		}
		return false
	} else {
		log.Fatalln("Please call modify() with argument of either \"index\" or \"ring\"")
	}
	return true
}

func is200(site string) bool {
	resp, err := http.Get(site)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func getRefDomain(request *http.Request) string {
	ref := request.Referer()
	parsed, err := url.Parse(ref)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(parsed.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	return domain
}
