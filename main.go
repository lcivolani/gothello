package main

import "fmt"

func main() {
	grid := `
	. . . . . . . .
	. . . . X . . .
	. . . O . . . .
	. . O . . . . .
	. . . . . X O O
	. . . . . O . .
	. . O O X . . .
	. . . . . . . .`
	board := mustParseGrid(grid)
	state := State{board, 'X', 0}
	fmt.Println(board)
	fmt.Println(state.Actions())
}
