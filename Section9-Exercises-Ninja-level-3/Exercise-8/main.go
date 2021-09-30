package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This is not going to be printed")
	case true:
		fmt.Println("This is going to be printed")
	}
}
