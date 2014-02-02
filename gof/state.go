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

func (c *Context) SetState(s State) {
	c.state = s
}

type State interface {
	Handle()
}

type ConcreteStateA struct {
}

func (c *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct {
}

func (c *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}

func main() {
	context := Context{new(ConcreteStateA)}
	context.Request()
	context.SetState(new(ConcreteStateB))
	context.Request()
}
