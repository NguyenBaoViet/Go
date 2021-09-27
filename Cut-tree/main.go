package main

func main() {

}

func Solution(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == A[i+1] && A[i] == A[i+2] || A[i] < A[i+1] && A[i+1] < A[i+2] && A[i+2] < A[i+3] {
			return -1
		}

	}
	return 0
}
