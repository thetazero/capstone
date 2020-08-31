package main

import (
	"math/big"
	"testing"
)

func Test_computePN(t *testing.T) {
	type args struct {
		alpha *big.Rat
		p     Vector
		q     Vector
		cap   int64
	}
	tests := []struct {
		args args
		want Vector
	}{
		{
			args{
				RInt(1),
				MakeIntVector([]int64{3, 1}),
				MakeIntVector([]int64{-2, 1}),
				0,
			},
			Vector{big.NewRat(-8, 3)},
		},
	}
	for _, tt := range tests {
		t.Run("⍺:"+tt.args.alpha.FloatString(5)+"p:"+tt.args.p.toString()+"q:"+tt.args.q.toString(), func(t *testing.T) {
			got := computePN(tt.args.alpha, tt.args.p, tt.args.q, tt.args.cap)
			if !got.equals(tt.want) {
				t.Errorf("computePN() = %v, want %v", got, tt.want)
			}
		})
	}
}
