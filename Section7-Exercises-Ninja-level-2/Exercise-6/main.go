package main

import "fmt"

const (
	a = 2021 - iota
	b = 2021 - iota
	c = 2021 - iota
	d = 2021 - iota
)

func main() {
	fmt.Print(a, b, c, d)
}
