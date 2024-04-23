package main

import "fmt"

// Тип продукта, который мы строим
type Product struct {
	part1 string
	part2 string
	part3 string
}

// Интерфейс строителя, определяющий шаги создания продукта
type Builder interface {
	BuildPart1() Builder
	BuildPart2() Builder
	BuildPart3() Builder
	GetProduct() *Product
}

// Конкретный строитель для создания конкретной конфигурации продукта
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (cb *ConcreteBuilder) BuildPart1() Builder {
	cb.product.part1 = "Part 1"
	return cb
}

func (cb *ConcreteBuilder) BuildPart2() Builder {
	cb.product.part2 = "Part 2"
	return cb
}

func (cb *ConcreteBuilder) BuildPart3() Builder {
	cb.product.part3 = "Part 3"
	return cb
}

func (cb *ConcreteBuilder) GetProduct() *Product {
	return cb.product
}

// Директор, который управляет процессом построения продукта
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Product {
	return d.builder.
		BuildPart1().
		BuildPart2().
		BuildPart3().
		GetProduct()
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	product := director.Construct()

	fmt.Println("Product:", *product)
}
