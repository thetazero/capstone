package main

import (
	"log"
	"math"
	"math/big"
	"strconv"

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
func (p Path) Draw(f EulerEquation, samples int, path string, wind bool) {
	plt, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(1)
	plt.X.Label.Text = "Real"
	plt.Y.Label.Text = "Imaginary"
	plt.BackgroundColor = colorful.Hsv(0, 0, .1)
	plt.X.Label.TextStyle.Color = colorful.Hsv(0, 0, 1)
	plt.X.Color = colorful.Hsv(0, 0, 1)
	plt.X.Tick.Color = colorful.Hsv(0, 0, 1)
	plt.X.Tick.Label.Color = colorful.Hsv(0, 0, 1)
	plt.Y.Label.TextStyle.Color = colorful.Hsv(0, 0, 1)
	plt.Y.Color = colorful.Hsv(0, 0, 1)
	plt.Y.Tick.Color = colorful.Hsv(0, 0, 1)
	plt.Y.Tick.Label.Color = colorful.Hsv(0, 0, 1)

	pts := make(plotter.XYs, samples)
	bar := pb.StartNew(samples)
	for i := 0; i < samples; i++ {
		val := f(p(i, samples))
		// val := p(i, j)

		real, _ := val[0].Float64()
		img, _ := val[1].Float64()
		pts[i] = plotter.XY{X: real, Y: img}
		bar.Increment()
	}
	bar.Finish()

	if wind {
		w := winding(pts)
		plt.Title.Text = "Winding: " + strconv.Itoa(w)
		plt.Title.TextStyle.Color = colorful.Hsv(0, 0, 1)
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	colors := moreland.Kindlmann() // Initialize a color map.
	colors.SetMax(float64(samples))
	colors.SetMin(0)
	scatter.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		hue := float64(i) / float64(samples) * 360
		c := colorful.Hsv(hue, 1, 1)
		return draw.GlyphStyle{Color: c, Radius: vg.Points(3), Shape: draw.CircleGlyph{}}
	}
	// fmt.Println(scatter.DataRange())
	plt.Add(scatter)

	//set start point to white
	scatter, err = plotter.NewScatter(plotter.XYs{pts[0]})
	scatter.GlyphStyle = draw.GlyphStyle{Color: colorful.Hsl(0, 0, 1), Radius: vg.Points(3), Shape: draw.CircleGlyph{}}
	plt.Add(scatter)

	// plotter.DefaultGlyphStyle.Radius = 2
	// plotter.DefaultGlyphStyle.Color = color.RGBA{0, 100, 100, 0xff}
	scatter, err = plotter.NewScatter(plotter.XYs{{X: 0, Y: 0}})
	scatter.GlyphStyle.Color = colorful.Hsv(0, 0, 1)
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

//make square with width=2, upper left corner=<x,yi>
func square(i, j int, w, x, y *big.Rat) Complex {
	if j%4 != 0 {
		panic("i%4 must equal zero")
	}
	c := Complex{x, y}
	zero := big.NewRat(0, 1)
	step := (j + 1) / 4
	if i < step {
		c.Sub(c, Complex{zero, new(big.Rat).Mul(w, big.NewRat(int64(i), int64(step)))})
	} else if i < 2*step {
		i -= step
		c.Sub(c, Complex{new(big.Rat).Mul(new(big.Rat).Mul(w, big.NewRat(-1, 1)), big.NewRat(int64(i), int64(step))), w})
	} else if i < 3*step {
		i -= step * 2
		c.Add(c, Complex{w, zero})
		c.Sub(c, Complex{zero, w})
		c.Add(c, Complex{zero, new(big.Rat).Mul(w, big.NewRat(int64(i), int64(step)))})
	} else {
		i -= step * 3
		c.Add(c, Complex{w, zero})
		c.Sub(c, Complex{new(big.Rat).Mul(w, big.NewRat(int64(i), int64(step))), zero})
	}
	return c
}

func winding(xys plotter.XYs) int {
	sum := 0.0
	last := math.Atan2(xys[0].Y, xys[0].X)
	for _, xy := range xys {
		cur := math.Atan2(xy.Y, xy.X)
		dθ := cur - last
		if dθ < -3 {
			dθ = 2*math.Pi + cur - last
		}
		sum += dθ
		last = cur
	}
	return int(math.Round(sum / math.Pi / 2))
}
