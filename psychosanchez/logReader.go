package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var re = regexp.MustCompile(`(?m).*\.(gz|old|xz|\d+){1}$`)
var logFiles map[string][]string

func showLog(w http.ResponseWriter, r *http.Request) {
	logFiles = make(map[string][]string, 0)
	if len(r.URL.Query().Get("folder")) == 0 {
		searchDir := "/var/log"
		fileList := []string{}

		filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

			// If we match regex
			if !f.IsDir() && !re.MatchString(f.Name()) {
				if logFiles[f.Name()] == nil {
					logFiles[f.Name()] = []string{}
					fileList = append(fileList, f.Name())
				}
			}

			return nil
		})

		filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
			for key := range logFiles {
				if strings.HasPrefix(f.Name(), key) {
					logFiles[key] = append(logFiles[key], path)
					break
				}
			}
			return nil
		})
		for key := range logFiles {
			fmt.Fprintln(w, key)
			for _, file := range logFiles[key] {
				fmt.Fprintln(w, file)
			}
		}
		return // don't call original handler
	}
}

func registerFiles() {

}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=alex123 dbname=test_db sslmode=disable password=1488")
	if err != nil {
		print(err.Error())
	} else {
		print("connected")
	}
	defer db.Close()
	http.HandleFunc("/", showLog)
	http.ListenAndServe(":8080", nil)
}
