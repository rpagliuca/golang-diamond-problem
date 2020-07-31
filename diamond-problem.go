package main

import "fmt"

// This file is not supposed to compile.
// Expected error: ambiguous selector t.Method

// See https://github.com/rpagliuca/golang-solving-the-diamond-problem
// for a solution

func main() {
	t := ComposedType{}
	t.Method() // The problem is here
}

type Type1 struct{}

type Type2 struct{}

type ComposedType struct {
	Type1
	Type2
}

func (Type1) Method() {
	fmt.Println("Called Type1.Method")
}

func (Type2) Method() {
	fmt.Println("Called Type2.Method")
}
