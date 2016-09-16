package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type function struct {
	numParams int                  // the number of arguments needed from stack
	result    func(args []int) int // the function to execute to obtain the result
}

// regular expression to ensure it's an int literal
var intExp = regexp.MustCompile("^-?\\d+$") // matches "2", "129832", "-843", etc.

// map of functions for operators
//
// example explanation for the + function:
// numParams = 2 means eval() will pass it the top 2 values from the stack
// the result function simply casts the first two values from the slice of values it got passed to int, adds them up, and returns the result
//
// here is a formatted version of the + line:
//	"+": function{
//		numParams: 2,
//		result: func(args []int) int {
//			return args[1] + args[0]
//		},
//	},
var functions = map[string]function{
	"+":       function{2, func(args []int) int { return args[1] + args[0] }}, // addition
	"·":       function{2, func(args []int) int { return args[1] * args[0] }}, // multiplication
	"-":       function{2, func(args []int) int { return args[1] - args[0] }}, // subtraction
	"/":       function{2, func(args []int) int { return args[1] / args[0] }}, // division
	"%":       function{2, func(args []int) int { return args[1] % args[0] }}, // modulo
	"squared": function{1, func(args []int) int { return args[0] * args[0] }}, // squaring
}

func eval(x string) {
	// the arguments we will later pass to the function
	args := []int{}

	// choose correct function
	f, ok := functions[x]
	if !ok {
		// probably dealing with an int literal, but let's make sure
		if !intExp.MatchString(x) {
			// invalid input
			fmt.Fprintln(os.Stderr, "invalid input:", x)
			os.Exit(1)
		}

		f = function{0, func(args []int) int { return args[0] }} // for int literals; returns its only argument

		// because x is an int literal, convert to int and add it as argument
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}

		args = append(args, i)
	}

	// pop needed arguments from stack
	for i := 0; i < f.numParams; i++ {
		x, err := s.Pop()
		if err != nil {
			panic(err)
		}

		args = append(args, x.(int))
	}

	// push result back onto stack
	s.Push(f.result(args))
}
