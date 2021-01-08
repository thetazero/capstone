package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(-1, 1), big.NewRat(1, 1)}, Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, 101)
	f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(2, 1), big.NewRat(0, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(-1, 1)}, 120)

	// f := MakeEulerEquation(big.NewRat(1, 1), Vector{big.NewRat(3, 1), big.NewRat(1, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(2, 1)}, 10)

	// f := MakeEulerEquation(big.NewRat(0, 1), Vector{big.NewRat(2, 1), big.NewRat(0, 1)}, Vector{big.NewRat(-1, 1), big.NewRat(-1, 1)}, 120)

	point := Complex{big.NewRat(1, 10), big.NewRat(0, 1)}
	// point := Complex{big.NewRat(1538, 1000), big.NewRat(0, 1)}
	fmt.Println("point:", point[0].FloatString(3), point[1].FloatString(3))
	aboutzero := f(point)
	fmt.Println("val:", aboutzero[0].FloatString(3), aboutzero[1].FloatString(3))

	p := Path(func(i, j int) Complex {
		a, b, theta := 0.5, 10.0, math.Pi/20
		acircle, bcircle, line := a*(math.Pi-2*theta), b*(math.Pi-2*theta), (b - a)
		len := acircle + bcircle + 2*line
		position := float64(i) / float64(j) * len
		if position < acircle { //inner circle
			position = acircle - position
			offset := -math.Pi/2 + theta
			real, imag := a*math.Cos(position/a+offset), a*math.Sin(position/a+offset)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= acircle
		if position < line { //lower line
			t := position / line
			angle := math.Pi/2 - theta
			real := t*b*math.Cos(angle) + (1-t)*a*math.Cos(angle)
			imag := t*-b*math.Sin(angle) + (1-t)*-a*math.Sin(angle)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= line
		if position < bcircle { //outer circle
			offset := -math.Pi/2 + theta
			real, imag := b*math.Cos(position/b+offset), b*math.Sin(position/b+offset)
			return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
		}
		position -= bcircle
		//upper line
		t := 1 - position/line
		angle := math.Pi/2 - theta
		real := t*b*math.Cos(angle) + (1-t)*a*math.Cos(angle)
		imag := t*b*math.Sin(angle) + (1-t)*a*math.Sin(angle)
		return Complex{new(big.Rat).SetFloat64(real), new(big.Rat).SetFloat64(imag)}
	})

	samples := 1000
	p.Draw(f, samples, "_result.png")
	p.Draw(func(x Complex) Complex {
		return x
	}, samples, "_contour.png")

}

func getCase(p, q Vector) string {
	r2 := p.SizeSquared()
	qpp := p.Add(q).SizeSquared()
	qmp := p.Add(q.Mul(big.NewRat(-1, 1))).SizeSquared()
	if qpp.Cmp(r2) == 1 && qmp.Cmp(r2) == 1 {
		return "i0"
	} else if qpp.Cmp(r2) == -1 && qmp.Cmp(r2) == 1 {
		return "ii"
	} else if qpp.Cmp(r2) == 0 && qmp.Cmp(r2) == 1 {
		return "i+"
	} else if qpp.Cmp(r2) == 1 && qmp.Cmp(r2) == 0 {
		return "i-"
	}
	return ""
}
