package main

import (
	"fmt"
)

func main() {
	s := InitialState()
	fmt.Println(s)
	fmt.Println("utility for X is", s.Utility('X'))
}
