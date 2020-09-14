package main

import "math/big"

// Navier-Stokes paper's rohn coefficients
func ns_pn(p, q Vector, debth int) Vector {
	x := make(Vector, debth*2+1)
	ps := p.SizeSquared()
	one := big.NewRat(1, 1)
	for i := range x {
		//1 - ||p||^2 / ||q+np||^2
		n := big.NewRat(int64(i-debth), 1)
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
func ns_an(λ, ν *big.Rat, p, q, pn Vector) []Vector {
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
	return []Vector{lamdas, constants}
}
