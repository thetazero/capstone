package main

import "testing"

func TestRationalFromContinued(t *testing.T) {
	var tests = []struct {
		input Vector
		want  Rational
	}{
		{
			Vector{1, 2, 3}, Rational{
				top: Vector{0, 4, 0, 6},
				bot: Vector{1, 0, 6, 0},
			},
		},
		{
			Vector{0, 4, 3, 2}, Rational{
				top: Vector{1, 0, 6, 0, 0},
				bot: Vector{0, 6, 0, 24, 0},
			},
		},
		{
			Vector{0, 1}, Rational{
				top: Vector{1, 0, 0},
				bot: Vector{0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		testname := tt.input.toString()
		t.Run(testname, func(t *testing.T) {
			ans := RationalFromContinued(tt.input)
			if !ans.equals(tt.want) {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}
