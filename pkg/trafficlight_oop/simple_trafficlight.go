package trafficlight_oop

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Light int

//goland:noinspection GoSnakeCaseUsage
const (
	RED Light = iota
	RED_AND_AMBER
	AMBER
	GREEN
	BUSSES_ALLOWED_ONLY
)

type SimpleTrafficLight struct {
	Light Light

	// for drawing to screen
	point pixel.Vec
	imd *imdraw.IMDraw
}

func NewSimpleTrafficLight(position pixel.Vec, imdP *imdraw.IMDraw) *SimpleTrafficLight {
	return &SimpleTrafficLight{point: position, Light: RED, imd: imdP}
}

func (tl *SimpleTrafficLight) NextLight(){
	switch tl.Light {
	case RED:
		tl.Light = RED_AND_AMBER
	case RED_AND_AMBER:
		tl.Light = BUSSES_ALLOWED_ONLY
	case AMBER:
		tl.Light = RED
	case BUSSES_ALLOWED_ONLY:
		tl.Light = GREEN
	case GREEN:
		tl.Light = AMBER
	}
}

func (tl *SimpleTrafficLight) CarPassingSpeed(speed int, licensePlate string){
	speedLimit := 50

	switch tl.Light {
	case BUSSES_ALLOWED_ONLY:
		speedLimit = 0

		if isNormalCar(licensePlate) && speed > 0 {
			redLightRunner(licensePlate)
		} else {
			if speed > speedLimit {
				redLightRunner(licensePlate)
			}
		}
	case RED:
		speedLimit = 0
		if speed > speedLimit {
			redLightRunner(licensePlate)
		}
	case RED_AND_AMBER:
		if speed > speedLimit {
			reportDriver(licensePlate)
		}
	case AMBER:
		if speed > speedLimit {
			reportDriver(licensePlate)
		}
	case GREEN:
		if speed > speedLimit {
			reportDriver(licensePlate)
		}
	}
}

func isNormalCar(licensePlate string) bool {
	// logic to distinguish between cars and buses based on their license plate
	return false
}

func redLightRunner(licensePlate string){}

func reportDriver(licensePlate string) {}

func (tl *SimpleTrafficLight) ShowRed() {
	tl.imd.Color = colornames.Darkred
	tl.imd.Push(pixel.V(tl.point.X, tl.point.Y+600))
	tl.imd.Ellipse(pixel.V(100, 100), 0)
}

func (tl *SimpleTrafficLight) ShowRedAndAmber() {
	tl.imd.Color = colornames.Darkred
	tl.imd.Push(pixel.V(tl.point.X, tl.point.Y+600))
	tl.imd.Ellipse(pixel.V(100, 100), 0)

	tl.imd.Color = colornames.Darkgoldenrod
	tl.imd.Push(pixel.V(tl.point.X, tl.point.Y+300))
	tl.imd.Ellipse(pixel.V(100, 100), 0)
}

func (tl *SimpleTrafficLight) ShowAmber() {
	tl.imd.Color = colornames.Darkgoldenrod
	tl.imd.Push(pixel.V(tl.point.X, tl.point.Y+300))
	tl.imd.Ellipse(pixel.V(100, 100), 0)
}

func (tl *SimpleTrafficLight) ShowGreen() {
	tl.imd.Color = colornames.Darkgreen
	tl.imd.Push(tl.point)
	tl.imd.Ellipse(pixel.V(100, 100), 0)
}

func (tl *SimpleTrafficLight) ShowPurple() {
	tl.imd.Color = colornames.Purple
	x := tl.point.X - 200.0
	point := pixel.Vec{X: x, Y: tl.point.Y}
	tl.imd.Push(point)
	tl.imd.Ellipse(pixel.V(100, 100), 0)
}

func (tl *SimpleTrafficLight) Draw() {
	switch tl.Light {
	case RED:
		tl.ShowRed()
	case RED_AND_AMBER:
		tl.ShowRedAndAmber()
	case AMBER:
		tl.ShowAmber()
	case GREEN:
		tl.ShowGreen()
	case BUSSES_ALLOWED_ONLY:
		tl.ShowPurple()
	}
}
