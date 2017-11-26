package main

// import "os"
// import "math"
// import "math/rand"
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func checkInArray(arr []string, element string) bool {
	for _, file := range arr {
		if file == element {
			return true
		}
	}
	return false
}

func endsWithNumber(element string) bool {
	parts := strings.Split(element, ".")
	number, error := strconv.ParseUint(parts[len(parts)-1], 10, 64)
	if error != nil {
		return false
	}
	if number > 0 {
		return true
	}

	return false
}

func seeLogDirContans(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("/var/log")

	fslice := make([]string, len(files))
	fsliceCandidates := make([]string, len(files))
	for _, f2 := range files {
		name := f2.Name()

		fmt.Fprintf(w, "%s \n", name)

		if f2.IsDir() {
			fslice = append(fslice, name)
		} else {
			parts := strings.Split(name, ".")
			fullName := strings.Join(parts[:len(parts)-1], ".")

			if len(parts) == 1 {
				fslice = append(fslice, name)
				continue
			}

			switch last := parts[len(parts)-1]; last {
			case "gz":
				{
					if endsWithNumber(fullName) {
						fsliceCandidates = append(fsliceCandidates, fullName)
					} else {
					}
				}
			case "old":
				{
					if !checkInArray(fslice, fullName) {
						fslice = append(fslice, name)
					} else {
					}
				}
			default:
				{
					if endsWithNumber(name) {
						fsliceCandidates = append(fsliceCandidates, fullName)
					} else {
						fslice = append(fslice, name)
					}
				}
			}
		}
	}

	fmt.Fprintf(w, "\n\n %s", fslice)
	fmt.Fprintf(w, "\n\n %s", fsliceCandidates)

	for _, fname := range fsliceCandidates {
		if fname == "" {
			continue
		}
		nameSlice := strings.Split(fname, ".")
		fullName := strings.Join(nameSlice[:len(nameSlice)-1], ".")

		if !checkInArray(fslice, fullName) && !checkInArray(fsliceCandidates, fullName) {
			fslice = append(fslice, fname)
		}
	}

	fmt.Fprintf(w, "\n\n %s", fslice)
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", seeLogDirContans)
	http.ListenAndServe(":8080", nil)
}
