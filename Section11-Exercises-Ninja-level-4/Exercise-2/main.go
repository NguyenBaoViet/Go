package main

import "fmt"

func main() {
	x := []int{56, 45, 85, 21, 5, 51, 25, 14, 58, 48}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
