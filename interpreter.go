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

// regular expressions to recognize operators 
var regularExpressions = []*regexp.Regexp{
	regexp.MustCompile("^-?\\d+$"),  // matches "2", "129832", "-843", etc.
	regexp.MustCompile("^\\+$"),     // matches "+"
	regexp.MustCompile("^\\*$"),     // matches "*"
	regexp.MustCompile("^-$"),       // matches "-"
	regexp.MustCompile("^/$"),       // matches "/"
	regexp.MustCompile("^%$"),       // matches "%"
	regexp.MustCompile("^squared$"), // matches "squared"
}

type function struct {
	stackParameters int                                  // the number of arguments needed from stack
	result          func(args []interface{}) interface{} // the function to execute to obtain the result
}

// corresponding functions to execute for literals and operators
var functions = []function{
	function{0, func(args []interface{}) interface{} { return args[0] }},                       // for int literals; returns its only argument
	function{2, func(args []interface{}) interface{} { return args[1].(int) + args[0].(int) }}, // addition
	function{2, func(args []interface{}) interface{} { return args[1].(int) * args[0].(int) }}, // multiplication
	function{2, func(args []interface{}) interface{} { return args[1].(int) - args[0].(int) }}, // subtraction
	function{2, func(args []interface{}) interface{} { return args[1].(int) / args[0].(int) }}, // division
	function{2, func(args []interface{}) interface{} { return args[1].(int) % args[0].(int) }}, // modulo
	function{1, func(args []interface{}) interface{} { return args[0].(int) * args[0].(int) }}, // squaring
}

// the global stack
var s *stack.Stack

func init() {
	s = stack.New()
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
		panic("stack not empty!")
	}
}

func eval(x string) {
	// choose correct function
	var f function

	for i, e := range regularExpressions {
		if e.MatchString(x) {
			f = functions[i]
		}
	}

	args := []interface{}{}

	// if x is an int literal, fix type to int and add x itself as argument
	if regularExpressions[0].MatchString(x) {
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
