package main

import (
	"example/hello/operations"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Enter two numbers to add:")

	var inputOperation string

	fmt.Scanln(&inputOperation)

	var x, y int
	
	x, _ = strconv.Atoi(strings.Split(inputOperation, " ")[0])
	y, _ = strconv.Atoi(strings.Split(inputOperation, " ")[1])

	fmt.Println("Enter the operation to perform:")
	fmt.Println("type 'add' to perform addition")

	var operation string
	fmt.Scanln(&operation)

	if operation == "add" {
		operations.Addition(x, y)
	} else if operation == "sub" {
		operations.Subtraction(x, y)

	} else {
		fmt.Println("Invalid operation")
	}
}
