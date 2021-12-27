package main

import (
	"html/template"
	"log"
	"net/http"
)

// Serves the webpage created by createRoot()
func root(writer http.ResponseWriter, request *http.Request) {
	r := parseList()
	table := "\n" + `<table>
  <tr>
    <th>Fedi handle</th>
    <th>Site URL</th>
  </tr>` + "\n"
	for _, member := range r {
		table = table + "  <tr>\n"
		table = table + "    <td>" + member.handle + "</td>\n"
		table = table + "    <td>" + link(member.url) + "</td>\n"
		table = table + "  </tr>\n"
	}
	table = table + "</table>\n"

	tmpl, err := template.ParseFiles(*flagIndex)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(writer, template.HTML(table))
}

// Redirects the visitor to the next member, wrapping around the list if the
// next would be out-of-bounds
func next(writer http.ResponseWriter, request *http.Request) {
}

// Redirects the visitor to the previous member, wrapping around the list of the
// next would be out-of-bounds
func previous(writer http.ResponseWriter, request *http.Request) {
}

// Redirects the visitor to a random member
func random(writer http.ResponseWriter, request *http.Request) {
}
