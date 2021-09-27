package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{4, 0, 1, -2}
	b := []int{-4, 4, 3, 0}
	fmt.Println(Solution(2, 4, a, b))

}
func Solution(inner int, outer int, points_x []int, points_y []int) int {
	count := 0
	for i := 0; i < len(points_x); i++ {
		x := math.Sqrt(math.Pow(float64(points_x[i]), 2) + math.Pow(float64(points_y[i]), 2))
		fmt.Println("x[", i, "]=", x)
		if x > float64(inner) && x < float64(outer) {
			count += 1
		}
	}
	return count
}
