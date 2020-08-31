package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPolynomialMul(t *testing.T) {
	var tests = []struct {
		input [2]Vector
		want  Vector
	}{
		{
			[2]Vector{Vector{RInt(1), RInt(2)}, Vector{RInt(1), RInt(2)}},
			Vector{RInt(1), RInt(4), RInt(4)},
		}, {
			[2]Vector{Vector{RInt(1), RInt(4), RInt(3)}, Vector{RInt(1), RInt(7), RInt(2)}},
			Vector{RInt(1), RInt(11), RInt(33), RInt(29), RInt(6)},
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

func TestVector_SizeSquared(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want *big.Rat
	}{
		{
			"<1,2>",
			Vector{RInt(1), RInt(2)},
			RInt(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.SizeSquared(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.SizeSquared() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Mul(t *testing.T) {
	tests := []struct {
		v    Vector
		k    *big.Rat
		want Vector
	}{
		{
			Vector{RInt(1), big.NewRat(2, 1), big.NewRat(3, 4)},
			big.NewRat(3, 8),
			Vector{big.NewRat(3, 8), big.NewRat(3, 4), big.NewRat(9, 32)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.v.toString(), func(t *testing.T) {
			got := tt.v.Mul(tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s.Mul(%s) = %s, want %s", tt.v.toString(), tt.k.FloatString(10), got.toString(), tt.want.toString())
			}
		})
	}
}

func TestVector_ShiftRight(t *testing.T) {
	tests := []struct {
		v    Vector
		want Vector
	}{
		{
			Vector{RInt(1), RInt(2), RInt(3)},
			Vector{RInt(0), RInt(1), RInt(2), RInt(3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.v.toString(), func(t *testing.T) {
			got := tt.v.ShiftRight()
			if !got.equals(tt.want) {
				t.Errorf("Vector.ShiftRight() = %s, want %s", got.toString(), tt.want.toString())
			}
		})
	}
}

func TestVector_Add(t *testing.T) {
	tests := []struct {
		a, b Vector
		want Vector
	}{
		{
			MakeIntVector([]int64{1, 2, 3}),
			MakeIntVector([]int64{2, 3, 4}),
			MakeIntVector([]int64{3, 5, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.a.toString()+"+"+tt.b.toString(), func(t *testing.T) {
			got := tt.a.Add(tt.b)
			if !got.equals(tt.want) {
				t.Errorf("%s+%s = %s, want %s", tt.a.toString(), tt.b.toString(), got.toString(), tt.want.toString())
			}
		})
	}
}

func MakeIntVector(a []int64) Vector {
	x := make(Vector, len(a))
	for i := range x {
		x[i] = RInt(a[i])
	}
	return x
}
