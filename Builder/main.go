package main

type House struct {
	Window string
	Floor  string
	Roof   string
}

type IHouseBuilder interface {
	reset()
	BuildRoof()
	BuildFloor()
	BuildWindow()
	GetHouse() House
}

type NormalHouseBuilder struct {
	house House
}

func (n *NormalHouseBuilder) reset() {
	n.house = House{}
}

func (n *NormalHouseBuilder) BuildRoof() {
	n.house.Roof = "Normal Roof"
}

func (n *NormalHouseBuilder) BuildFloor() {
	n.house.Floor = "Normal Floor"
}

func (n *NormalHouseBuilder) BuildWindow() {
	n.house.Window = "Normal Window"
}

func (n *NormalHouseBuilder) GetHouse() House {
	return n.house
}

type IHouseDirector interface {
	ChangeBuilder(b IHouseBuilder)
	Make(typeHouse string) House
}
type HouseDictor struct {
	houseBuilder IHouseBuilder
}

func (h *HouseDictor) ChangeBuilder(b IHouseBuilder) {
	h.houseBuilder = b
}

func (h *HouseDictor) Make(typeHouse string) House {
	h.houseBuilder.reset()
	if typeHouse == "normal" {
		h.houseBuilder.BuildRoof()
		h.houseBuilder.BuildFloor()
		h.houseBuilder.BuildWindow()
	} else {
		h.houseBuilder.BuildRoof()
		h.houseBuilder.BuildFloor()
	}
	return h.houseBuilder.GetHouse()
}

func main() {
	normalHouseBuilder := &NormalHouseBuilder{}
	houseDirector := &HouseDictor{houseBuilder: normalHouseBuilder}
	house := houseDirector.Make("normal")
	println(house.Floor)
	println(house.Roof)
	println(house.Window)
}
