package main

import "fmt"

func main() {
	x := []string{`Alabama`, `Alaska`, `Arizone`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
