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
	c := getCase(p, q)
	for i := range x {
		//1 - ||p||^2 / ||q+np||^2
		n := big.NewRat(int64(i)-debth, 1)
		if (c == "i-" && int64(i) < debth) || (c == "i+" && int64(i) > debth) {
			x[i] = big.NewRat(42, 1)
			continue
		}
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
	fmt.Println(x)
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

func solve_ns(ν *big.Rat, p, q Vector, debth int64) (RationalFunc, RationalFunc, Vector, Vector, Vector, string) {
	c := getCase(p, q)
	if c == "i0" {
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
		gp1, g := RationalFromContinuedVector(negativeCoeff)
		// computer := func(x *big.Rat) *big.Rat {
		// 	a := f.Compute(x)
		// 	a.Add(a, g.Compute(x))
		// 	a.Add(a, Vector{coeff[0][debth], coeff[1][debth]}.Compute(x))
		// 	return a
		// }
		// yvals := "["
		// xvals := "["
		// steps := int64(1000)
		// lower, upper := int64(-19), int64(-10)
		// for i := 0; i < 1000; i++ {
		// 	if i != 0 {
		// 		yvals += ","
		// 		xvals += ","
		// 	}
		// 	x := big.NewRat(int64(i), steps/(upper-lower))
		// 	x.Add(x, big.NewRat(lower, 1))
		// 	yvals += computer(x).FloatString(5)
		// 	xvals += x.FloatString(5)
		// }
		// yvals += "]"
		// xvals += "]"
		// fmt.Println(xvals)
		// fmt.Println("------")
		// fmt.Println(yvals)
		// fmt.Println(computer(big.NewRat(0, 1)).FloatString(5))
		return f, g, fp1.bot, gp1.bot, Vector{coeff[0][debth], coeff[1][debth]}, c
	} else if c == "i+" {
		debth++
		pn := ns_pn(p, q, debth)
		coeff := ns_an(ν, p, q, pn)
		fmt.Println(coeff)
		negativeCoeff := make([]Vector, len(coeff))
		for j := range coeff {
			negativeCoeff[j] = coeff[j][:debth]
			for i := 0; i < len(negativeCoeff[j])/2; i++ {
				negativeCoeff[j][i], negativeCoeff[j][len(negativeCoeff[j])-1-i] = negativeCoeff[j][len(negativeCoeff[j])-1-i], negativeCoeff[j][i]
			}
			negativeCoeff[j] = append(Vector{big.NewRat(0, 1)}, negativeCoeff[j]...)

		}
		gp1, g := RationalFromContinuedVector(negativeCoeff)
		return RationalFunc{}, g, Vector{}, gp1.bot, Vector{coeff[0][debth], coeff[1][debth]}, c
	} else if c == "i-" {
		debth++
		pn := ns_pn(p, q, debth)
		fmt.Println("pn for i-,", pn)
		coeff := ns_an(ν, p, q, pn)
		fmt.Println(coeff)
		positiveCoeff := make([]Vector, len(coeff))
		for j := range coeff {
			positiveCoeff[j] = coeff[j][debth+1:]
			positiveCoeff[j] = append(Vector{big.NewRat(0, 1)}, positiveCoeff[j]...)
		}
		fp1, f := RationalFromContinuedVector(positiveCoeff)
		return f, RationalFunc{}, fp1.bot, Vector{}, Vector{coeff[0][debth], coeff[1][debth]}, c
	}
	return RationalFunc{}, RationalFunc{}, Vector{}, Vector{}, Vector{}, "error"
}
