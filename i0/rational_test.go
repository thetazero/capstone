package main

import (
	"math/big"
	"testing"
)

func TestRationalFromContinued(t *testing.T) {
	var tests = []struct {
		input Vector
		want  Rational
	}{
		{
			Vector{RInt(1), RInt(2), RInt(3)}, Rational{
				top: Vector{RInt(0), RInt(4), RInt(0), RInt(6)},
				bot: Vector{RInt(1), RInt(0), RInt(6), RInt(0)},
			},
		},
		{
			Vector{RInt(0), RInt(4), RInt(3), RInt(2)}, Rational{
				top: Vector{RInt(1), RInt(0), RInt(6), RInt(0), RInt(0)},
				bot: Vector{RInt(0), RInt(6), RInt(0), RInt(24), RInt(0)},
			},
		},
		{
			Vector{RInt(0), RInt(1)}, Rational{
				top: Vector{RInt(1), RInt(0), RInt(0)},
				bot: Vector{RInt(0), RInt(1), RInt(0)},
			},
		},
	}
	for _, tt := range tests {
		testname := tt.input.toString()
		t.Run(testname, func(t *testing.T) {
			ans, _ := RationalFromContinued(tt.input)
			if !ans.equals(tt.want) {
				t.Errorf("got %s/%s, want %s/%s", ans.top.toString(), ans.bot.toString(), tt.want.top.toString(), tt.want.bot.toString())
			}
		})
	}
}

func RInt(a int64) *big.Rat {
	return big.NewRat(a, 1)
}
