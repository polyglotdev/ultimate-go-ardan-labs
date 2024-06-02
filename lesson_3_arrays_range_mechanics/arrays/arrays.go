// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the pointer semantic form of the for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Original value of friends[1]: %s\n", friends[1])

	for i := range friends {
		friends[1] = "Elijah"
		friends[2] = "Ezra"

		if i == 1 {
			fmt.Printf("Altered value of friends[1]: %s\n", friends[1])
		}
		if i == 2 {
			fmt.Printf("Altered value of friends[2]: %s\n", friends[2])
		}
	}

	fmt.Println("friends:", friends)
	fmt.Println("------------------------------")

	// Using the value semantic form of the for range.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
}
