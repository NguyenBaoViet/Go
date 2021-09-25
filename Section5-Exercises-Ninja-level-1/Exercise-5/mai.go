package main

import "fmt"

type mytype string

var x mytype
var y string

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x := "my string"
	fmt.Println(x)

	y := string(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
