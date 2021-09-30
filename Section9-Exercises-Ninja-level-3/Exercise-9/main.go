package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Jame Bond"

	switch x {
	case "Jame Bond":
		fmt.Println(x)
	case "jame bond":
		fmt.Println(strings.ToUpper(x))
	case "JAME BOND":
		fmt.Println(strings.ToLower(x))

	}
}
