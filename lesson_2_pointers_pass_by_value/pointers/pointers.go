package main

import "fmt"

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	fmt.Println("------------------------------")
	fmt.Println("ValueIncrement")
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	valueIncrement(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//
//go:noinline
func increment(inc int) {

	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func valueIncrement(inc *int) {
	*inc++
	println("inc:\tValue Of[", *inc, "]\tAddr Of[", inc, "]")
}
