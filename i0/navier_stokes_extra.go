package main

import (
	"fmt"
	"math/big"
)

// Navier-Stokes paper's rohn coefficients
func ns_pn(p, q Vector, debth int64) Vector {
	x := make(Vector, debth*2+1)
	ps := p.SizeSquared()
	one := big.NewRat(1, 1)
	for i := range x {
		//1 - ||p||^2 / ||q+np||^2
		n := big.NewRat(int64(i)-debth, 1)
		sq := big.NewRat(0, 1)
		for j := range q {
			t := new(big.Rat).Add(q[j], new(big.Rat).Mul(p[j], n))
			t.Mul(t, t)
			sq.Add(sq, t)
		}
		x[i] = new(big.Rat).Inv(sq)
		x[i].Mul(x[i], ps)
		x[i].Sub(one, x[i])
	}
	return x
}

//ns_an determines the aₙ terms for the navier stokes case. The vectors returned represent λ coefficient followed by the constant coefficient for all aₙ.
func ns_an(ν *big.Rat, p, q, pn Vector) []Vector {
	lamdas := make(Vector, len(pn))
	constants := make(Vector, len(pn))
	for i := range lamdas {
		// (λ + ν||q+np||^2)/pn
		n := big.NewRat(int64(i-len(pn)/2), 1)
		lamdas[i] = new(big.Rat).Inv(pn[i])
		constants[i] = big.NewRat(0, 1)
		for j := range p {
			t := new(big.Rat).Add(q[j], new(big.Rat).Mul(p[j], n))
			t.Mul(t, t)
			constants[i].Add(constants[i], t)
		}
		constants[i].Mul(constants[i], ν)
		constants[i].Quo(constants[i], pn[i])
	}
	return []Vector{constants, lamdas}
}

// returns f, g, fp1.top, gp1.top, p0

func solve_ns(ν *big.Rat, p, q Vector, debth int64) (RationalFunc, RationalFunc, Vector, Vector, Vector) {
	debth++
	pn := ns_pn(p, q, debth)
	coeff := ns_an(ν, p, q, pn)
	fmt.Println(coeff)
	positiveCoeff := make([]Vector, len(coeff))
	negativeCoeff := make([]Vector, len(coeff))
	for j := range coeff {
		positiveCoeff[j] = coeff[j][debth+1:]
		positiveCoeff[j] = append(Vector{big.NewRat(0, 1)}, positiveCoeff[j]...)
		negativeCoeff[j] = coeff[j][:debth]
		for i := 0; i < len(negativeCoeff[j])/2; i++ {
			negativeCoeff[j][i], negativeCoeff[j][len(negativeCoeff[j])-1-i] = negativeCoeff[j][len(negativeCoeff[j])-1-i], negativeCoeff[j][i]
		}
		negativeCoeff[j] = append(Vector{big.NewRat(0, 1)}, negativeCoeff[j]...)

	}
	fp1, f := RationalFromContinuedVector(positiveCoeff)
	// fmt.Println(f)
	gp1, g := RationalFromContinuedVector(negativeCoeff)
	return f, g, fp1.bot, gp1.bot, Vector{coeff[0][debth], coeff[1][debth]}
}
