package main

type IPrototype interface {
	Clone() IPrototype
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Clone() IPrototype {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
}

type IPrototypeRegistry interface {
	GetClone(name string) IPrototype
	Register(name string, prototype IPrototype)
}

type PrototypeRegistry struct {
	prototypes map[string]IPrototype
}

func (p *PrototypeRegistry) GetClone(name string) IPrototype {
	if p.prototypes[name] == nil {
		panic("Prototype not found")
	}
	return p.prototypes[name].Clone()
}

func (p *PrototypeRegistry) Register(name string, prototype IPrototype) {
	p.prototypes[name] = prototype
}

func main() {
	registry := &PrototypeRegistry{
		prototypes: make(map[string]IPrototype),
	}

	person := &Person{
		Name: "John",
		Age:  30,
	}

	registry.Register("person", person)

	personClone := registry.GetClone("person").(*Person)
	personClone.Age = 10
	println("Clone Age: ", personClone.Age)
	println("Original Age: ", person.Age)
}
