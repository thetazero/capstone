package main

import "testing"

func Test_getCase(t *testing.T) {
	tests := []struct {
		p    Vector
		q    Vector
		want string
	}{
		{
			MakeIntVector([]int64{3, 1}),
			MakeIntVector([]int64{-1, 2}),
			"i0",
		}, {
			MakeIntVector([]int64{3, 1}),
			MakeIntVector([]int64{-1, 1}),
			"ii",
		}, {
			MakeIntVector([]int64{3, 1}),
			MakeIntVector([]int64{0, -2}),
			"i+",
		}, {
			MakeIntVector([]int64{3, 1}),
			MakeIntVector([]int64{2, -2}),
			"i-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.p.toString()+" "+tt.q.toString(), func(t *testing.T) {
			if got := getCase(tt.p, tt.q); got != tt.want {
				t.Errorf("getCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
