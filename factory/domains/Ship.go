package domains

type Ship struct {
	name  string
	power int
}

func (s *Ship) GetName() string {
	return s.name
}

func (s *Ship) GetPower() int {
	return s.power
}

func (Ship) Start() {
	println("Ship is started")
}
