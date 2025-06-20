package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	structOne := Employee{
		"whitetail",
		"deer",
		1,
	}

	structTwo := Employee{
		firstName: "ruffed",
		lastName:  "grouse",
		id:        2,
	}

	var structThree = Employee{
		structOne.firstName,
		structTwo.lastName,
		3,
	}

	fmt.Println("employee 1:", structOne)
	fmt.Println("employee 2:", structTwo)
	employee3 := fmt.Sprintf("employee %d: %v\n", structThree.id, structThree)
	fmt.Printf(employee3)
}
