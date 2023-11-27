package main

import "fmt"

// the main thing in builder pattern is that we do have a director jo ke
// we gave our builders to that director
// director can operate and Change the builder Anytime as they want

type House struct {
	Floor   int
	Doors   int
	Windows int
}

type Ibuilder interface {
	setWindows()
	setDoors()
	setFloors()
	getHouse() House
}

func GetBuilder(builderType int) Ibuilder {
	if builderType == 1 {
		return newNormalBuilder()
	}

	return newFashionBuilder()
}

type NormalBuilder struct {
	Floor   int
	Doors   int
	Windows int
}

func (nb *NormalBuilder) setFloors() {
	nb.Floor = 4
}

func (nb *NormalBuilder) setWindows() {
	nb.Windows = 8
}

func (nb *NormalBuilder) setDoors() {
	nb.Doors = 12
}

func (nb *NormalBuilder) getHouse() House {
	return House{
		Floor:   nb.Floor,
		Windows: nb.Windows,
		Doors:   nb.Doors,
	}
}

func newNormalBuilder() Ibuilder {
	return &NormalBuilder{}
}

type FashionBuilder struct {
	Floor   int
	Doors   int
	Windows int
}

func (f *FashionBuilder) setFloors() {
	f.Floor = 2
}

func (f *FashionBuilder) setWindows() {
	f.Windows = 5
}

func (f *FashionBuilder) setDoors() {
	f.Doors = 7
}

func (f *FashionBuilder) getHouse() House {
	return House{
		Floor:   f.Floor,
		Windows: f.Windows,
		Doors:   f.Doors,
	}
}

func newFashionBuilder() Ibuilder {
	return &FashionBuilder{}
}

// below are all director Related things
type Director struct { // director would be controlling the Builder , not the Factory Itself
	builder Ibuilder
}

func newDirector(b Ibuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Ibuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoors()
	d.builder.setFloors()
	d.builder.setWindows()
	return d.builder.getHouse()
}

func main() {
	NormalBuilder := newNormalBuilder()
	FashionBuilder := newFashionBuilder()

	Director := newDirector(NormalBuilder)
	normalHouse := Director.buildHouse()
	fmt.Println(normalHouse)

	Director.setBuilder(FashionBuilder)
	FasionHouse := Director.buildHouse()
	fmt.Println(FasionHouse)

}
