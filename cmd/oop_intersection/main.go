package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"trafficLight/pkg/trafficlight_oop"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Intersection",
		Bounds: pixel.R(0, 0, 1024, 1024),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Precision = 100

	NorthernTrafficLight := trafficlight_oop.NewSimpleTrafficLight(pixel.V(341, 200), imd)
	EasternTrafficLight := trafficlight_oop.NewSimpleTrafficLight(pixel.V(682, 200), imd)

	intersection := trafficlight_oop.NewSimpleIntersection(NorthernTrafficLight, EasternTrafficLight)

	for !win.Closed() {
		// update objects
		intersection.Step()

		// erase canvas
		win.Clear(colornames.Black)
		imd.Clear()

		// draw
		intersection.Draw()
		imd.Draw(win)

		// fetch new events (key presses, mouse moves and clicks, etc.)
		// and redraw the window
		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
