package main

import "fmt"

func main() {
	x := [5]int{21, 54, 8, 5, 4}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	fmt.Printf("%T", x)
}
