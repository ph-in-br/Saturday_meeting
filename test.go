package main;

import "fmt"
import "math"

func main (){
	fmt.Println("Result ", checkIfFibo(6765));
}

func checkIfFibo(n uint64) bool {
	var left float64 = math.Mod(math.Sqrt(float64(5*n*n - 4))*10,10);

	if left == 0 {
		return true;
	}

	left = math.Mod(math.Sqrt(float64(5*n*n + 4))*10,10);

	if left == 0 {
		return true;
	}

	return false;
}