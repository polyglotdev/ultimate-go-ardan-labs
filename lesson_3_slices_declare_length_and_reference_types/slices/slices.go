package main

import "fmt"

func main() {
	nflTeams := make([]string, 5)
	nflTeams[0] = "Dallas Cowboys"
	nflTeams[1] = "Philadelphia Eagles"
	nflTeams[2] = "New York Giants"
	nflTeams[3] = "Washington Redskins"
	nflTeams[4] = "Chicago Bears"

	for i, team := range nflTeams {
		fmt.Printf("%d. %s\n", i+1, team)
	}

	inspectSlice(nflTeams)
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length: %d\nCapacity: %d\n", len(slice), cap(slice))
	for i, team := range slice {
		fmt.Printf("%d. %s | %+v\n", i+1, team, &slice[i])
	}
}
