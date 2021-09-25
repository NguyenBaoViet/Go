package main

import "fmt"

var y = 43

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
}

func foo() {
	fmt.Println("again:", y)
}
