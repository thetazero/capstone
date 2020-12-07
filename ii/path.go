package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/cheggaaa/pb"
	"github.com/lucasb-eyer/go-colorful"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

//Path defines a complex path. First argument is the step and Second is the total number of steps for a cycle. Prefer (relativly) equidistant samples.
type Path func(int, int) Complex

//Draw result of going along the path
func (p Path) Draw(f EulerEquation, path string) {
	plt, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(1)
	plt.X.Label.Text = "Real"
	plt.Y.Label.Text = "Imaginary"
	// plt.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	// 	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	// })
	// plt.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	// 	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	// })

	j := 100
	pts := make(plotter.XYs, j)
	bar := pb.StartNew(j)
	for i := 0; i < j; i++ {
		val := f(p(i, j))
		// val := p(i, j)

		real, _ := val[0].Float64()
		img, _ := val[1].Float64()
		pts[i] = plotter.XY{X: real, Y: img}
		if i == 0 {
			fmt.Println(pts[i])
		}
		// pts = append(pts, plotter.XY{X: 1, Y: 1})
		bar.Increment()
	}
	bar.Finish()
	// fmt.Println(pts)
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	colors := moreland.Kindlmann() // Initialize a color map.
	colors.SetMax(float64(j))
	colors.SetMin(0)
	scatter.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		hue := float64(i) / float64(j) * 360
		c := colorful.Hsv(hue, 1, 1)
		if i == 0 {
			c = colorful.Hsv(hue, 0, 0)
		}
		return draw.GlyphStyle{Color: c, Radius: vg.Points(3), Shape: draw.CircleGlyph{}}
	}
	// fmt.Println(scatter.DataRange())
	plt.Add(scatter)

	// plotter.DefaultGlyphStyle.Radius = 2
	// plotter.DefaultGlyphStyle.Color = color.RGBA{0, 100, 100, 0xff}
	scatter, err = plotter.NewScatter(plotter.XYs{{X: 0, Y: 0}})
	if err != nil {
		log.Panic(err)
	}
	plt.Add(scatter)

	err = plt.Save(500, 500, path)
	if err != nil {
		log.Panic(err)
	}
}

func circle(i, j int, r, x, y *big.Rat) Complex {
	real := new(big.Rat).SetFloat64(math.Cos(2 * math.Pi * float64(i) / float64(j)))
	real.Mul(real, r)
	real.Add(real, x)
	imag := new(big.Rat).SetFloat64(math.Sin(2 * math.Pi * float64(i) / float64(j)))
	imag.Mul(imag, r)
	imag.Add(imag, y)
	return Complex{real, imag}
}
