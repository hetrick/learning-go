Solutions to exercises from the book _[Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)_. Sample code and the author's solutions can be found [here](https://github.com/learning-go-book-2e).

<ins>Chapter 1</ins>
- [Hello World](./ch1/hello.go)

<ins>Chapter 2</ins>
- **Exercise 1**: ([solution](./ch2/ex1/main.go))
- **Exercise 2**: ([solution](./ch2/ex2/main.go))
- **Exercise 3**: ([solution](./ch2/ex3/main.go))

<ins>Chapter 3</ins>
- **Exercise 1**: ([solution](./ch3/ex1/main.go))
- **Exercise 2**: ([solution](./ch3/ex2/main.go))
- **Exercise 3**: ([solution](./ch3/ex3/main.go))

<ins>Chapter 4</ins>
- **Exercise 1**: ([solution](./ch4/ex1/main.go))
- **Exercise 2**: ([solution](./ch4/ex2/main.go))
- **Exercise 3**: ([solution](./ch4/ex3/main.go))

<ins>Chapter 5</ins>
- **Exercise 1**: The simple calculator program doesn't handle one error case: division by zero. Change the function signature for the math operations to return both an int and an error. In the div function, if the divisor is 0, return 0, errors.New("division by zero"). In all other cases, return the calculated value and nil. Adjust the main function to check for this error. ([solution](./ch5/ex1/main.go))
