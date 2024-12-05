package domains

type Truck struct {
	name  string
	power int
}

func (t *Truck) GetName() string {
	return t.name
}

func (t *Truck) GetPower() int {
	return t.power
}

func (t *Truck) Start() {
	println("Truck is started")
}
