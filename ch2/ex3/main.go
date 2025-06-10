package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b++
	smallI++
	bigI++

	fmt.Printf("max uint8 + 1: %d\n", b)
	fmt.Printf("max int32 + 1: %d\n", smallI)
	fmt.Printf("max uint64 + 1: %d", bigI)
}
