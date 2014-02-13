package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Number struct {
	number int
}

func (n *Number) Interpret(variables map[string]Expression) int {
	return n.number
}

type Plus struct {
	leftOperand  Expression
	rightOperand Expression
}

func (p *Plus) Interpret(variables map[string]Expression) int {
	return p.leftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

type Minus struct {
	leftOperand  Expression
	rightOperand Expression
}

func (m *Minus) Interpret(variables map[string]Expression) int {
	return m.leftOperand.Interpret(variables) - m.rightOperand.Interpret(variables)
}

type Variable struct {
	name string
}

func (v *Variable) Interpret(variables map[string]Expression) int {
	if variables[v.name] == nil {
		return 0
	}
	return variables[v.name].Interpret(variables)
}

type Evaluator struct {
	syntaxTree Expression
}

func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		if token == "+" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		} else if token == "-" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Minus{left, right}
			expressionStack.Push(subExpression)
		} else {
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree}
}

func (e *Evaluator) Interpret(context map[string]Expression) int {
	return e.syntaxTree.Interpret(context)
}

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	top  *Element
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

func main() {
	expression := "w x z - +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Number{5}
	variables["x"] = &Number{10}
	variables["z"] = &Number{42}
	result := sentence.Interpret(variables)
	fmt.Println(result)
}
