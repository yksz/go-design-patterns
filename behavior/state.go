package main

import "fmt"

type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func (c *Context) SetState(state State) {
	c.state = state
}

type State interface {
	Handle(*Context)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(ctx *Context) {
	fmt.Println("ConcreteStateA.Handle()")
	ctx.SetState(new(ConcreteStateB))
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(ctx *Context) {
	fmt.Println("ConcreteStateB.Handle()")
	ctx.SetState(new(ConcreteStateA))
}

func main() {
	context := Context{new(ConcreteStateA)}
	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
