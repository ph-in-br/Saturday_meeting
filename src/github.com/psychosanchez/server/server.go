package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/psychosanchez/logReader"
	"github.com/psychosanchez/server/logger"
)

func loadLogHandler(w http.ResponseWriter, r *http.Request) {
	logFiles := logReader.GetLog(r)
	for key := range logFiles {
		fmt.Fprintln(w, key)
		for _, file := range logFiles[key] {
			fmt.Fprintln(w, file)
		}
	}
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=alex123 dbname=test_db sslmode=disable password=1488")
	if err != nil {
		logger.Log(err.Error())
	} else {
		logger.Log("connected")
	}

	defer db.Close()
	http.HandleFunc("/", loadLogHandler)

	http.ListenAndServe(":8080", nil)
}
