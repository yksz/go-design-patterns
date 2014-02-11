package main

import (
	"fmt"
)

type Subject interface {
	DoAction()
}

type RealSubject struct{}

func (s *RealSubject) DoAction() {
	fmt.Println("RealSubject.DoAction()")
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) DoAction() {
	if p.realSubject == nil {
		p.realSubject = new(RealSubject)
	}
	fmt.Println("Proxy.DoAction()")
	p.realSubject.DoAction()
}

func main() {
	subject := new(Proxy)
	subject.DoAction()
}
