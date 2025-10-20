package main

import (
	"fmt"

	"github.com/dytsou/calc/internal/calc"
)

func main() {
	fmt.Println("Calculator")

	var a, b int
	var op string
	fmt.Println("Enter a op(+,-,*,/) b")
	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	switch op {
	case "+":
		fmt.Println("Result:", calc.Add(a, b))
	case "-":
		fmt.Println("Result:", calc.Subtract(a, b))
	case "*":
		fmt.Println("Result:", calc.Multiply(a, b))
	case "/":
		result, err := calc.Divide(a, b)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Result: %.2f\n", result)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
