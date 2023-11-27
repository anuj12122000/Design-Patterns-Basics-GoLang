package main

import "fmt"

type IGun interface {
	GetPower() string
	GetName() string
}

type Gun struct {
	power string
	name  string
}

func (g Gun) GetName() string {
	return g.name
}

func (g Gun) GetPower() string {
	return g.power
}

type Ak47 struct {
	Gun // here we embed / inject the Gun struct in the ak47 struct
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "ak47",
			power: "2",
		},
	}
}

type Sniper struct {
	Gun
}

func newSniper() IGun {
	return &Sniper{
		Gun: Gun{
			name:  "Sniper",
			power: "4",
		},
	}
}

func getFactory(id int) IGun {
	if id == 1 {
		return newAk47()
	}
	return newSniper()
}

func main() {
	firstType := getFactory(1)
	fmt.Println(firstType)
}
