package main

import "fmt"

func main() {
	nbaPlayers := []string{"Lebron James", "Kevin Durant", "Stephen Curry", "Kawhi Leonard", "James Harden"}
	inspectSlice(nbaPlayers)
	fmt.Println(`----------------------------------`)

	// append a new player to the slice
	nbaPlayers = append(nbaPlayers, "Micheal Jordan")
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	nbaPlayers = append(nbaPlayers, "Magic Johnson")
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	nbaPlayers = append(nbaPlayers, "Larry Bird")
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	nbaPlayers = append(nbaPlayers, "Shaquille O'Neal")
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	nbaPlayers = append(nbaPlayers, "Kobe Bryant")
	inspectSlice(nbaPlayers)

	fmt.Println(`----------------------------------`)
	nbaPlayers = append(nbaPlayers, "Tim Duncan")
	inspectSlice(nbaPlayers)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length: %d\nCapacity: %d\n", len(slice), cap(slice))
	for i, player := range slice {
		fmt.Printf("%d. %s | %+v\n", i+1, player, &slice[i])
	}
}
