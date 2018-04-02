package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Rect struct {
	rectX0,
	rectY0,
	rectX1,
	rectY1 float64
}

func (r *Rect) horizontal() float64 {
	return r.rectX1 - r.rectX0
}
func (r *Rect) vertical() float64 {
	return r.rectY1 - r.rectY0
}

func (r *Rect) safeVertical(radius float64) float64 {
	return r.vertical() - 2*radius
}

func (r *Rect) safeHorizontal(radius float64) float64 {
	return r.horizontal() - 2*radius
}

func run() {

	winX := 1.3 * 1024.0
	winY := 1.3 * 768.0

	field := Rect{
		rectX0: winX / 4,
		rectY0: winY / 4,
		rectX1: 3 * winX / 4,
		rectY1: 3 * winY / 4,
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, winX, winY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	radius := 120.0

	lengthX := 0.0

	//startDx := 50.0
	//startDy := 50.0

	//lineX0 := field.rectX0 + radius + startDx
	lineX0 := field.rectX0 + radius

	lengthY := 0.0
	//lineY0 := field.rectY0 + radius + startDy
	lineY0 := field.rectY0 + radius

	line := pixel.V(lineX0, lineY0)
	last := time.Now()

	safeVertical := field.safeVertical(radius)
	safeHorizontal := field.safeHorizontal(radius)

	for !win.Closed() {
		win.Clear(colornames.Black)
		dt := time.Since(last).Seconds() * 300
		last = time.Now()

		lengthX = run2(&line.X, lengthX, safeHorizontal, dt, lineX0)
		lengthY = run2(&line.Y, lengthY, safeVertical, dt, lineY0)

		imd := imdraw.New(nil)
		circle(line, imd, radius)

		imd.Color = colornames.Blueviolet
		imd.Push(pixel.V(field.rectX0, field.rectY0), pixel.V(field.rectX1, field.rectY1))
		imd.Rectangle(1)

		imd.Draw(win)

		win.Update()
	}
}

func run2(vectorAxis *float64, lengthD, safeLength, delta, lineD0 float64) float64 {
	if lengthD >= 2*safeLength {
		lengthD = 0.0
	}
	if lengthD >= safeLength {
		*vectorAxis -= delta
	} else {
		*vectorAxis = lineD0 + lengthD
	}

	lengthD += delta

	return lengthD
}

type Moving struct {
	vector                             pixel.Vec
	lengthD, safeLength, delta, lineD0 float64
	axis                               string
}

func (a Moving) run() (float64, pixel.Vec) {
	if a.lengthD >= 2*a.safeLength {
		a.lengthD = 0.0
	}
	if a.lengthD >= a.safeLength {
		if a.axis == "X" {
			a.vector.X -= a.delta
		} else {
			a.vector.Y -= a.delta
		}
	} else {
		a.vector.X = a.lineD0 + a.lengthD
	}

	a.lengthD += a.delta

	return a.lengthD, a.vector
}

func circle(line pixel.Vec, imd *imdraw.IMDraw, radius float64) {
	imd.Color = colornames.Darkgray
	imd.Push(line)
	imd.Circle(radius, 0)
}

func setLine(line *pixel.Vec, length float64, field Rect, radius float64) {
	if line.X+length >= field.rectX1 {
		line.X = field.rectX0 + radius
	} else {
		line.X += length
	}

	if line.Y+length >= field.rectY1 {
		line.Y = field.rectY0 + radius
		length = 0.0
	} else {
		line.Y += length
	}
}

func main() {
	pixelgl.Run(run)
}
