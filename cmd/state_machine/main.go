package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
	"trafficLight/pkg/trafficlight_state"
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

	NorthernTrafficLight := trafficlight_state.NewSimpleTrafficLight(50, pixel.V(512, 200), imd)
	NorthernTrafficLight.State = trafficlight_state.NewAmberState()

	interval := time.NewTicker(1000 * time.Millisecond)
	for !win.Closed() {
		select {
		case <-interval.C:
			// erase canvas
			win.Clear(colornames.Black)
			imd.Clear()

			// have object draw to imd
			NorthernTrafficLight.State.NextLight(NorthernTrafficLight)
			NorthernTrafficLight.State.Draw(NorthernTrafficLight)
			NorthernTrafficLight.State.CarPassingSpeed(NorthernTrafficLight, 25, "abc 123")
		default:
		}

		// draw
		imd.Draw(win)
		// fetch new events (key presses, mouse moves and clicks, etc.)
		// and redraw the window
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
