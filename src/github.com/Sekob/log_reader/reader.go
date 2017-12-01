package reader

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Sekob/utils"
)

var dir = "/var/log1"

//AllApplications gets all available applications in log dir
func AllApplications(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, read(), html.EscapeString(r.URL.Path))
}

//GetConcreteApplication gets all logs for concrete application
func GetConcreteApplication(w http.ResponseWriter, r *http.Request, appName string) {
	files, err := filepath.Glob(dir + "/" + appName + "*")

	utils.CheckError(err)

	strFilesNames := ""
	for _, v := range files {
		strFilesNames += "\n" + filepath.Base(v)
	}

	fmt.Fprintf(w, strFilesNames)
}

//GetConcreteLog gets concrete log file
func GetConcreteLog(w http.ResponseWriter, r *http.Request, logName string) {
	data, err := ioutil.ReadFile(dir + "/" + logName)
	utils.CheckError(err)
	fmt.Fprintf(w, "%s", data)
}

func read() string {
	files, err := ioutil.ReadDir(dir)

	utils.CheckError(err)

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

	if utils.IsNumber(s[i]) {
		return true
	}

	return false
}

//SetupDir changes dir
func SetupDir(s string) {
	dir = s
}
