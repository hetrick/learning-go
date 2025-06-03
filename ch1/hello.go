package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var planet string
	n, err := fmt.Sscanf("The world in this case is Earth", "The world in this case is %s", &planet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d world found: %s\n", n, planet)
}
