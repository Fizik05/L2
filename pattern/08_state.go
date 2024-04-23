package main

import "fmt"

// State interface
type State interface {
	DoAction(*Context)
}

// ConcreteStateA struct
type ConcreteStateA struct{}

// DoAction method for ConcreteStateA
func (s *ConcreteStateA) DoAction(context *Context) {
	fmt.Println("State A")
	context.state = &ConcreteStateB{}
}

// ConcreteStateB struct
type ConcreteStateB struct{}

// DoAction method for ConcreteStateB
func (s *ConcreteStateB) DoAction(context *Context) {
	fmt.Println("State B")
	context.state = &ConcreteStateA{}
}

// Context struct
type Context struct {
	state State
}

// Request method for Context
func (c *Context) Request() {
	c.state.DoAction(c)
}

func main() {
	context := &Context{}
	context.state = &ConcreteStateA{}

	context.Request()
	context.Request()
}
