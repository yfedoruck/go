package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Limegreen
	imd.Push(pixel.V(500, 500))
	imd.Circle(100, 0)

	for !win.Closed() {
		win.Clear(colornames.Black)

		imd.Draw(win)

		mat := pixel.IM
		mat.Moved(win.Bounds().Center())

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
