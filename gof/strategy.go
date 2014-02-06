package main

import (
	"fmt"
)

type Context struct {
	strategy func()
}

func (c *Context) Execute() {
	c.strategy()
}

func (c *Context) SetStrategy(strategy func()) {
	c.strategy = strategy
}

func main() {
	concreteStrategyA := func() {
		fmt.Println("concreteStrategyA()")
	}
	concreteStrategyB := func() {
		fmt.Println("concreteStrategyB()")
	}
	context := Context{concreteStrategyA}
	context.Execute()
	context.SetStrategy(concreteStrategyB)
	context.Execute()
}
