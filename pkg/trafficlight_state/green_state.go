package trafficlight_state

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type greenState struct{
	DefaultLightState
}

func NewGreenState() *greenState{
	state :=  &greenState{}
	state.StateName = "GREEN"
	return state
}

func (state *greenState) Draw(light *TrafficLight){
	light.imd.Color = colornames.Darkgreen
	light.imd.Push(light.point)
	light.imd.Ellipse(pixel.V(100, 100), 0)
}

func (state *greenState) NextLight(light *TrafficLight){
	light.TransitionState(NewAmberState())
}