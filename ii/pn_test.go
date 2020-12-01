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
		{
			args{
				big.NewRat(3, 4),
				MakeIntVector([]int64{6, 2}),
				MakeIntVector([]int64{-3, 4}),
				3,
			},
			Vector{big.NewRat(354861, 357869), big.NewRat(88837, 91845), big.NewRat(10269, 13277), big.NewRat(-1803, 1205), big.NewRat(781, 3789), big.NewRat(35301, 38309), big.NewRat(188157, 191165)},
		},
		{
			args{
				big.NewRat(3, 4),
				MakeIntVector([]int64{6, 2}),
				MakeIntVector([]int64{-3, 4}),
				100000,
			},
			Vector{}, //hacky speed test used to take 0.9s, now 0.45s
		},
	}
	for _, tt := range tests {
		t.Run("⍺:"+tt.args.alpha.FloatString(5)+"p:"+tt.args.p.toString()+"q:"+tt.args.q.toString(), func(t *testing.T) {
			got := computePN(tt.args.alpha, tt.args.p, tt.args.q, tt.args.cap)
			if len(tt.want) == 0 {
				return
			}
			if !got.equals(tt.want) {
				t.Errorf("computePN() = %v, want %v", got, tt.want)
			}
		})
	}
}
