package domains

type AirPlane struct {
	name  string
	power int
}

func (a *AirPlane) GetName() string {
	return a.name
}

func (a *AirPlane) GetPower() int {
	return a.power
}

func (AirPlane) Start() {
	println("AirPlane is started")
}
