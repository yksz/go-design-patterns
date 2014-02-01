package main

import (
	"fmt"
)

type AbstractClass struct {
	template template
}

func (a *AbstractClass) TemplateMethod() {
	a.template.method1()
	a.template.method2()
}

type template interface {
	method1()
	method2()
}

type ConcreteClass struct {
}

func (c *ConcreteClass) method1() {
	fmt.Println("ConcreteClass.method1()")
}

func (c *ConcreteClass) method2() {
	fmt.Println("ConcreteClass.method2()")
}

func main() {
	abstractClass := AbstractClass{new(ConcreteClass)}
	abstractClass.TemplateMethod()
}
