package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type function struct {
	stackParameters int                                  // the number of arguments needed from stack
	result          func(args []interface{}) interface{} // the function to execute to obtain the result
}

// regular expression to ensure it's an int literal
var intExp = regexp.MustCompile("^-?\\d+$") // matches "2", "129832", "-843", etc.

// map of functions for operators
//
// example explanation for the + function:
// stackParameters = 2 means eval() will pass it the top 2 values from the stack
// the result function simply casts the first two values from the slice of values it got passed to int, adds them up, and returns the result
//
// here is a formatted version of the + line:
//	"+": function{
//		stackParameters: 2,
//		result: func(args []interface{}) interface{} {
//			return args[1].(int) + args[0].(int)
//		},
//	},
var functions map[string]function = map[string]function{
	"+":       function{2, func(args []interface{}) interface{} { return args[1].(int) + args[0].(int) }}, // addition
	"·":       function{2, func(args []interface{}) interface{} { return args[1].(int) * args[0].(int) }}, // multiplication
	"-":       function{2, func(args []interface{}) interface{} { return args[1].(int) - args[0].(int) }}, // subtraction
	"/":       function{2, func(args []interface{}) interface{} { return args[1].(int) / args[0].(int) }}, // division
	"%":       function{2, func(args []interface{}) interface{} { return args[1].(int) % args[0].(int) }}, // modulo
	"squared": function{1, func(args []interface{}) interface{} { return args[0].(int) * args[0].(int) }}, // squaring
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
