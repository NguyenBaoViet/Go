package main

import "fmt"

func main() {
	s := struct {
		first    string
		friends  map[string]int
		favorite []string
	}{
		first: "viet",
		friends: map[string]int{
			"thanh": 1234,
			"hoang": 4321,
			"bao":   34555,
		},
		favorite: []string{"ice cream", "game"},
	}

	fmt.Println(s)

	fmt.Println("Name:", s.first)
	fmt.Println("Friends:")
	for k, v := range s.friends {
		fmt.Println("\t", k, v)
	}
	fmt.Println("Favorite:")
	for i, v := range s.favorite {
		fmt.Println("\t", i+1, v)
	}
}
