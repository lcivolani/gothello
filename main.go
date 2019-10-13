package main

import (
	"fmt"
	"os"
)

func main() {
	s := InitialState()
	fmt.Print(s)
	for i := 0; i < 5; i++ {
		acts := s.Actions()
		fmt.Println("actions:", acts)
		fmt.Println("going to execute the first one...")
		ns, err := s.Result(acts[0])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		s = ns
		fmt.Print(s)
	}
}
