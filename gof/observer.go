package main

import (
	"container/list"
	"fmt"
)

type Observer interface {
	Notify()
}

type Subject struct {
	observerCollection *list.List
}

func NewSubject() *Subject {
	return &Subject{list.New()}
}

func (s *Subject) RegisterObserver(o Observer) {
	s.observerCollection.PushBack(o)
}

func (s *Subject) UnregisterObserver(o Observer) {
	for e := s.observerCollection.Front(); e != nil; e = e.Next() {
		if o == e.Value.(Observer) {
			s.observerCollection.Remove(e)
		}
	}
}

func (s *Subject) NotifyObservers() {
	for e := s.observerCollection.Front(); e != nil; e = e.Next() {
		observer := e.Value.(Observer)
		observer.Notify()
	}
}

type ConcreteObserver struct {
}

func (c *ConcreteObserver) Notify() {
	fmt.Println("ConcreteObserver.Notify()")
}

func main() {
	concreteObserver := new(ConcreteObserver)

	subject := NewSubject()
	subject.RegisterObserver(concreteObserver)
	subject.NotifyObservers()
	subject.UnregisterObserver(concreteObserver)
	subject.NotifyObservers()
}
