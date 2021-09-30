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

	m := map[string]person{
		p1.first: p1,
	}

	fmt.Println(m)

	fmt.Println("First name:", m[p1.first].first)
	fmt.Println("Last name:", m[p1.first].last)
	fmt.Println("Favorite:")
	for i, v := range m[p1.first].favorite {
		fmt.Println("\t", i+1, v)
	}
}
