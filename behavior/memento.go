package main

import (
	"fmt"
)

type Originator struct {
	state string
}

func (o *Originator) Set(state string) {
	fmt.Println("Setting state to " + state)
	o.state = state
}

func (o *Originator) SaveToMemento() Memento {
	fmt.Println("Saving to Memento.")
	return Memento{o.state}
}

func (o *Originator) RestoreFromMemento(memento Memento) {
	o.state = memento.GetSavedState()
	fmt.Println("State after restoring from Memento: " + o.state)
}

type Memento struct {
	state string
}

func (m *Memento) GetSavedState() string {
	return m.state
}

func main() {
	savedStates := make([]Memento, 0)
	originator := new(Originator)
	originator.Set("State1")
	originator.Set("State2")
	savedStates = append(savedStates, originator.SaveToMemento())
	originator.Set("State3")
	savedStates = append(savedStates, originator.SaveToMemento())
	originator.Set("State4")
	originator.RestoreFromMemento(savedStates[1])
}
