package main

import (
	"math/big"
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
		{
			"test0",
			[]Vector{
				MakeIntVector([]int64{1, 3, 5, 7}),
				MakeIntVector([]int64{1, 2, 4, 6}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{151, 425, 464, 236, 48}),
				bot: MakeIntVector([]int64{115, 252, 188, 48, 0}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{21, 42, 30, 8}),
				bot: MakeIntVector([]int64{16, 22, 8, 0}),
			},
		},
		{
			"test1",
			[]Vector{
				MakeIntVector([]int64{1, 1}),
				MakeIntVector([]int64{1, 1}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{2, 2, 1}),
				bot: MakeIntVector([]int64{1, 1, 0}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{1, 1}),
				bot: MakeIntVector([]int64{1, 0}),
			},
		}, {
			"test2",
			[]Vector{
				MakeIntVector([]int64{1, 3}),
				MakeIntVector([]int64{2, 4}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{4, 10, 8}),
				bot: MakeIntVector([]int64{3, 4, 0}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{1, 2}),
				bot: MakeIntVector([]int64{1, 0}),
			},
		},
		{
			"only x^1 terms test",
			[]Vector{
				MakeIntVector([]int64{0, 0, 0}),
				MakeIntVector([]int64{1, 2, 3}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{0, 4, 0, 6}),
				bot: MakeIntVector([]int64{1, 0, 6, 0}),
			},
			RationalFunc{
				top: MakeIntVector([]int64{1, 0, 2}),
				bot: MakeIntVector([]int64{0, 2, 0}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RationalFromContinuedVector(tt.a)
			if !got.equals(tt.want) {
				t.Errorf("RationalFromContinuedVector() got = %v/%v, want %v/%v", got.top, got.bot, tt.want.top, tt.want.bot)
			}
			if !got1.equals(tt.want1) {
				t.Errorf("RationalFromContinuedVector() got1 = %v/%v, want %v/%v", got1.top, got1.bot, tt.want1.top, tt.want1.bot)
			}
		})
	}
}
