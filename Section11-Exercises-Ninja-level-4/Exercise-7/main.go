package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stired"}
	y := []string{"Miss", "Monneypenny", "Heloooo, James."}

	z := [][]string{x, y}

	for i, v := range z {
		fmt.Println(i)
		fmt.Println("\t", v)
		for j, k := range v {
			fmt.Println("\t\t", j, k)
		}
	}
}
