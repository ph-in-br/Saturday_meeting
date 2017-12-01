package utils

import (
	"log"
	"strconv"
)

//CheckError checks if error is nil or not, if nil log.Fatal will be used
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//IsNumber checks if s is a number
func IsNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
