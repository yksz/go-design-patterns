package main

import (
	"fmt"
)

type Handler interface {
	Request(flag bool)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) Request(flag bool) {
	fmt.Println("ConcreteHandlerA.Request()")
	if flag {
		h.next.Request(flag)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) Request(flag bool) {
	fmt.Println("ConcreteHandlerB.Request()")
}

func main() {
	handlerA := &ConcreteHandlerA{new(ConcreteHandlerB)}
	handlerA.Request(true)
}
