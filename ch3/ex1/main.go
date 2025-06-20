package main

import "fmt"

func main() {
	greetings := []string{
		"Hello",
		"Hola",
		"Aloha",
		"Howdy",
		"Wazzup",
	}

	first_second := greetings[:2]
	second_third_fourth := greetings[1:4]
	fourth_fifth := greetings[3:5]

	fmt.Println("          greetings:", greetings)
	fmt.Println("       first_second:", first_second)
	fmt.Println("second_third_fourth:", second_third_fourth)
	fmt.Println("       fourth_fifth:", fourth_fifth)
}
