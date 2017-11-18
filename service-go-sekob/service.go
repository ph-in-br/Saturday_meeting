package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var dir = "/var/log"

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func allApplications(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, read(), html.EscapeString(r.URL.Path))
}

func concreteApplication(w http.ResponseWriter, r *http.Request, appName string) {
	files, err := filepath.Glob(dir + "/" + appName + "*")

	checkError(err)

	strFilesNames := ""
	for _, v := range files {
		strFilesNames += "\n" + filepath.Base(v)
	}

	fmt.Fprintf(w, strFilesNames)
}

func concreteLog(w http.ResponseWriter, r *http.Request, logName string) {
	data, err := ioutil.ReadFile(dir + "/" + logName)
	checkError(err)
	fmt.Fprintf(w, "%s", data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	urlArray := strings.Split(r.URL.Path, "/")

	if urlArray[1] == "" {
		allApplications(w, r)
		return
	}

	if len(urlArray) == 2 {
		concreteApplication(w, r, urlArray[1])
		return
	}

	if len(urlArray) == 3 && urlArray[1] == "show" {
		concreteLog(w, r, urlArray[2])
		return
	}
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func checkToContinue(splited *[]string) bool {
	s := *splited
	length := len(s)
	i := length - 1

	if length == 1 {
		return false
	}

	if s[i] == "gz" || s[i] == "old" {
		return true
	}

	if isNumber(s[i]) {
		return true
	}

	return false
}

func read() string {
	files, err := ioutil.ReadDir(dir)

	checkError(err)

	m := make(map[string]bool)
	result := ""
	for _, f := range files {
		name := f.Name()

		splited := strings.Split(name, ".")

		if checkToContinue(&splited) {
			continue
		}

		m[name] = true
	}

	for key := range m {
		result += "\n" + key
	}

	return result
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", emptyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
