package main

import "project/domains"

type LogisticsFactory struct {
}

func (t *LogisticsFactory) CreateLogistics(logisticsType string) ILogistics {
	if logisticsType == "Truck" {
		return &domains.Truck{}
	}
	return nil
}

type ILogistics interface {
	GetName() string
	GetPower() int
	Start()
}

func (t *LogisticsFactory) CreateTransport(transportType string) ILogistics {
	if transportType == "Truck" {
		return &domains.Truck{}
	}
	if transportType == "Ship" {
		return &domains.Ship{}
	}
	if transportType == "AirPlane" {
		return &domains.AirPlane{}
	}
	return nil
}

func main() {
	factory := &LogisticsFactory{}
	t := factory.CreateTransport("AirPlane")
	t.Start()
}
