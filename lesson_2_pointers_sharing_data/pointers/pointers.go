package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	a := 10
	b := 0

	// Perform a math operation.
	addResult, _ := maths("add", &a, &b)
	fmt.Println("Add:", addResult)

	subResult, _ := maths("sub", &a, &b)
	fmt.Println("Subtract:", subResult)

	mulResult, _ := maths("mul", &a, &b)
	fmt.Println("Multiply:", mulResult)

	divResult, err := maths("div", &a, &b)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Divide:", divResult)
	}
}

// maths take in the sort of math operation to perform and two integers
// and returns the result of the operation and an error if any.
func maths(op string, a, b *int) (int, error) {
	switch op {
	case "add":
		log.Println("Performing addition")
		return *a + *b, nil
	case "sub":
		log.Println("Performing subtraction")
		return *a - *b, nil
	case "mul":
		log.Println("Performing multiplication")
		return *a * *b, nil
	case "div":
		if *b != 0 {
			log.Println("Performing division")
			return *a / *b, nil
		} else {
			// Handle division by zero if necessary
			log.Println("Error: Division by zero")
			return 0, errors.New("division by zero")
		}
	default:
		log.Println("Error: Invalid operation")
		return 0, errors.New("invalid operation")
	}
}
