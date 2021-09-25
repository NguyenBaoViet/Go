package main

import (
	"fmt"
	"strings"
)

var count int = 1

func main() {
	S := "CgCoodlBClCuck"
	arr := strings.Split(S, "")
	S = ""
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == "C" {
			arr[i] = "0"
		}
		if arr[i] == "B" {
			arr[i] = "1"
		}
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == "0" {
			count += 1
			arr[i] = "_"
			if count%2 == 0 {
				for j := i + 1; j < len(arr); j++ {
					if arr[j] != "0" {
						arr[j] = strings.ToUpper(arr[j])
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] == "1" {
			arr[i] = "_"
			for j := i - 1; j > -1; j-- {
				if arr[j] != "_" {
					arr[j] = "_"
					break
				}
			}
		}
	}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != "_" {
			S += arr[i]
		}
	}
	fmt.Print(S)
}
