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

	length := 0.0
	radius := 20.0
	lineY0 := field.rectY0 + radius
	line := pixel.V(field.rectX0+radius, lineY0)
	last := time.Now()

	for !win.Closed() {
		win.Clear(colornames.Black)

		dt := time.Since(last).Seconds() * 300

		//fmt.Println(dt)
		last = time.Now()
		//line := pixel.V(field.rectX0+radius, field.rectY0+radius)

		//setLine(&line, length, field, radius)
		//if line.X+length >= field.rectX1 {
		//	line.X = field.rectX0 + radius
		//} else {
		//	line.X += length
		//}

		if length >= 2*field.vertical() {
			length = 0.0
		}

		if (length >= field.vertical()) && (length < 2*field.vertical()) {
			line.Y -= dt
		} else {
			line.Y = lineY0 + length
		}
		length += dt

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
