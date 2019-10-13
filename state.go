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

// Opponent defines which player is playing against the current one.
// It returns the corresponding mark.
func (s State) Opponent() rune {
	if s.Player() == 'X' {
		return 'O'
	}
	return 'X'
}

// Actions returns the set of legal moves in a state.
func (s State) Actions() []Action {
	// TODO: possible duplicates!
	acts := make([]Action, 0)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// this operation cannot fail
			cell, _ := s.board.Cell(i, j)
			if cell != s.toMove {
				continue
			}
			for _, dir := range directions {
				ni, nj := dir.Next(i, j)
				ncell, ok := s.board.Cell(ni, nj)
				if !ok || ncell != s.Opponent() {
					continue
				}
				// walk over opponent's cells
				for ok && ncell == s.Opponent() {
					ni, nj = dir.Next(ni, nj)
					ncell, ok = s.board.Cell(ni, nj)
				}
				if !ok || ncell != 0 {
					continue
				}
				acts = append(acts, Action{s.toMove, ni, nj})
			}
		}
	}
	return acts
}

func (s State) Result(a Action) (State, error) {
	if a.mark != s.toMove {
		return s, fmt.Errorf("not %c's turn", a.mark)
	}
	nb, err := s.board.SetCell(a.row, a.col, a.mark)
	if err != nil {
		return s, fmt.Errorf("invalid action %s: %v", a, err)
	}
	for _, dir := range directions {
		ni, nj := dir.Next(a.row, a.col)
		ncell, ok := s.board.Cell(ni, nj)
		if !ok || ncell != s.Opponent() {
			continue
		}
		count := 0
		for ok && ncell == s.Opponent() {
			count++
			ni, nj = dir.Next(ni, nj)
			ncell, ok = s.board.Cell(ni, nj)
		}
		if !ok || ncell != s.toMove {
			continue
		}
		// TODO: flip opponent's pieces
	}
	s.board = nb
	s.toMove = s.Opponent()
	s.steps++
	return s, nil
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
