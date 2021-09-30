package main

import "fmt"

func main() {
	yb := 1999

	for {
		if yb > 2021 {
			break
		}
		fmt.Println(yb)
		yb++
	}
}
