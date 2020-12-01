package main

import "math/big"

func computePN(alpha *big.Rat, p, q Vector, cap int64) Vector {
	x := make(Vector, cap*2+1)
	// top := (ps * (1 + alpha*alpha*ps))
	one := big.NewRat(1, 1)
	alpha2 := new(big.Rat).Mul(alpha, alpha)
	top := new(big.Rat).Mul(alpha2, p.SizeSquared())
	top.Add(top, one)
	top.Mul(top, p.SizeSquared())

	bottom := new(big.Rat)
	for n := -cap; n <= cap; n++ {
		// bottom := q.Add(p.Mul(float64(n))).SizeSquared() * (1 + alpha*alpha*q.Add(p.Mul(float64(n))).SizeSquared())
		nn := big.NewRat(n, 1)

		qnp := big.NewRat(0, 1)
		t := new(big.Rat)
		for i := range p {
			t.Mul(nn, p[i])
			t.Add(t, q[i])
			t.Mul(t, t)
			qnp.Add(qnp, t)
		}
		bottom.Mul(alpha2, qnp)
		bottom.Add(bottom, one)
		bottom.Mul(bottom, qnp)
		// x[n+cap] = 1 - top/bottom
		x[n+cap] = new(big.Rat).Sub(one, bottom.Quo(top, bottom))
	}
	return x
}
