package main

import (
	"fmt"
)

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct {
}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementA.Accept()")
	visitor.VisitA(c)
}

type ConcreteElementB struct {
}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementB.Accept()")
	visitor.VisitB(c)
}

type Visitor interface {
	VisitA(*ConcreteElementA)
	VisitB(*ConcreteElementB)
}

type ConcreteVisitor struct {
}

func (c *ConcreteVisitor) VisitA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor.VisitA()")
}

func (c *ConcreteVisitor) VisitB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor.VisitB()")
}

func main() {
	visitor := new(ConcreteVisitor)
	elementA := new(ConcreteElementA)
	elementB := new(ConcreteElementB)
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
