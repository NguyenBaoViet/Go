package main

import "fmt"

func main() {
	age := 22
	year_born := 1999
	year_live := year_born
	for year_live <= year_born+age {
		fmt.Println(year_live)
		year_live++
	}
}
