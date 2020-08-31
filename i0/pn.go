package main

import "math/big"

func computePN(alpha *big.Rat, p, q Vector, cap int64) Vector {
	x := make(Vector, cap*2+1)
	ps := p.SizeSquared()
	for n := -cap; n <= cap; n++ {
		// top := (ps * (1 + alpha*alpha*ps))
		top := new(big.Rat).Mul(alpha, alpha)
		top.Mul(top, ps)
		top.Add(top, big.NewRat(1, 1))
		top.Mul(top, ps)
		// bottom := q.Add(p.Mul(float64(n))).SizeSquared() * (1 + alpha*alpha*q.Add(p.Mul(float64(n))).SizeSquared())
		bottom := q.Add(p.Mul(big.NewRat(n, 1))).SizeSquared()
		temp := new(big.Rat).Mul(alpha, alpha)
		temp.Mul(temp, q.Add(p.Mul(big.NewRat(n, 1))).SizeSquared())
		temp.Add(big.NewRat(1, 1), temp)
		bottom.Mul(bottom, temp)
		bottom.Inv(bottom)
		// x[n+cap] = 1 - top/bottom
		x[n+cap] = new(big.Rat).Sub(big.NewRat(1, 1), top.Mul(top, bottom))
	}
	return x
}
