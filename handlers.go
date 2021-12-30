package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// Serves the webpage created by createRoot()
func (s server) root(writer http.ResponseWriter, request *http.Request) {
	if s.modify("ring") {
		s.parseList()
	} else if s.modify("index") {
		s.parseIndex()
	}
	var table string
	for _, member := range s.ring {
		table = table + "  <tr>\n"
		table = table + "    <td>" + member.handle + "</td>\n"
		table = table + "    <td>" + link(member.url) + "</td>\n"
		table = table + "  </tr>\n"
	}
	s.index.Execute(writer, template.HTML(table))
}

// Redirects the visitor to the next member, wrapping around the list if the
// next would be out-of-bounds
func (s server) next(writer http.ResponseWriter, request *http.Request) {
	if s.modify("ring") {
		s.parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range s.ring {
		if item.url == host {
			if i+1 >= len(s.ring) {
				dest = dest + s.ring[0].url
				http.Redirect(writer, request, dest, 302)
				success = true
				break
			}
			dest = dest + s.ring[i+1].url
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
func (s server) previous(writer http.ResponseWriter, request *http.Request) {
	if s.modify("ring") {
		s.parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range s.ring {
		if item.url == host {
			if i-1 < 0 {
				dest = dest + s.ring[len(s.ring)-1].url
				http.Redirect(writer, request, dest, 302)
				break
			}
			dest = dest + s.ring[i-1].url
			http.Redirect(writer, request, dest, 302)
			break
		}
	}
	if success == false {
		http.Error(writer, "Ring member '"+host+"' not found.", 404)
	}
}

// Redirects the visitor to a random member
func (s server) random(writer http.ResponseWriter, request *http.Request) {
	if s.modify("ring") {
		s.parseList()
	}
	rand.Seed(time.Now().Unix())
	dest := "https://" + s.ring[rand.Intn(len(s.ring)-1)].url
	http.Redirect(writer, request, dest, 302)
}
