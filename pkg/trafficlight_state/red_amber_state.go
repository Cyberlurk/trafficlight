package trafficlight_state

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type redAmberState struct {
	DefaultLightState
}

func NewRedAmberState() *redAmberState{
	state :=  &redAmberState{}
	state.StateName = "RED & AMBER"
	return state
}

func (state *redAmberState) Draw(light *TrafficLight){
	light.imd.Color = colornames.Darkred
	light.imd.Push(pixel.V(light.point.X, light.point.Y+600))
	light.imd.Ellipse(pixel.V(100, 100), 0)

	light.imd.Color = colornames.Darkgoldenrod
	light.imd.Push(pixel.V(light.point.X, light.point.Y+300))
	light.imd.Ellipse(pixel.V(100, 100), 0)
}

func (state *redAmberState)	NextLight(light *TrafficLight){
	light.TransitionState(NewGreenState())
}