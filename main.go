package main

import (
	"fmt"
	"os"
)

var grid = `
........
........
........
..X.....
...O...X
....O.O.
........
........
`

func main() {
	s := InitialState()
	s.board = mustParseGrid(grid)
	fmt.Print(s)
	for i := 0; i < 5; i++ {
		acts := s.Actions()
		fmt.Println("actions:", acts)
		if len(acts) == 0 {
			fmt.Println("no possible actions")
			break
		}
		fmt.Println("going to execute the first one")
		ns, err := s.Result(acts[0])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		s = ns
		fmt.Print(s)
	}
}
