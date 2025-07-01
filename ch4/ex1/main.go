package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var rNums []int

	n := 0
	for ; n < 100; n++ {
		randNum := rand.Intn(100)
		rNums = append(rNums, randNum)
	}

	fmt.Println("count of random numbers:", len(rNums))
}
