package main

import (
	"fmt"
	"log"
)

func main() {
	nbaPlayers := [5]string{"Lebron James", "Kevin Durant", "Stephen Curry", "Kawhi Leonard", "James Harden"}

	for i, player := range nbaPlayers {
		log.Printf("Address of player %d: %p\n", i, &player)
		fmt.Println(i, player)
	}
	log.Println("nbaPlayers:", nbaPlayers)
	log.Printf("Address of nbaPlayers: %p\n", &nbaPlayers)
	fmt.Println("------------------------------")
	slice1 := nbaPlayers[:]
	slice2 := nbaPlayers[1:3]

	slice2[0] = "Klay Thompson"
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
}
