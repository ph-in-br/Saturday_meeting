package main

import "fmt"
import "os"
import "math"
import "strconv"

func forced(cnt2 uint64) {
	var prev uint64 = 1
	var prevprev uint64 = 1
	var cnt uint64 = 0
	for i := uint64(1); i <= cnt2; i += prevprev {
		cnt += 1
		prevprev = prev
		prev = i
		fmt.Println(i)
		if cnt2 == i {
			fmt.Printf("Number % is % in fibonachi ordering numbers\n", cnt2, cnt)
		}
	}
}

func correct(cnt2 uint64) {
	var cnt float64 = float64(cnt2)
	var plus float64 = math.Sqrt(5*cnt*cnt + 4)
	var minus float64 = math.Sqrt(5*cnt*cnt - 4)
	var pplus = float64(uint64(plus))
	var pminus = float64(uint64(minus))
	if plus-pplus == 0 || minus-pminus == 0 {
		fmt.Printf("Number % is % in fibonachi ordering numbers\n", cnt2, cnt)
	} else {
		fmt.Printf("Number is not in  fibonachi ordering numbers\n")
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("not enough params \n")
		return
	}
	cnt2, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
	}

	forced(cnt2)
	correct(cnt2)

}
