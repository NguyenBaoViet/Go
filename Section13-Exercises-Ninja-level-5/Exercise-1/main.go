package main

import "fmt"

type person struct {
	first    string
	last     string
	favorite []string
}

func main() {
	p1 := person{
		first:    "Viet",
		last:     "Nguyen Bao",
		favorite: []string{"ice cream", "game"},
	}

	fmt.Println("First name:", p1.first)
	fmt.Println("Last name:", p1.last)
	fmt.Println("Favorite:")
	for i, v := range p1.favorite {
		fmt.Println("\t", i+1, v)
	}
}
