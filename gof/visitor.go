package main

import (
	"fmt"
)

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementA.Accept()")
	visitor.VisitA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementB.Accept()")
	visitor.VisitB(e)
}

type Visitor interface {
	VisitA(*ConcreteElementA)
	VisitB(*ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor.VisitA()")
}

func (v *ConcreteVisitor) VisitB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor.VisitB()")
}

func main() {
	visitor := new(ConcreteVisitor)
	elementA := new(ConcreteElementA)
	elementB := new(ConcreteElementB)
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
