package main

import (
	"image"
	"math/rand"

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
	//win.SetSmooth(true)

	spriteSheet, err := loadPicture("trees.png")
	if err != nil {
		panic(err)
	}

	var treeFrames []pixel.Rect

	var (
		trees    []*pixel.Sprite
		matrices []pixel.Matrix
	)

	for x := spriteSheet.Bounds().Min.X; x < spriteSheet.Bounds().Max.X; x += 32 {
		for y := spriteSheet.Bounds().Min.Y; y < spriteSheet.Bounds().Max.Y; y += 32 {
			treeFrames = append(treeFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	//tree := pixel.NewSprite(spriteSheet, treeFrames[5])
	win.Clear(colornames.Black)

	//last := time.Now()
	//angle := 0.0
	for !win.Closed() {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := pixel.NewSprite(spriteSheet, treeFrames[rand.Intn(len(treeFrames))])
			trees = append(trees, tree)
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(win.MousePosition()))
		}
		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}
		//tree.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))

		//dt := time.Since(last).Seconds()
		//last = time.Now()
		//angle += 10 * dt

		//win.Clear(colornames.Black)

		//mat := pixel.IM
		//mat = mat.Rotated(pixel.ZV, angle)
		//mat = mat.Moved(win.Bounds().Center())

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
