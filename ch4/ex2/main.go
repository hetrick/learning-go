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

	for i, rN := range rNums {
		fmt.Printf("index %d random number: %d\n", i, rN)
	}
}
