package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	name  string
	price int
}

func (p *Product) String() string {
	return "Product [name=" + p.name + ", price=" + strconv.Itoa(p.price) + "]"
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() *Product {
	d.builder.SetName()
	d.builder.SetPrice()
	return d.builder.GetResult()
}

type Builder interface {
	SetName()
	SetPrice()
	GetResult() *Product
}

type AppleBuilder struct {
	product *Product
}

func NewAppleBuilder() *AppleBuilder {
	return &AppleBuilder{new(Product)}
}

func (a *AppleBuilder) SetName() {
	a.product.name = "apple"
}

func (a *AppleBuilder) SetPrice() {
	a.product.price = 10
}

func (a *AppleBuilder) GetResult() *Product {
	return a.product
}

type OrangeBuilder struct {
	product *Product
}

func NewOrangeBuilder() *OrangeBuilder {
	return &OrangeBuilder{new(Product)}
}

func (o *OrangeBuilder) SetName() {
	o.product.name = "orange"
}

func (o *OrangeBuilder) SetPrice() {
	o.product.price = 20
}

func (o *OrangeBuilder) GetResult() *Product {
	return o.product
}

func main() {
	director1 := &Director{NewAppleBuilder()}
	director2 := &Director{NewOrangeBuilder()}
	product1 := director1.Construct()
	product2 := director2.Construct()
	fmt.Println(product1)
	fmt.Println(product2)
}
