package main

type IChair interface {
	SitOn()
}

type ISofa interface {
	LieOn()
}

type ArtDecoChair struct {
}

func (a *ArtDecoChair) SitOn() {
	println("ArtDecoChair: sit on")
}

type ArtDecoSofa struct{}

func (a *ArtDecoSofa) LieOn() {
	println("ArtDecoSofa: lie on")
}

type ArtDecoTable struct{}

func (a *ArtDecoTable) PutOn() {
	println("ArtDecoTable: put on")
}

type ModernChair struct{}

func (m *ModernChair) SitOn() {
	println("ModernChair: sit on")
}

type ModernSofa struct{}

func (m *ModernSofa) LieOn() {
	println("ModernSofa: lie on")
}

type ModernTable struct{}

func (m *ModernTable) PutOn() {
	println("ModernTable: put on")
}

type IFurnitureFactory interface {
	createChair() IChair
	createSofa() ISofa
}

type ArtDecoFurnitureFactory struct{}

func (a *ArtDecoFurnitureFactory) createChair() IChair {
	return &ArtDecoChair{}
}

func (a *ArtDecoFurnitureFactory) createSofa() ISofa {
	return &ArtDecoSofa{}
}

type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) createChair() IChair {
	return &ModernChair{}
}

func (m *ModernFurnitureFactory) createSofa() ISofa {
	return &ModernSofa{}
}

type IWorker interface {
	createNewChair(f *IFurnitureFactory) IChair
	createNewSofa(f *IFurnitureFactory) ISofa
}

type Worker struct{}

func (w *Worker) createNewChair(f IFurnitureFactory) IChair {
	return f.createChair()
}

func (w *Worker) createNewSofa(f IFurnitureFactory) ISofa {
	return f.createSofa()
}

func main() {
	worker := &Worker{}
	artDecoFactory := &ArtDecoFurnitureFactory{}
	worker.createNewChair(artDecoFactory).SitOn()
}
