package main

import (
	"math"
	"testing"
)

func TestVectorSize(t *testing.T) {
	var tests = []struct {
		input Vector
		want  float64
	}{
		{
			Vector{1, 2}, math.Pow(5, 0.5),
		},
		{
			Vector{-3, -4}, 5.0,
		},
		{
			Vector{0, 0}, 0,
		},
	}
	for _, tt := range tests {
		testname := tt.input.toString()
		t.Run(testname, func(t *testing.T) {
			ans := tt.input.size()
			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}

func TestPolynomialMul(t *testing.T) {
	var tests = []struct {
		input [2]Vector
		want  Vector
	}{
		{
			[2]Vector{Vector{1, 2}, Vector{1, 2}},
			Vector{1, 4, 4},
		}, {
			[2]Vector{Vector{1, 4, 3}, Vector{1, 7, 2}},
			Vector{1, 11, 33, 29, 6},
		},
	}
	for _, tt := range tests {
		testname := tt.input[0].toString() + "x" + tt.input[0].toString()
		t.Run(testname, func(t *testing.T) {
			ans := tt.input[0].PolynomialMul(tt.input[1])
			if !ans.equals(tt.want) {
				t.Errorf("got %s, want %s", ans.toString(), tt.want.toString())
			}
		})
	}
}
