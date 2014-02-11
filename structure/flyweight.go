package main

import (
	"fmt"
)

type FlyweightFactory struct {
	pool map[string]*Flyweight
}

func (f *FlyweightFactory) GetFlyweight(str string) *Flyweight {
	flyweight := f.pool[str]
	if flyweight == nil {
		fmt.Println("new: " + str)
		f.pool[str] = &Flyweight{str}
	}
	return flyweight
}

type Flyweight struct {
	str string
}

func main() {
	factory := &FlyweightFactory{make(map[string]*Flyweight)}
	factory.GetFlyweight("a")
	factory.GetFlyweight("p")
	factory.GetFlyweight("p")
	factory.GetFlyweight("l")
	factory.GetFlyweight("e")
	factory.GetFlyweight("!")
	factory.GetFlyweight("o")
	factory.GetFlyweight("r")
	factory.GetFlyweight("a")
	factory.GetFlyweight("n")
	factory.GetFlyweight("g")
	factory.GetFlyweight("e")
	factory.GetFlyweight("!")
}
