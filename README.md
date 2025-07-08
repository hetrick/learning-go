Solutions to exercises from the book _[Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)_. Sample code and the author's solutions can be found [here](https://github.com/learning-go-book-2e).

<ins>Chapter 1</ins>
- [Hello World](./ch1/hello.go)

<ins>Chapter 2</ins>
- **Exercise 1**: Write a program where you declare an integer variable called `i` with the value `20`. Assign `i` to a floating point variable named `f`. Print out `i` and `f`. ([solution](./ch2/ex1/main.go))
- **Exercise 2**: Write a program where you declare a constant called `value` that can be assigned to both an integer and a floating point variable. Assign it to an integer called `i` and a floating point variable called `f`. Print out `i` and `f`. ([solution](./ch2/ex2/main.go))
- **Exercise 3**: Write a program with three variables, one named `b` of type `byte`, one named `smallI` of type `int32`, and one named `bigI` of type `uint64`. Assign each variable the maximum legal value for its type, then add `1` to each variable. Print out their values. ([solution](./ch2/ex3/main.go))

<ins>Chapter 3</ins>
- **Exercise 1**: Write a program that defines a variable named `greetings` of type slice of strings with the following values: `"Hello"`, `"Hola"`, `"नमस्कार"`, `"こんにちは"`, and `"Привіт"`. Create a sub-slice containing the first two values, a second subslice with the second, third, and fourth values, and a third subslice with the fourth and fifth values. Print out all four slices. ([solution](./ch3/ex1/main.go))
- **Exercise 2**: Write a program that defines a string variable called `message` with the value `"Hi 👩 and 👨"` and prints the 4th rune in it as a character, not a number. ([solution](./ch3/ex2/main.go))
- **Exercise 3**: Write a program that defines a struct called `Employee` with three fields: `firstName`, `lastName`, and `id`. The first two fields are of type `string` and the last field (`id`) is of type `int`. Create three instances of this struct using whatever values you'd like. Initialize the first one using the struct literal style without keys, the second using the struct literal style with keys, and the third with var declaration. Use dot notation to populate the fields in the third struct. Print out all three structs. ([solution](./ch3/ex3/main.go))

<ins>Chapter 4</ins>
- **Exercise 1**: Write a `for` loop that puts 100 random number between 0 and 100 into an `int` slice. ([solution](./ch4/ex1/main.go))
- **Exercise 2**: Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules: a. If the value is divisible by 2, print "Two!" b. If the value is divisible by 3, print "Three!" c. If the value is divisible by 2 and 3, print "Six!". Don't print anything else. d. Otherwise, print "Never mind". ([solution](./ch4/ex2/main.go))
- **Exercise 3**: Start a new program. In `main`, declare an `int` variable called `total`. Write a `for` loop that uses a variable named `i` to iterate from 0 (inclusive) to 10 (exclusive). After the `for` loop, print out the value of `total`. What is printed out? What is the likely bug in this code? ([solution](./ch4/ex3/main.go))

<ins>Chapter 5</ins>
- **Exercise 1**: The simple calculator program doesn't handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return 0, errors.New("division by zero"). In all other cases, return the calculated value and nil. Adjust the main function to check for this error. ([solution](./ch5/ex1/main.go))
