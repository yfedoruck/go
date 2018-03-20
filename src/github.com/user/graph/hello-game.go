package main

import (
    _ "io"
    _ "strings"
    _ "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel"
    _ "golang.org/x/image/colornames"
    "golang.org/x/image/colornames"
    "os"
    "image"
    _ "image/png"
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
    cfg := pixelgl.WindowConfig{
        Title:  "Pixel Rocks!",
        Bounds: pixel.R(0, 0, 1024, 768),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    pic, err := loadPicture("vim-go.png")
    if err != nil {
        panic(err)
    }
    sprite := pixel.NewSprite(pic, pic.Bounds())
    win.Clear(colornames.Greenyellow)
    sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

    for !win.Closed() {
        win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}
