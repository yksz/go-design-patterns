package main

import (
	"fmt"
)

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type ConcreteIterator struct {
	concreteAggregate *ConcreteAggregate
	index             int
}

func (i *ConcreteIterator) Next() interface{} {
	if i.index >= len(i.concreteAggregate.slice) {
		return nil
	}
	defer func() {
		i.index++
	}()
	return i.concreteAggregate.slice[i.index]
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.concreteAggregate.slice)
}

type Aggregate interface {
	Iterator() Iterator
}

type ConcreteAggregate struct {
	slice []string
}

func (a *ConcreteAggregate) Iterator() Iterator {
	return &ConcreteIterator{concreteAggregate: a}
}

func main() {
	slice := []string{"a", "b", "c", "d"}
	aggregate := &ConcreteAggregate{slice}
	for i := aggregate.Iterator(); i.HasNext(); {
		v := i.Next()
		fmt.Println(v)
	}
}
