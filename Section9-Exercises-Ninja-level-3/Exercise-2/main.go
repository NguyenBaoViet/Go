package main

import "fmt"

func main() {
	for i := 1; i < 27; i++ {
		fmt.Println(i)
		for j := 0; j < 4; j++ {
			fmt.Printf("\t%#U\n", i+64)
		}
	}
}
