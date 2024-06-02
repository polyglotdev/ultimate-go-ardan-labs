// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	arr1 := [5]string{}

	fmt.Println("arr1", arr1)
	fmt.Println("------------------------------")
	// Declare an array of 5 strings and pre-populate it with names.
	arr2 := [5]string{"Apple", "Orange", "Banana", "Grape", "Plum"}

	// Assign the populated array to the array of zero values.
	arr1 = arr2

	// Iterate over the first array declared.
	for i, fruit := range arr1 {
		fmt.Println(i, &fruit, fruit)
	}
	fmt.Println("------------------------------")
	// Display the string value and address of each element.
	for i, fruit := range arr2 {
		fmt.Println(i, &fruit, fruit)
	}
}
