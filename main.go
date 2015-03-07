package main

import (
	log "github.com/cihub/seelog"
	"html/template"
	"net/http"
	"os"
)

type page struct {
	URL string
}

func checkError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	url := r.FormValue("url")

	if url == "" {
		log.Error("not url pramater")
		t := template.Must(template.ParseFiles("error.html"))
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, nil)
	} else {
		log.Info("url: " + url)
		t := template.Must(template.ParseFiles("index.html"))
		page := page{
			URL: url,
		}
		t.Execute(w, page)
	}
}

func main() {
	// root action
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	checkError(err)
	log.ReplaceLogger(logger)
	log.Info("server start")
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	checkError(err)
}
