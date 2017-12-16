package logReader

import (
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m).*\.(gz|old|xz|\d+){1}$`)
var logFiles map[string][]string

// GetLog return application names that created their log files
func GetLog(r *http.Request) map[string][]string {
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
	}
	return logFiles
}
