package main

import (
	"fmt"
)

type Component interface {
	Operation(int)
}

type File struct {
	name string
}

func (f *File) Operation(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println("File: " + f.name)
}

type Directory struct {
	components []Component
	name       string
}

func (d *Directory) Operation(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println("Directory: " + d.name)
	for _, component := range d.components {
		component.Operation(depth + 1)
	}
}

func (d *Directory) Add(component Component) {
	d.components = append(d.components, component)
}

func (d *Directory) Remove(component Component) {
	for i, v := range d.components {
		if v == component {
			d.remove(i)
			return
		}
	}
}

func (d *Directory) remove(i int) {
	reslice := append(d.components[:i], d.components[i+1:]...)
	newslice := make([]Component, len(reslice))
	copy(newslice, reslice)
	d.components = newslice
}

func (d *Directory) GetChildren() []Component {
	return d.components
}

func main() {
	root := &Directory{name: "root"}
	usr := &Directory{name: "usr"}
	local := &Directory{name: "local"}
	home := &Directory{name: "home"}
	user1 := &Directory{name: "user1"}
	file1 := &File{name: "file1"}
	root.Add(usr)
	usr.Add(local)
	root.Add(home)
	home.Add(user1)
	user1.Add(file1)
	root.Operation(0)
	root.Remove(home)
	root.Operation(0)
}
