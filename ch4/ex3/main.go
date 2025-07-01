package main

import "fmt"

func main() {
	var total int
	var i int
	for ; i < 10; i++ {
		// total := total + i
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)

}
