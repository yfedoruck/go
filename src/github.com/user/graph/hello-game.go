package main

import (
    _ "io"
    _ "strings"
    _ "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel"
    _ "golang.org/x/image/colornames"
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

    win.Clear(colornames.Darkgrey)

    for !win.Closed() {
        win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}
