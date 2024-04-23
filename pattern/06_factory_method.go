package main

import "fmt"

// Product interface
type Product interface {
	Use() string
}

// ConcreteProduct struct
type ConcreteProduct struct{}

// Use method for ConcreteProduct
func (p *ConcreteProduct) Use() string {
	return "Using ConcreteProduct"
}

// Creator interface
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreator struct
type ConcreteCreator struct{}

// CreateProduct method for ConcreteCreator
func (c *ConcreteCreator) CreateProduct() Product {
	return &ConcreteProduct{}
}

func main() {
	creator := &ConcreteCreator{}
	product := creator.CreateProduct()
	fmt.Println(product.Use())
}
