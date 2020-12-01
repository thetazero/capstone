package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestComplex_Add(t *testing.T) {
	type args struct {
		a Complex
		b Complex
	}
	tests := []struct {
		name string
		c    *Complex
		args args
		want Complex
	}{
		{
			"(1+i)+(1+i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 1), big.NewRat(1, 1)},
				Complex{big.NewRat(1, 1), big.NewRat(1, 1)},
			},
			Complex{big.NewRat(2, 1), big.NewRat(2, 1)},
		}, {
			"(1-i)+(2+i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 1), big.NewRat(-1, 1)},
				Complex{big.NewRat(2, 1), big.NewRat(1, 1)},
			},
			Complex{big.NewRat(3, 1), big.NewRat(0, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Add(tt.args.a, tt.args.b); !got.Equals(tt.want) {
				t.Errorf("Complex.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex_Sub(t *testing.T) {
	type args struct {
		a Complex
		b Complex
	}
	tests := []struct {
		name string
		c    *Complex
		args args
		want Complex
	}{
		{
			"1+i-(3-2i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 1), big.NewRat(1, 1)},
				Complex{big.NewRat(3, 1), big.NewRat(-2, 1)},
			},
			Complex{big.NewRat(-2, 1), big.NewRat(3, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex_Mul(t *testing.T) {
	type args struct {
		a Complex
		b Complex
	}
	tests := []struct {
		name string
		c    *Complex
		args args
		want Complex
	}{
		{
			"(1+2i)*(3-4i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 1), big.NewRat(2, 1)},
				Complex{big.NewRat(3, 1), big.NewRat(-4, 1)},
			},
			Complex{big.NewRat(11, 1), big.NewRat(2, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Mul(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex_Quo(t *testing.T) {
	type args struct {
		a Complex
		b Complex
	}
	tests := []struct {
		name string
		c    *Complex
		args args
		want Complex
	}{
		{
			"(1+2i)/(3+4i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 1), big.NewRat(2, 1)},
				Complex{big.NewRat(3, 1), big.NewRat(4, 1)},
			},
			Complex{big.NewRat(11, 25), big.NewRat(2, 25)},
		}, {
			"(1/2-3i)/(7/9+2i)",
			&Complex{},
			args{
				Complex{big.NewRat(1, 2), big.NewRat(-3, 1)},
				Complex{big.NewRat(7, 9), big.NewRat(2, 1)},
			},
			Complex{big.NewRat(-909, 746), big.NewRat(-270, 373)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Quo(tt.args.a, tt.args.b)
			if !got.Equals(tt.want) {
				t.Errorf("Complex.Quo() = %v, want %v", got, tt.want)
			}
		})
	}
}
