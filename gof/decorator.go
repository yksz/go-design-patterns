package main

import (
	"fmt"
)

type Component interface {
	Operation() string
}

type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent.Operation()"
}

type ConcreteDecoratorA struct {
	component Component
}

func (d *ConcreteDecoratorA) Operation() string {
	if d.component == nil {
		return "ConcreteDecoratorA.Operation()"
	} else {
		return d.component.Operation() + " and ConcreteDecoratorA.Operation()"
	}
}

type ConcreteDecoratorB struct {
	component Component
}

func (d *ConcreteDecoratorB) Operation() string {
	if d.component == nil {
		return "ConcreteDecoratorB.Operation()"
	} else {
		return d.component.Operation() + " and ConcreteDecoratorB.Operation()"
	}
}

func main() {
	component := &ConcreteDecoratorA{&ConcreteDecoratorB{new(ConcreteComponent)}}
	fmt.Println(component.Operation())
}
