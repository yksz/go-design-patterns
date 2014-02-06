package main

import (
	"fmt"
)

type DrawingAPI interface {
	DrawCircle(x, y, radius float32)
}

type DrawingAPI1 struct{}

func (a *DrawingAPI1) DrawCircle(x, y, radius float32) {
	fmt.Printf("API1.circle at %f:%f radius %f\n", x, y, radius)
}

type DrawingAPI2 struct{}

func (a *DrawingAPI2) DrawCircle(x, y, radius float32) {
	fmt.Printf("API2.circle at %f:%f radius %f\n", x, y, radius)
}

type Shape interface {
	Draw()
	ResizeByPercentage(pct float32)
}

type CircleShape struct {
	x, y, radius float32
	drawingAPI   DrawingAPI
}

func (s *CircleShape) Draw() {
	s.drawingAPI.DrawCircle(s.x, s.y, s.radius)
}

func (s *CircleShape) ResizeByPercentage(pct float32) {
	s.radius *= pct
}

func main() {
	shapes := []Shape{&CircleShape{1, 2, 3, new(DrawingAPI1)}, &CircleShape{5, 7, 11, new(DrawingAPI2)}}
	for _, shape := range shapes {
		shape.ResizeByPercentage(2.5)
		shape.Draw()
	}
}
