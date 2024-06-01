package main

import (
	"fmt"
	"log"
)

// User represents a user in the system.
type User struct {
	Name   string
	Email  string
	Logins int
}

func main() {
	// Declare and initialize a variable named elijah of type User.
	elijah := User{
		Name:  "Elijah",
		Email: "elijah@ardanlabs.com",
	}

	// Display the initial state of elijah
	display(&elijah)

	// Increment the logins field of elijah
	increment(&elijah.Logins)

	// Display the state of elijah after incrementing logins
	display(&elijah)
	fmt.Println("------------------------------")
	ezra := User{
		Name:  "Ezra",
		Email: "ezra@ardanlabs.com",
	}

	display(&ezra)
	increment(&ezra.Logins)
	increment(&ezra.Logins)
}

// increment declares logins as a pointer variable whose value is
// always an address and points to values of type int.
func increment(logins *int) {
	*logins++
	log.Printf("Incremented logins to %d\n", *logins)
}

// display declares u as User pointer variable whose value is always an address
// and points to values of type User.
func display(u *User) {
	log.Printf("Displaying user: %+v\n", *u)
}
