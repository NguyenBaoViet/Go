package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "Bond" {
		fmt.Println(`Missing "James"`)
	} else {
		fmt.Println("You got wrong string")
	}
}
