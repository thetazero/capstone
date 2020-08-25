package main

func computePN(alpha float64, p, q Vector, cap int) []float64 {
	x := make([]float64, cap*2+1)
	ps := p.SizeSquared()
	for n := -cap; n <= cap; n++ {
		top := (ps * (1 + alpha*alpha*ps))
		bottom := q.Add(p.Mul(float64(n))).SizeSquared() * (1 + alpha*alpha*q.Add(p.Mul(float64(n))).SizeSquared())
		x[n+cap] = 1 - top/bottom
	}
	return x
}
