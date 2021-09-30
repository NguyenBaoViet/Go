package main

import "fmt"

func main() {
	a := (40 == 42)
	b := (40 <= 42)
	c := (40 >= 42)
	d := (40 != 42)
	e := (40 < 42)
	f := (40 > 42)

	fmt.Println("40 == 42 ", a)
	fmt.Println("40 <= 42 ", b)
	fmt.Println("40 >= 42 ", c)
	fmt.Println("40 != 42 ", d)
	fmt.Println("40 < 42 ", e)
	fmt.Println("40 > 42 ", f)
}
