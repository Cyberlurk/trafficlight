package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
	"trafficLight/pkg/trafficlight_oop"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Traffic light 1",
		Bounds: pixel.R(0, 0, 1024, 1024),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Precision = 100

	NorthernTrafficLight := trafficlight_oop.NewSimpleTrafficLight(pixel.V(512, 200), imd)
	NorthernTrafficLight.Light = trafficlight_oop.AMBER

	for !win.Closed() {
		NorthernTrafficLight.CarPassingSpeed(30, "abc")
		NorthernTrafficLight.NextLight()

		// erase canvas
		win.Clear(colornames.Black)
		imd.Clear()

		// draw
		NorthernTrafficLight.Draw()
		imd.Draw(win)

		// fetch new events (key presses, mouse moves and clicks, etc.)
		// and redraw the window
		win.Update()

		time.Sleep(1 * time.Second)
	}
}

func main() {
	pixelgl.Run(run)
}
