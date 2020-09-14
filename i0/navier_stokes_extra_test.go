package main

import (
	"math/big"
	"testing"
)

func Test_ns_pn(t *testing.T) {
	tests := []struct {
		name  string
		p, q  Vector
		debth int
		want  Vector
	}{
		{
			"debth 2 test",
			MakeIntVector([]int64{3, 1}),
			MakeIntVector([]int64{-1, 2}),
			2,
			Vector{big.NewRat(39, 49), big.NewRat(7, 17), big.NewRat(-1, 1), big.NewRat(3, 13), big.NewRat(31, 41)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ns_pn(tt.p, tt.q, tt.debth); !got.equals(tt.want) {
				t.Errorf("ns_pn() = %v, want %v", got, tt.want)
			}
		})
	}
}
