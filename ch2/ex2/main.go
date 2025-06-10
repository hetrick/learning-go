package main

import "fmt"

func main() {
	const value = 30

	var i int = value
	var f float64 = value

	fmt.Printf("value %d as\nint: %d\nfloat: %f", value, i, f)
}
