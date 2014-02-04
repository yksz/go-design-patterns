package main

import (
	"fmt"
)

type Prototype interface {
	Clone()
}

type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Clone() *ConcretePrototype {
	return &ConcretePrototype{c.name}
}

func (c *ConcretePrototype) String() string {
	return "ConcretePrototype [name=" + c.name + "]"
}

func main() {
	prototype := &ConcretePrototype{"prototype1"}
	fmt.Println(prototype)
	fmt.Println(prototype.Clone())
}
