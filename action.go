package main

import "fmt"

// Action represents a move made by a player over the course of a game.
type Action struct {
	mark     rune
	row, col int
}

func (a Action) String() string {
	return fmt.Sprintf("%q in (%d, %d)", a.mark, a.row, a.col)
}
