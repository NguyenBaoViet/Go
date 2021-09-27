package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		B []int
		X int
		Y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: " A = [100, 200, 100], B = [50, 100, 100], X = 100 and Y = 100",
			args: args{
				A: []int{100, 200, 100},
				B: []int{50, 100, 100},
				X: 100,
				Y: 100,
			},
			want: 2,
		},
		{
			name: " A = [100, 200, 100], B = [50, 100, 100], X = 100 and Y = 70",
			args: args{
				A: []int{100, 200, 100},
				B: []int{50, 100, 100},
				X: 100,
				Y: 70,
			},
			want: 0,
		},
		{
			name: " A = [100, 200, 100], B = [50, 100, 100], X = 200 and Y = 60",
			args: args{
				A: []int{100, 200, 100},
				B: []int{50, 100, 100},
				X: 200,
				Y: 60,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B, tt.args.X, tt.args.Y); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
