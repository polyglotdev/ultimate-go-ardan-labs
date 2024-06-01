package main

import "fmt"

func main() {

	value := 20

	fmt.Println("Address of value:", &value, "\nValue of value:", value)
	fmt.Println("------------------------------")
	var ptrVariable *int = &value
	fmt.Println("Address of ptrVariable:", &ptrVariable, "\nValue of ptrVariable:", ptrVariable)

	fmt.Println("------------------------------")
	fmt.Println("value == *ptrVariable:", value == *ptrVariable)

	fmt.Println("------------------------------")
	p := &value
	fmt.Println("Address of p:", &p, "\nValue of p:", p, "\nPoints to:", *p)
	fmt.Println("------------------------------")
	if value == *ptrVariable && value == *p {
		fmt.Println("value, *ptrVariable, and *p have the same value.")
	} else {
		fmt.Println("value, *ptrVariable, and *p do not have the same value.")
	}
}
