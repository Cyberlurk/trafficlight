package trafficlight_state

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type amberState struct {
	DefaultLightState
}

func NewAmberState() *amberState{
	state :=  &amberState{}
	state.StateName = "AMBER"
	return state
}

func (state *amberState) Draw(light *TrafficLight){
	light.imd.Color = colornames.Darkgoldenrod
	light.imd.Push(pixel.V(light.point.X, light.point.Y+300))
	light.imd.Ellipse(pixel.V(100, 100), 0)
}

func (state *amberState) NextLight(light *TrafficLight){
	light.TransitionState(NewRedState())
}