package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Serves the webpage created by createRoot()
func root(writer http.ResponseWriter, request *http.Request) {
	r := parseList()
	var table string
	for _, member := range r {
		table = table + "  <tr>\n"
		table = table + "    <td>" + member.handle + "</td>\n"
		table = table + "    <td>" + link(member.url) + "</td>\n"
		table = table + "  </tr>\n"
	}

	tmpl, err := template.ParseFiles(*flagIndex)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(writer, template.HTML(table))
}

// Redirects the visitor to the next member, wrapping around the list if the
// next would be out-of-bounds
func next(writer http.ResponseWriter, request *http.Request) {
	host := request.URL.Query().Get("host")
	dest := "https://"
	list := parseList()
	for i, item := range list {
		if item.url == host {
			if i+1 >= len(list) {
				dest = dest + list[0].url
				http.Redirect(writer, request, dest, 302)
				break
			}
			dest = dest + list[i+1].url
			http.Redirect(writer, request, dest, 302)
			break
		}
	}
	// TODO: Print output to the page telling the visitor that the member
	// couldn't be found
}

// Redirects the visitor to the previous member, wrapping around the list of the
// next would be out-of-bounds
func previous(writer http.ResponseWriter, request *http.Request) {
	host := request.URL.Query().Get("host")
	dest := "https://"
	list := parseList()
	for i, item := range list {
		if item.url == host {
			if i-1 < 0 {
				dest = dest + list[len(list)-1].url
				http.Redirect(writer, request, dest, 302)
				break
			}
			dest = dest + list[i-1].url
			http.Redirect(writer, request, dest, 302)
			break
		}
	}
	// TODO: Print output to the page telling the visitor that the member
	// couldn't be found
}

// Redirects the visitor to a random member
func random(writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().Unix())
	list := parseList()
	dest := "https://" + list[rand.Intn(len(list)-1)].url
	http.Redirect(writer, request, dest, 302)
}
