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
		templ_file, err := Asset("file/error.html")
		checkError(err)
		t := template.Must(template.New("error.html").Parse(string(templ_file)))
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, nil)
	} else {
		log.Info("url: " + url)
		templ_file, err := Asset("file/index.html")
		checkError(err)
		t := template.Must(template.New("index.html").Parse(string(templ_file)))
		page := page{
			URL: url,
		}
		t.Execute(w, page)
	}
}

func main() {
	// root action
	log_file, _ := Asset("file/seelog.xml")
	logger, err := log.LoggerFromConfigAsString(string(log_file))
	checkError(err)
	log.ReplaceLogger(logger)
	log.Info("server start")
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	checkError(err)
}
