package parser

import (
	"fmt"
)

// copied from https://golangbyexample.com/floatstack-in-golang/
type floatstack struct {
	stack []float64
}

func (c *floatstack) push(value float64) {
	c.stack = append(c.stack, value)
}

func (c *floatstack) pop() (float64, error) {
	if len(c.stack) > 0 {
		ele := c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
		return ele, nil
	}
	return 0, fmt.Errorf("Pop Error: Stack is empty")
}

func (c *floatstack) front() (float64, error) {
	if len(c.stack) > 0 {
		val := c.stack[len(c.stack)-1]
		return val, nil
	}
	return 0, fmt.Errorf("Peep Error: Stack is empty")
}

func (c *floatstack) size() int {
	return len(c.stack)
}

func (c *floatstack) empty() bool {
	return c.size() == 0
}

func fs() floatstack {
	return floatstack{
		stack: make([]float64, 0),
	}
}
