package main

import (
	"fmt"
)

func main() {
	grid := `
	. . . . . . . .
	. . . . . . . .
	. . . . . . . .
	. . . X O . . .
	. . . O X . . .
	. . . . . . . .
	. . . . . . . .
	. . . . . . . .
	`
	board := mustParseGrid(grid)

	fmt.Println(board)
}
