package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution("test", "tent"))
}

func Solution(s1 string, s2 string) string {
	rs := ""
	if s1 == s2 {
		rs = "NOTHING"
	} else if len(s1) == len(s2) {
		flag := false
		for _, char := range s2 {
			if strings.Contains(s1, string(char)) {
				if strings.Index(s2, string(char))-strings.Index(s1, string(char)) > -2 {
					flag = true
					if strings.Index(s2, string(char))-strings.Index(s1, string(char)) > 0 {
						rs = "MOVE " + string(char)
					}
				}
			} else {
				flag = false
				break
			}
		}
		if !flag {
			count := 0
			for i := 0; i < len(s1); i++ {
				if s2[i] != s1[i] {
					count += 1
					rs = "CHANGE " + string(s1[i]) + " " + string(s2[i])
				}
				if count > 1 {
					rs = "IMPOSSIBLE"
					break
				}
			}
		}
	} else if len(s1)-len(s2) == -1 {
		if s1 == s2[:len(s2)-1] {
			rs = "ADD " + string(s2[len(s2)-1])
		}
	} else {
		rs = "IMPOSSIBLE"
	}

	return rs
}
