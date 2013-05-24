// A calculator for reversed polish notation, using function composition.
// This was an exercise of concatenative programming; it only supports integer numbers. It is inspired by http://evincarofautumn.blogspot.mx/2012/02/why-concatenative-programming-matters.html.
// For more information see http://en.wikipedia.org/wiki/Reverse_Polish_notation.
package main

import (
	"fmt"
	"github.com/sauerbraten/stack"
	"os"
)

// the global stack
var s *stack.Stack = stack.New()

func main() {
	fmt.Println(os.Args)
	// os.Args[0] is the program path, os.Args[1] the actual input string
	for _, v := range os.Args[1:] {
		eval(v)
	}

	// final result is should be the only element now on the stack
	result, err := s.Pop()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// make sure stack is now empty
	_, err = s.Pop()
	if err == nil {
		fmt.Fprintln(os.Stderr, "stack not empty!")
	}
}
