package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{100, 200, 100}
	b := []int{50, 100, 100}
	fmt.Println(Solution(a, b, 100, 100))

}

func Solution(A []int, B []int, X int, Y int) int {
	for i := 0; i < len(A); i++ {
		x := math.Sqrt(math.Pow(float64(A[i]-X), 2) + math.Pow(float64(int(B[i])-Y), 2))
		if x <= 20 {
			return i
		}
	}
	return -1
}
