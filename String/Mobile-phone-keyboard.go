package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// count % 2
// 1 = caps lock deactivated, 0 = caps lock activated

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your string:")
		S, _ := reader.ReadString('\n')

		fmt.Println(Solution(S))

	}

}

func Solution(S string) string {
	//
	var count int = 1
	arr := strings.Split(S, "")
	S = ""
	for i := 0; i < len(arr); i++ {
		if arr[i] == "C" {
			arr[i] = "0"
		}
		if arr[i] == "B" {
			arr[i] = "1"
		}
	}

	// C process
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
	//B process
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

	for i := 0; i < len(arr); i++ {
		if arr[i] != "_" {
			S += arr[i]
		}
	}

	return S
}
