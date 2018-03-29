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
	lineX0 := field.rectX0 + radius

	lengthY := 0.0
	lineY0 := field.rectY0 + radius

	line := pixel.V(field.rectX0+radius, lineY0)
	last := time.Now()

	safeVertical := field.vertical() - 2*radius
	safeHorizontal := field.horizontal() - 2*radius

	for !win.Closed() {
		win.Clear(colornames.Black)

		dt := time.Since(last).Seconds() * 100

		last = time.Now()

		if lengthX >= 2*safeHorizontal {
			lengthX = 0.0
		}
		if lengthX >= safeHorizontal {
			line.X -= dt
		} else {
			line.X = lineX0 + lengthX
		}

		if lengthY >= 2*safeVertical {
			lengthY = 0.0
		}
		if lengthY >= safeVertical {
			line.Y -= dt
		} else {
			line.Y = lineY0 + lengthY
		}

		lengthX += dt
		lengthY += dt

		circle(line, win, radius)

		imd := imdraw.New(nil)

		imd.Color = colornames.Blueviolet
		imd.Push(pixel.V(field.rectX0, field.rectY0), pixel.V(field.rectX1, field.rectY1))
		imd.Rectangle(1)

		imd.Draw(win)

		win.Update()
	}
}

func circle(line pixel.Vec, win *pixelgl.Window, radius float64) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Darkgray
	imd.Push(line)
	imd.Circle(radius, 0)
	imd.Draw(win)
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
