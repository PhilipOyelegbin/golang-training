/*Create a Go program that takes two integer inputs from the user and an operation (+, -, *, /).
Based on the operation, perform the calculation and print the result.
Handle division by zero.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to sum two numbers
func sum2Num(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

// Function to subtract two numbers
func sub2Num(num1 int, num2 int) (result int) {
	result = num1 - num2
	return result
}

// Function to multiply two numbers
func mul2Num(num1 int, num2 int) (result int) {
	result = num1 * num2
	return result
}

// Function to multiply two numbers
func div2Num(num1 int, num2 int) (result int) {
	if num2 == 0 {
		panic("Division by zero is not allowed")
	} else if num1 == 0 {
		return 0
	} else {
		result = num1 / num2
		return result
	}
}

func main() {
	fmt.Println("##################################################################")
	fmt.Println("Welcome to the Calculator!")
	fmt.Println("##################################################################")

	num1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
	num1Str, _ := num1.ReadString('\n')
	num1Int, _ := strconv.Atoi(strings.TrimSpace(num1Str))

	num2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter second number: ")
	num2Str, _ := num2.ReadString('\n')
	num2Int, _ := strconv.Atoi(strings.TrimSpace(num2Str))


	operation := bufio.NewReader(os.Stdin)
	fmt.Print("Enter operation (+, -, *, /): ")
	operationStr, _ := operation.ReadString('\n')

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("You entered:", num1Int, num2Int, "and operation:", operationStr)

	switch operationStr {
		case "+\n":
			fmt.Println("Performing addition...")
			addition := sum2Num(num1Int, num2Int)
			fmt.Printf("Result: %d\n", addition)
			fmt.Println("----------------------------------------------------------------")
		case "-\n":
			fmt.Println("Performing subtraction...")
			subtraction := sub2Num(num1Int, num2Int)
			fmt.Printf("Result: %d\n", subtraction)
			fmt.Println("----------------------------------------------------------------")
		case "*\n":
			fmt.Println("Performing multiplication...")
			multiplication := mul2Num(num1Int, num2Int)
			fmt.Printf("Result: %d\n", multiplication)
			fmt.Println("----------------------------------------------------------------")
		case "/\n":
			fmt.Println("Performing division...")
			division := div2Num(num1Int, num2Int)
			fmt.Printf("Result: %d\n", division)
			fmt.Println("----------------------------------------------------------------")
		default:
			fmt.Println("Invalid operation")
			fmt.Println("----------------------------------------------------------------")
	}
	
}
