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

func (s *Subject) RegisterObserver(observer Observer) {
	s.observerCollection.PushBack(observer)
}

func (s *Subject) UnregisterObserver(observer Observer) {
	for e := s.observerCollection.Front(); e != nil; e = e.Next() {
		if observer == e.Value.(Observer) {
			s.observerCollection.Remove(e)
			return
		}
	}
}

func (s *Subject) NotifyObservers() {
	for e := s.observerCollection.Front(); e != nil; e = e.Next() {
		observer := e.Value.(Observer)
		observer.Notify()
	}
}

type ConcreteObserver struct{}

func (c *ConcreteObserver) Notify() {
	fmt.Println("ConcreteObserver.Notify()")
}

func main() {
	observer := new(ConcreteObserver)
	subject := NewSubject()
	subject.RegisterObserver(observer)
	subject.NotifyObservers()
	subject.UnregisterObserver(observer)
	subject.NotifyObservers()
}
