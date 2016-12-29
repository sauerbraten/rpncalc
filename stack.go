package main

import "errors"

type stack []float64

// Push puts a new element on top of the stack
func (s *stack) Push(value float64) {
	*s = append(*s, value)
}

// Pop removes the top element and returns it, along with an error that is only != nil if the stack is empty
func (s *stack) Pop() (float64, error) {
	temp, err := s.Peek()
	if err != nil {
		return 0.0, err
	}

	*s = (*s)[0 : len(*s)-1]

	return temp, nil
}

func (s *stack) Peek() (float64, error) {
	if len(*s) == 0 {
		return 0.0, errors.New("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

// Len returns the amount of elements in the stack
func (s *stack) Len() int {
	return len(*s)
}
