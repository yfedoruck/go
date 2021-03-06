package main

import (
	"fmt"
	"math"
	"math/rand"
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

	last := time.Now()

	const count = 3

	var circles [count]Circle
	for i := 0; i < count; i++ {
		circles[i] = Circle{}
		circles[i].build(&field)
	}

	fmt.Println(circles)

	for !win.Closed() {

		win.Clear(colornames.Black)

		imd := imdraw.New(nil)

		for i := 0; i < count; i++ {
			dt := time.Since(last).Seconds() * circles[i].velocity
			next := i + 1
			if next == count {
				next = 0
			}
			circles[i].move1(dt, &circles[next])
			circles[i].draw(imd)
		}

		last = time.Now()

		imd.Color = colornames.Blueviolet
		imd.Push(pixel.V(field.rectX0, field.rectY0), pixel.V(field.rectX1, field.rectY1))
		imd.Rectangle(1)

		imd.Draw(win)

		win.Update()
	}
}

type Circle struct {
	line     pixel.Vec
	direct   Direction
	radius   float64
	rec      *Rect
	velocity float64
}

func (c *Circle) build(rec *Rect) {
	c.radius = 60.0

	lineX0 := rec.rectX0 + c.radius + float64(rand.Intn(int(rec.rectX1-c.radius)))

	lineY0 := rec.rectY0 + c.radius + float64(rand.Intn(int(rec.rectY1-c.radius)))

	c.line = pixel.V(lineX0, lineY0)
	c.direct = Direction{
		X: randomBool(),
		Y: randomBool(),
	}
	c.rec = rec
	c.velocity = float64(rand.Intn(500))
}

func (c *Circle) draw(imd *imdraw.IMDraw) {
	imd.Color = colornames.Darkgray
	imd.Push(c.line)
	imd.Circle(c.radius, 0)
}

func (c *Circle) move1(delta float64, cr *Circle) {

	vector := &c.line
	dir := &c.direct
	radius := c.radius
	rec := c.rec

	if dir.X == true {
		vector.X += delta
	} else {
		vector.X -= delta
	}

	if dir.Y == true {
		vector.Y += delta
	} else {
		vector.Y -= delta
	}

	// X - axis
	if vector.X <= rec.rectX0+radius {
		vector.X = rec.rectX0 + radius
		dir.X = true
	}

	if vector.X >= rec.rectX1-radius {
		vector.X = rec.rectX1 - radius
		dir.X = false
	}

	// Y - axis
	if vector.Y <= rec.rectY0+radius {
		vector.Y = rec.rectY0 + radius
		dir.Y = true
	}

	if vector.Y >= rec.rectY1-radius {
		vector.Y = rec.rectY1 - radius
		dir.Y = false
	}

	length := math.Sqrt(math.Pow(vector.X-cr.line.X, 2) + math.Pow(vector.Y-cr.line.Y, 2))

	if length <= 2*radius {

		//dx := math.Abs(vector.X - cr.line.X)

		if vector.X > cr.line.X {
			if dir.X == cr.direct.X {
				cr.direct.X = !cr.direct.X
			} else if dir.X == false {
				dir.X = !dir.X
			}
		} else {
			if dir.X == cr.direct.X {
				dir.X = !dir.X
			} else if cr.direct.X == false {
				cr.direct.X = !cr.direct.X
			}
		}

		if vector.Y > cr.line.Y {
			if dir.Y == cr.direct.Y {
				cr.direct.Y = !cr.direct.Y
			} else if dir.Y == false {
				dir.Y = !dir.Y
			}
		} else {
			if dir.Y == cr.direct.Y {
				dir.Y = !dir.Y
			} else if cr.direct.Y == false {
				cr.direct.Y = !cr.direct.Y
			}
		}
	}
}

func (c *Circle) move(delta float64) {

	vector := &c.line
	dir := &c.direct
	radius := c.radius
	rec := c.rec

	// X - axis
	if dir.X == true {
		vector.X += delta
	} else {
		vector.X -= delta
	}

	if vector.X <= rec.rectX0+radius {
		vector.X = rec.rectX0 + radius
		dir.X = true
	}

	if vector.X >= rec.rectX1-radius {
		vector.X = rec.rectX1 - radius
		dir.X = false
	}

	// Y - axis
	if dir.Y == true {
		vector.Y += delta
	} else {
		vector.Y -= delta
	}

	if vector.Y <= rec.rectY0+radius {
		vector.Y = rec.rectY0 + radius
		dir.Y = true
	}

	if vector.Y >= rec.rectY1-radius {
		vector.Y = rec.rectY1 - radius
		dir.Y = false
	}
}

func random(min, max int) float64 {
	rand.Seed(time.Now().Unix())
	return float64(rand.Intn(max-min) + min)
}

func randomBool() bool {
	return rand.Intn(2) == 0
}

type Direction struct {
	X, Y bool
}

func circle(line pixel.Vec, imd *imdraw.IMDraw, radius float64) {
	imd.Color = colornames.Darkgray
	imd.Push(line)
	imd.Circle(radius, 0)
}

func main() {
	pixelgl.Run(run)
}
