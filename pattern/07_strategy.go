package main

import "fmt"

// Strategy interface
type Strategy interface {
	DoOperation(int, int) int
}

// ConcreteStrategyAdd struct
type ConcreteStrategyAdd struct{}

// DoOperation method for ConcreteStrategyAdd
func (s *ConcreteStrategyAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// ConcreteStrategySubtract struct
type ConcreteStrategySubtract struct{}

// DoOperation method for ConcreteStrategySubtract
func (s *ConcreteStrategySubtract) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// Context struct
type Context struct {
	strategy Strategy
}

// ExecuteStrategy method for Context
func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	context := &Context{}

	addStrategy := &ConcreteStrategyAdd{}
	context.strategy = addStrategy
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	subtractStrategy := &ConcreteStrategySubtract{}
	context.strategy = subtractStrategy
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}
