package main

import "fmt"

type Car interface {
	getCar() string
}

type Sedan struct {
	Name string
}

func getNewSedan() *Sedan {
	return &Sedan{}
}

func (s *Sedan) getCar() string {
	return "SEDAN CAR"
}

type HatchBack struct {
	Name string
}

func getNewHatchBack() *HatchBack {
	return &HatchBack{}
}

func (h *HatchBack) getCar() string {
	return "HATCH BACK CAR"
}

func factory(CarType int) Car {
	if CarType == 1 {
		return getNewSedan()
	}
	return getNewHatchBack()
}

func main() {
	CarObj := factory(1) // here client interacted with factor and gave the type client Want , Client DOn't Care abput how Factory would make it or Get It
	fmt.Println(CarObj.getCar())
}
