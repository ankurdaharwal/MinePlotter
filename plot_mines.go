package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	xys, err := generateData()
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func generateData() (plotter.XYs, error) {

	var xys plotter.XYs
	var noMinePoints int
	for i := -200; i < 200; i++ {
		for j := -200; j < 200; j++ {

			fmt.Printf("\nI: %v", i)
			fmt.Printf("| J: %v", j)
			sum := sumDigits(i, j)
			if sum > 21 {
				xys = append(xys, struct{ X, Y float64 }{float64(i), float64(j)})
			}
			if sum < 22 {
				noMinePoints++
			}
		}
	}
	fmt.Printf("\n\nTotal Points %v\n", noMinePoints)
	return xys, nil
}

func sumDigits(x int, y int) (sum int) {

	var sumX, sumY int

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Printf("\nX: %v", x)
	fmt.Printf("| Y: %v", y)
	for x > 0 {
		sumX += x % 10
		x = x / 10
	}
	for y > 0 {
		sumY += y % 10
		y = y / 10
	}
	fmt.Printf("| X+Y: %v", sumX+sumY)
	return sumX + sumY
}

func plotData(path string, xys plotter.XYs) error {

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, B: 55, A: 255}
	p.Add(s)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
