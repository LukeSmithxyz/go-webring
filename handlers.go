// SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>
//
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Serves the webpage created by createRoot()
func (m model) root(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	if m.modify("index") {
		log.Println("Index modified, clearing field and re-parsing")
		m.parseIndex()
	}
	var table string
	for _, member := range m.ring {
		table = table + "  <tr>\n"
		table = table + "    <td>" + member.handle + "</td>\n"
		table = table + "    <td>" + link(member.url) + "</td>\n"
		table = table + "  </tr>\n"
	}
	m.index.Execute(writer, template.HTML(table))
}

// Redirects the visitor to the next member, wrapping around the list if the
// next would be out-of-bounds
func (m model) next(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range m.ring {
		if item.url == host {
			if i+1 >= len(m.ring) {
				dest = dest + m.ring[0].url
				http.Redirect(writer, request, dest, 302)
				success = true
				break
			}
			dest = dest + m.ring[i+1].url
			http.Redirect(writer, request, dest, 302)
			success = true
			break
		}
	}
	if success == false {
		http.Error(writer, "Ring member '"+host+"' not found.", 404)
	}
}

// Redirects the visitor to the previous member, wrapping around the list if the
// next would be out-of-bounds
func (m model) previous(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range m.ring {
		if item.url == host {
			if i-1 < 0 {
				dest = dest + m.ring[len(m.ring)-1].url
				http.Redirect(writer, request, dest, 302)
				break
			}
			dest = dest + m.ring[i-1].url
			http.Redirect(writer, request, dest, 302)
			break
		}
	}
	if success == false {
		http.Error(writer, "Ring member '"+host+"' not found.", 404)
	}
}

// Redirects the visitor to a random member
func (m model) random(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	rand.Seed(time.Now().Unix())
	dest := "https://" + m.ring[rand.Intn(len(m.ring)-1)].url
	http.Redirect(writer, request, dest, 302)
}
