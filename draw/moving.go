package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func moveCircle() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	length := 0.0
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		length += 20 * dt

		imd := imdraw.New(nil)
		imd.Color = colornames.Darkgray
		line := pixel.V(100, 100)
		line.X += length
		line.Y += length
		imd.Push(line)
		imd.Circle(10, 1)

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}
