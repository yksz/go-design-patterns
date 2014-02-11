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

func (b *AppleBuilder) SetName() {
	b.product.name = "apple"
}

func (b *AppleBuilder) SetPrice() {
	b.product.price = 10
}

func (b *AppleBuilder) GetResult() *Product {
	return b.product
}

type OrangeBuilder struct {
	product *Product
}

func NewOrangeBuilder() *OrangeBuilder {
	return &OrangeBuilder{new(Product)}
}

func (b *OrangeBuilder) SetName() {
	b.product.name = "orange"
}

func (b *OrangeBuilder) SetPrice() {
	b.product.price = 20
}

func (b *OrangeBuilder) GetResult() *Product {
	return b.product
}

func main() {
	director1 := &Director{NewAppleBuilder()}
	director2 := &Director{NewOrangeBuilder()}
	product1 := director1.Construct()
	product2 := director2.Construct()
	fmt.Println(product1)
	fmt.Println(product2)
}
