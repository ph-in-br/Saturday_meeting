package main

import (
	"fmt"
	"math"
)

func perfectSquare(x int) bool {
	fX := float64(x)
	s := int(math.Sqrt(fX))
	return s*s == x
}

func isFibonacci(val int) bool {
	return perfectSquare(5*val*val-4) || perfectSquare(5*val*val+4)
}

func readValue() int {
	var input int

	fmt.Print("\nEnter value: ")

	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return input
}

func main() {
	value := readValue()
	if value < 0 {
		return
	}
	var result string
	if isFibonacci(value) {
		result = "YEA BOII, ITS FIBONACCI"
	} else {
		result = "OH NO, ITS NOT FIBONACCI"
	}
	fmt.Println(result)
}
