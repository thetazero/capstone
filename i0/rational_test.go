package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestRationalFromContinued(t *testing.T) {
	var tests = []struct {
		input Vector
		want  RationalFunc
	}{
		{
			Vector{RInt(1), RInt(2), RInt(3)}, RationalFunc{
				top: Vector{RInt(0), RInt(4), RInt(0), RInt(6)},
				bot: Vector{RInt(1), RInt(0), RInt(6)},
			},
		},
		{
			Vector{RInt(0), RInt(4), RInt(3), RInt(2)}, RationalFunc{
				top: Vector{RInt(1), RInt(0), RInt(6), RInt(0), RInt(0)},
				bot: Vector{RInt(0), RInt(6), RInt(0), RInt(24)},
			},
		},
		{
			Vector{RInt(0), RInt(1)}, RationalFunc{
				top: Vector{RInt(1), RInt(0), RInt(0)},
				bot: Vector{RInt(0), RInt(1)},
			},
		},
		{
			MakeIntVector([]int64{0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12}), RationalFunc{
				top: MakeIntVector([]int64{1, 0, 756, 0, 83640, 0, 2820960, 0, 30983040, 0, 79833600, 0, 0}),
				bot: MakeIntVector([]int64{0, 39, 0, 9456, 0, 575400, 0, 11659680, 0, 70899840, 0, 79833600}),
			},
		},
		{
			MakeCountingVector(2048), RationalFunc{ //3.65s → 1.7s
				top: Vector{},
				bot: Vector{},
			},
		},
	}
	for _, tt := range tests {
		testname := tt.input.toString()
		t.Run(testname, func(t *testing.T) {
			ans, _ := RationalFromContinued(tt.input)
			if len(tt.want.top) == 0 && len(tt.want.top) == 0 {
				return
			}
			if !ans.equals(tt.want) {
				t.Errorf("got %s/%s, want %s/%s", ans.top, ans.bot, tt.want.top, tt.want.bot)
			}
		})
	}
}

func MakeCountingVector(length int) Vector {
	x := make(Vector, length)
	for i := range x {
		x[i] = big.NewRat(int64(i), 1)
	}
	return x
}

func RInt(a int64) *big.Rat {
	return big.NewRat(a, 1)
}

func TestRationalFromContinuedVector(t *testing.T) {
	tests := []struct {
		name  string
		a     []Vector
		want  RationalFunc
		want1 RationalFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RationalFromContinuedVector(tt.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RationalFromContinuedVector() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RationalFromContinuedVector() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
