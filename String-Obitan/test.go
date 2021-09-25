package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("nice", "nicer"))
}

func Solution(s1 string, s2 string) string {
	rs := ""
	if s1 == s2 {
		rs = "NOTHING"
	} else {
		if len(s1) == len(s2) {

		}
		if len(s1) < len(s2) {
			if s1 == s2[:len(s2)-1] {
				rs = "ADD " + string(s2[len(s2)-1])
			}
		}
	}

	return rs
}
