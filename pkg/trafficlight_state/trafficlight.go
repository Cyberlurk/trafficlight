package trafficlight_state

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type LightState interface {
	Draw(light *TrafficLight)
	EnterState()
	NextLight(light *TrafficLight)
	CarPassingSpeed(*TrafficLight, int, string)
}

type DefaultLightState struct {
	StateName string
}

func (state DefaultLightState) CarPassingSpeed(road *TrafficLight, speed int, licensePlate string){
	if speed > road.SpeedLimit {
		fmt.Printf("Car with license %s was speeding\n", licensePlate)
	}
}

func (state *DefaultLightState) EnterState(){
	fmt.Println("changed state to:", state.StateName)
}

func NewSimpleTrafficLight(speedLimit int, position pixel.Vec, imDraw *imdraw.IMDraw) *TrafficLight {
	return &TrafficLight{SpeedLimit: speedLimit, point: position, State: new(redState), imd: imDraw}
}

type TrafficLight struct {
	State LightState
	SpeedLimit int

	// for drawing to screen
	point pixel.Vec
	imd *imdraw.IMDraw
}

func (tl *TrafficLight) TransitionState(newState LightState) {
	tl.State = newState
	tl.State.EnterState()
}


