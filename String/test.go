package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var flag bool = false
var count int = 1

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your string:")
		S, _ := reader.ReadString('\n')

		re := regexp.MustCompile(`[ADEFGHIJKLMNOPTUVWXYZ\d\s]`)
		match := re.MatchString(S)
		if match == false {
			fmt.Println("SUCCESS")
		} else {
			fmt.Println("ERROR")
		}
	}

}
