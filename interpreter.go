// A calculator for reversed polish notation, using function composition.
// This was an exercise of concatenative programming; it only supports integer numbers. It is inspired by http://evincarofautumn.blogspot.mx/2012/02/why-concatenative-programming-matters.html.
// For more information see http://en.wikipedia.org/wiki/Reverse_Polish_notation.
package main

import (
	"fmt"
	"github.com/sauerbraten/stack"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type function struct {
	stackParameters int                                  // the number of arguments needed from stack
	result          func(args []interface{}) interface{} // the function to execute to obtain the result
}

// regular expression to ensure it's an int literal
var intExp = regexp.MustCompile("^-?\\d+$") // matches "2", "129832", "-843", etc.

// map of functions for operators
var functions map[string]function

// the global stack
var s *stack.Stack

func init() {
	// set up the stack
	s = stack.New()

	// set up the functions; one function per operator
	functions = make(map[string]function)
	functions["+"] = function{2, func(args []interface{}) interface{} { return args[1].(int) + args[0].(int) }}       // addition
	functions["*"] = function{2, func(args []interface{}) interface{} { return args[1].(int) * args[0].(int) }}       // multiplication
	functions["-"] = function{2, func(args []interface{}) interface{} { return args[1].(int) - args[0].(int) }}       // subtraction
	functions["/"] = function{2, func(args []interface{}) interface{} { return args[1].(int) / args[0].(int) }}       // division
	functions["%"] = function{2, func(args []interface{}) interface{} { return args[1].(int) % args[0].(int) }}       // modulo
	functions["squared"] = function{1, func(args []interface{}) interface{} { return args[0].(int) * args[0].(int) }} // squaring
}

func main() {
	// os.Args[0] is the program path, os.Args[1] the actual input string
	for _, v := range strings.Split(os.Args[1], " ") {
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

func eval(x string) {
	// the arguments we will later pass to the function
	args := []interface{}{}

	// choose correct function
	f, ok := functions[x]
	if !ok {
		// probably dealing with an int literal, but let's make sure
		if !intExp.MatchString(x) {
			// invalid input
			fmt.Fprintln(os.Stderr, "invalid input:", x)
			os.Exit(1)
		}

		f = function{0, func(args []interface{}) interface{} { return args[0] }} // for int literals; returns its only argument

		// because x is an int literal, convert to int and add it as argument
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}

		args = append(args, i)
	}

	// pop needed arguments from stack
	for i := 0; i < f.stackParameters; i++ {
		x, err := s.Pop()
		if err != nil {
			panic(err)
		}

		args = append(args, x)
	}

	// push result back onto stack
	s.Push(f.result(args))
}
