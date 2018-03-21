package main

import (
	"image"

	_ "image/png"
	_ "io"
	"math"
	"os"
	_ "strings"

	"github.com/faiface/pixel"
	_ "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	var _ = math.Pi

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	pic, err := loadPicture("picture.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	win.Clear(colornames.Black)
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	//mat = mat.Rotated(win.Bounds().Center(), math.Pi/4)
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.5, 2))
	sprite.Draw(win, mat)
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
