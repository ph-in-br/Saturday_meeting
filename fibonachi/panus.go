package main

import "fmt"
import "os"
import "math"
import "strconv"
import "math/rand"
import "time"

func forced(cnt2 uint64) {
	var prev uint64 = 1
	var prevprev uint64 = 1
	var cnt uint64
	for i := uint64(1); i <= cnt2; i += prevprev {
		cnt++
		prevprev = prev
		prev = i
		fmt.Println(i)
		if cnt2 == i {
			fmt.Printf("Число %d является %d-ым числом фибоначи", cnt2, cnt)
			break
		}
	}
	fmt.Printf("Число %d не является числом фибоначи\n", cnt2)
}

func correct(cnt2 uint64) {
	cnt := float64(cnt2)
	plus := math.Sqrt(5*cnt*cnt + 4)
	minus := math.Sqrt(5*cnt*cnt - 4)
	if plus-float64(uint64(plus)) == 0 || minus-float64(uint64(minus)) == 0 {
		fmt.Printf("Число %d является числом фибоначи\n", cnt2)
	} else {
		fmt.Printf("Число %d не является числом фибоначи\n", cnt2)
	}
}

func main() {
	var number uint64
	if len(os.Args) <= 1 {
		fmt.Println("Не хватает параметров в вызове команды")
		fmt.Println("Будет использовано случайное значение")
		date := time.Now().Unix()
		rand.Seed(date)
		number = uint64(rand.Int63())
	} else {
		numberIn, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("Что-то пошло не так")
		}
		number = numberIn
	}

	fmt.Println("Пробуем с числом - ", number)
	forced(number)
	fmt.Printf("\n\n")
	correct(number)

}
