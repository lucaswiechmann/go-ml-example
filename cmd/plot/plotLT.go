package plot

import (
	"encoding/csv"
	"flag"
	"fmt"
	"image/color"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}

func PlotLR() {
	flag.Parse()
	if len(flag.Args()) < 3 {
		fmt.Println("usage: plotLR filename a b")
		return
	}

	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)

	a, err := strconv.ParseFloat(flag.Args()[1], 64)
	if err != nil {
		fmt.Println(a, "not a valid float")
		return
	}

	b, err := strconv.ParseFloat(flag.Args()[2], 64)
	if err != nil {
		fmt.Println(b, "not a valid float")
		return
	}

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	size := len(records)

	data := xy{
		x: make([]float64, size),
		y: make([]float64, size),
	}

	for i, v := range records {
		if len(v) != 2 {
			fmt.Println("Expected two elements per line")
			continue
		}

		s, err := strconv.ParseFloat(v[0], 64)
		if err == nil {
			data.y[i] = s
		}

		s, err = strconv.ParseFloat(v[1], 64)
		if err == nil {
			data.x[i] = s
		}
	}

	line := plotter.NewFunction(func(x float64) float64 { return a*x + b })
	line.Color = color.RGBA{B: 255, A: 255}

	p := plot.New()

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(scatter, line)

	w, err := p.WriterTo(300, 300, "svg")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = w.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
}
