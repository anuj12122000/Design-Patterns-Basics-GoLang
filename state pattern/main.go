package main

import "fmt"

type tvState interface {
	state()
}

type on struct{}

func (on *on) state() {
	fmt.Println("TV is in On State")
}

type off struct{}

func (off *off) state() {
	fmt.Println("TV is in Off State")
}

type stateContext struct {
	currentTvState tvState // here we can observe that we can pass the interface as type too
	// this would be able to take both on and off type TvState
}

func getContext() *stateContext {
	return &stateContext{
		currentTvState: &off{},
	}
}

func (sc *stateContext) setState(state tvState) {
	sc.currentTvState = state
}

func (sc *stateContext) getCurrentState() {
	sc.currentTvState.state()
}

func main() {
	tvContext := getContext()
	tvContext.getCurrentState()
	tvContext.setState(&on{}) // here we changed the State of the tv
	tvContext.getCurrentState()
}
