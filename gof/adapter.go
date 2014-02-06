package main

import (
	"fmt"
)

type Adaptee struct{}

func (a *Adaptee) OldMethod() {
	fmt.Println("Adaptee.OldMethod()")
}

type Target interface {
	RequiredMethod()
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter() *Adapter {
	return &Adapter{new(Adaptee)}
}

func (a *Adapter) RequiredMethod() {
	fmt.Println("Adapter.RequiredMethod()")
	a.adaptee.OldMethod()
}

func main() {
	target := NewAdapter()
	target.RequiredMethod()
}
