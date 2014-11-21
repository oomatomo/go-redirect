package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	URL string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	url := r.FormValue("url")

	if url == "" {
		log.Println("not url pramater")
		t := template.Must(template.ParseFiles("error.html"))
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, nil)
	} else {
		log.Println("url: " + url)
		t := template.Must(template.ParseFiles("index.html"))
		page := Page{
			URL: url,
		}
		t.Execute(w, page)
	}
}

func main() {
	// root action
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
