package trafficlight_state

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"log"
)

type redState struct{
	DefaultLightState
}

func NewRedState() *redState {
	state :=  &redState{}
	state.StateName = "RED"
	return state
}

func (state *redState) Draw(light *TrafficLight){
	light.imd.Color = colornames.Darkred
	light.imd.Push(pixel.V(light.point.X, light.point.Y+600))
	light.imd.Ellipse(pixel.V(100, 100), 0)
}

func (state redState) CarPassingSpeed(light *TrafficLight, speed int, licensePlate string){
	if speed > 0 {
		log.Printf("Car with license \"%s\" ran a red light!\n", licensePlate)
	}
}

func (state *redState) NextLight(light *TrafficLight){
	light.TransitionState(NewRedAmberState())
}