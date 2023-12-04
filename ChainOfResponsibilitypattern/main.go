package main

import "fmt"

type step interface {
	Execute()
	setNext(step)
}

type OpenFlap struct {
	next step
}

func (o *OpenFlap) setNext(nextStep step) {
	o.next = nextStep
}

func (o *OpenFlap) Execute() {
	fmt.Println("OPENING THE LAPTOP FLAP")
	o.next.Execute()
}

type PressButton struct {
	next step
}

func (pb *PressButton) setNext(nextStep step) {
	pb.next = nextStep
}

func (pb *PressButton) Execute() {
	fmt.Println("PRESSING THE BUTTON")
	pb.next.Execute()
}

type EnterPassword struct {
	next step
}

func (eP *EnterPassword) Execute() {
	fmt.Println("Entering Password and Loggin In")
	fmt.Println("Laptop Opened")
}

func (ep *EnterPassword) setNext(nextStep step) {
	ep.next = nextStep
}

func main() {

	step3 := &EnterPassword{}

	step2 := &OpenFlap{next: step3}

	step1 := &PressButton{next: step2}

	step1.Execute()

	return
}
