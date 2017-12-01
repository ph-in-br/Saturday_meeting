package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/Sekob/log_reader"
)

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	urlArray := strings.Split(r.URL.Path, "/")

	if urlArray[1] == "" {
		reader.AllApplications(w, r)
		return
	}

	if len(urlArray) == 2 {
		reader.GetConcreteApplication(w, r, urlArray[1])
		return
	}

	if len(urlArray) == 3 && urlArray[1] == "show" {
		reader.GetConcreteLog(w, r, urlArray[2])
		return
	}
}

//StartServer starts srver with hendlers
func StartServer() {
	reader.SetupDir("/var/log")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", emptyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
