package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Serves the webpage created by createRoot()
func root(writer http.ResponseWriter, request *http.Request) {
	if modify() {
		r = parseList()
		fmt.Println("Parsed the list again")
	}
	var table string
	for _, member := range *r {
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
	if modify() {
		r = parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range *r {
		if item.url == host {
			if i+1 >= len(*r) {
				dest = dest + (*r)[0].url
				http.Redirect(writer, request, dest, 302)
				success = true
				break
			}
			dest = dest + (*r)[i+1].url
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
func previous(writer http.ResponseWriter, request *http.Request) {
	if modify() {
		r = parseList()
	}
	host := request.URL.Query().Get("host")
	dest, success := "https://", false
	for i, item := range *r {
		if item.url == host {
			if i-1 < 0 {
				dest = dest + (*r)[len(*r)-1].url
				http.Redirect(writer, request, dest, 302)
				break
			}
			dest = dest + (*r)[i-1].url
			http.Redirect(writer, request, dest, 302)
			break
		}
	}
	if success == false {
		http.Error(writer, "Ring member '"+host+"' not found.", 404)
	}
}

// Redirects the visitor to a random member
func random(writer http.ResponseWriter, request *http.Request) {
	if modify() {
		r = parseList()
	}
	rand.Seed(time.Now().Unix())
	dest := "https://" + (*r)[rand.Intn(len(*r)-1)].url
	http.Redirect(writer, request, dest, 302)
}
