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

	for _, rN := range rNums {
		if rN%2 == 0 {
			fmt.Printf("%d is divisible by two!\n", rN)
		} else if rN%3 == 0 {
			fmt.Printf("%d is divisible by three!\n", rN)
		} else if rN%2 == 0 && rN%3 == 0 {
			fmt.Printf("%d is divisible by six!\n", rN)
			break
		} else {
			fmt.Print("never mind\n")
		}
	}
}
