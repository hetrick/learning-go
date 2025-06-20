package main

import "fmt"

func main() {
	var message string = "Hi 😎 🥸"

	var runeMessage []rune = []rune(message)

	fmt.Println("4th rune:", runeMessage[3])
}
