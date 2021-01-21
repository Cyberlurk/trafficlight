package trafficlight_oop

type SimpleIntersection struct {
	xDirection 	*SimpleTrafficLight
	yDirection 	*SimpleTrafficLight
	xActive 	bool
}

func (intersection *SimpleIntersection) Step() {
	if intersection.xActive {
		intersection.xDirection.NextLight()
		if intersection.xDirection.Light == RED {
			intersection.xActive = false
			intersection.yDirection.NextLight()
		}
	} else {
		intersection.yDirection.NextLight()
		if intersection.yDirection.Light == RED {
			intersection.xActive = true
			intersection.xDirection.NextLight()
		}
	}
}

func (intersection *SimpleIntersection) Draw() {
	intersection.xDirection.Draw()
	intersection.yDirection.Draw()
}

func NewSimpleIntersection(xDirection *SimpleTrafficLight, yDirection *SimpleTrafficLight) *SimpleIntersection{
	return &SimpleIntersection{xDirection, yDirection, true}
}