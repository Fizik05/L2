package main

import "fmt"

type Subsystem1 struct {
}

func (s *Subsystem1) operation1() {
	fmt.Println("The first subsystem do something")
}

type Subsystem2 struct {
}

func (s *Subsystem2) operation2() {
	fmt.Println("The second subsystem do something")
}

type Facade struct {
	Subsystem1
	Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		Subsystem1: Subsystem1{},
		Subsystem2: Subsystem2{},
	}
}

func (f *Facade) operation() {
	f.Subsystem1.operation1()
	f.Subsystem2.operation2()
}

func main() {
	facade := NewFacade()

	facade.operation()
}
