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
// next would be out-of-bounds, and ensuring the destination returns a 200 OK
// status before performing the redirect.
func (m model) next(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	host := request.URL.Query().Get("host")
	scheme, success := "https://", false
	length := len(m.ring)
	for i, item := range m.ring {
		if item.url == host {
			for j := i + 1; j < length+i; j++ {
				dest := scheme + m.ring[j%length].url
				log.Println("Checking '" + dest + "'")
				if is200(dest) {
					log.Println("Redirecting visitor to '" + dest + "'")
					http.Redirect(writer, request, dest, 302)
					success = true
					break
				}
				log.Println("Something went wrong accessing '" + dest + "', skipping site")
			}
		}
	}
	if success == false {
		http.Error(writer, "Ring member '"+host+"' not found.", 404)
	}
}

// Redirects the visitor to the previous member, wrapping around the list if the
// next would be out-of-bounds, and ensuring the destination returns a 200 OK
// status before performing the redirect.
func (m model) previous(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	host := request.URL.Query().Get("host")
	scheme := "https://"
	length := len(m.ring)
	for index, item := range m.ring {
		if item.url == host {
			// from here to start of list
			for i := index - 1; i > 0; i-- {
				dest := scheme + m.ring[i].url
				if is200(dest) {
					log.Println("Redirecting visitor to '" + dest + "'")
					http.Redirect(writer, request, dest, 302)
					return
				}
			}
			// from end of list to here
			for i := length - 1; i > index; i-- {
				dest := scheme + m.ring[i].url
				if is200(dest) {
					log.Println("Redirecting visitor to '" + dest + "'")
					http.Redirect(writer, request, dest, 302)
					return
				}
			}
			http.Error(writer, `It would appear that either none of the ring members are accessible
(unlikely) or the backend is broken (more likely). In either case,
please email amolith@secluded.site and let him (me) know what's up.`, 500)
			return
		}
	}
	http.Error(writer, "Ring member '"+host+"' not found.", 404)
	return
}

// Redirects the visitor to a random member
func (m model) random(writer http.ResponseWriter, request *http.Request) {
	if m.modify("ring") {
		log.Println("Ring modified, clearing field and re-parsing")
		m.parseList()
	}
	rand.Seed(time.Now().Unix())
	dest := "https://" + m.ring[rand.Intn(len(m.ring))].url
	http.Redirect(writer, request, dest, 302)
}
