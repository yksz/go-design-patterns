package main

import (
	"fmt"
)

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle()
}

func (c *Context) SetState(state State) {
	c.state = state
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

func main() {
	context := Context{new(ConcreteStateA)}
	context.Request()
	context.SetState(new(ConcreteStateB))
	context.Request()
}
