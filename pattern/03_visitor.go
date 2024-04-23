package main

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
}

func (cea *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(cea)
}

type ConcreteElementB struct {
}

func (ceb *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ceb)
}

type Visitor interface {
	VisitConcreteElementA(el *ConcreteElementA)
	VisitConcreteElementB(el *ConcreteElementB)
}

type ConcreteVisitor struct {
}

func (cv *ConcreteVisitor) VisitConcreteElementA(cea *ConcreteElementA) {
	fmt.Println("Element A")
}

func (cv *ConcreteVisitor) VisitConcreteElementB(ceb *ConcreteElementB) {
	fmt.Println("Element B")
}

type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(el Element) {
	os.elements = append(os.elements, el)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, el := range os.elements {
		el.Accept(visitor)
	}
}

func main() {
	ObjectStructure := &ObjectStructure{}

	ObjectStructure.Attach(&ConcreteElementA{})
	ObjectStructure.Attach(&ConcreteElementB{})

	visitor := &ConcreteVisitor{}

	ObjectStructure.Accept(visitor)
}
