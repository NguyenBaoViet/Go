package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "add",
			args: args{
				s1: "nice",
				s2: "nicer",
			},
			want: "ADD r",
		},
		{
			name: "change",
			args: args{
				s1: "test",
				s2: "tent",
			},
			want: "CHANGE s n",
		},
		{
			name: "move",
			args: args{
				s1: "beans",
				s2: "banes",
			},
			want: "MOVE e",
		},
		{
			name: "impssible",
			args: args{
				s1: "0",
				s2: "odd",
			},
			want: "IMPOSSIBLE",
		},
		{
			name: "nothing",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: "NOTHING",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
