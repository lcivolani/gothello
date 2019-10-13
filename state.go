package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// State represents a stage of an Othello gameplay.
type State struct {
	board  Board
	toMove rune
	steps  int
}

// InitialState specifies how the game is set up at the starting point.
func InitialState() State {
	s := State{}
	mid := size / 2
	s.board.matrix[mid-1][mid-1] = 'X'
	s.board.matrix[mid-1][mid] = 'O'
	s.board.matrix[mid][mid-1] = 'O'
	s.board.matrix[mid][mid] = 'X'
	if rand.Intn(2) == 0 {
		s.toMove = 'X'
	} else {
		s.toMove = 'O'
	}
	return s
}

// Player defines which player has the move in a state.
// It returns the corresponding mark.
func (s State) Player() rune {
	return s.toMove
}

// Utility defines the final numeric value for a game that ends in the terminal
// state s, for a given player.
// It panics if called on a non-terminal state.
func (s State) Utility(mark rune) int {
	// TODO: make sure that the state is terminal
	return s.board.Count(mark)
}

func (s State) String() string {
	buf := bytes.Buffer{}
	fmt.Fprint(&buf, s.board)
	fmt.Fprintf(&buf, "next player: %c\n", s.toMove)
	fmt.Fprintf(&buf, "total moves: %d\n", s.steps)
	return buf.String()
}
