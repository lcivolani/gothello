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
	board  *Board
	player rune
	steps  int
}

// InitialState specifies how the game is set up at the starting point.
func InitialState() *State {
	s := new(State)
	s.board = new(Board)
	mid := size / 2
	s.board.matrix[mid-1][mid-1] = 'X'
	s.board.matrix[mid-1][mid] = 'O'
	s.board.matrix[mid][mid-1] = 'O'
	s.board.matrix[mid][mid] = 'X'
	s.player = randomPlayer()
	return s
}

// Player defines which player has the move in a state.
// It returns the corresponding mark.
func (s *State) Player() rune {
	return s.player
}

// Opponent defines which player is playing against the current one.
// It returns the corresponding mark.
func (s *State) Opponent() rune {
	if s.Player() == 'X' {
		return 'O'
	}
	return 'X'
}

// Actions returns the set of legal moves in a state.
func (s *State) Actions() []Action {
	acts := make([]Action, 0)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// this operation cannot fail
			cell, _ := s.board.Cell(i, j)
			if cell != 0 {
				continue
			}
			if !s.captures(i, j) {
				continue
			}
			acts = append(acts, Action{s.player, i, j})
		}
	}
	return acts
}

// captures checks if the current player can capture any opponent pieces by placing
// a piece in the specified cell
func (s *State) captures(row, col int) bool {
	capt := false
	for _, dir := range directions {
		if s.capturesAlong(row, col, dir) {
			capt = true
			break
		}
	}
	return capt
}

// captures checks if the current player can capture any opponent pieces along the
// specified direction by placing a piece in the specified cell
func (s *State) capturesAlong(x, y int, dir *Direction) bool {
	nx, ny := dir.Next(x, y)
	ncell, ok := s.board.Cell(nx, ny)
	if !ok || ncell != s.Opponent() {
		return false
	}
	// walk over opponent's cells
	for ok && ncell == s.Opponent() {
		nx, ny = dir.Next(nx, ny)
		ncell, ok = s.board.Cell(nx, ny)
	}
	if !ok || ncell != s.player {
		return false
	}
	return true
}

func (s *State) Result(a Action) (*State, error) {
	if a.mark != s.player {
		return s, fmt.Errorf("invalid action %s: not %c's turn", a, a.mark)
	}
	res := s.Copy()
	err := res.board.SetCell(a.row, a.col, a.mark)
	if err != nil {
		return s, fmt.Errorf("invalid action %s: %v", a, err)
	}
	res.player = s.Opponent()
	res.steps++

	captured := 0
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
		if !ok || ncell != s.player {
			continue
		}
		// flip opponent's pieces (on result Board)
		ci, cj := dir.Next(a.row, a.col)
		for c := 0; c < count; c++ {
			res.board.Flip(ci, cj)
			ci, cj = dir.Next(ci, cj)
		}
		captured += count
	}
	if captured == 0 {
		return s, fmt.Errorf("invalid action %s: no capture", a)
	}
	fmt.Println("captured", captured, "pieces")
	return res, nil
}

func (s *State) Terminal() bool {
	// TODO: incorrect: state is terminal if *both* players cannot play
	return s.board.Count('X')+s.board.Count('O') == size*size ||
		len(s.Actions()) == 0
}

// Utility defines the final numeric value for a game that ends in the terminal
// state s, for a given player.
// It panics if called on a non-terminal state.
func (s *State) Utility(mark rune) int {
	if !s.Terminal() {
		panic("cannot compute utility: state is not terminal")
	}
	// TODO: by convention, empty squares at the end are added to the winner's score
	return s.board.Count(mark)
}

func (s *State) Copy() *State {
	copy := new(State)
	copy.board = s.board.Copy()
	copy.player = s.player
	copy.steps = s.steps
	return copy
}

func (s *State) String() string {
	buf := bytes.Buffer{}
	fmt.Fprint(&buf, s.board)
	fmt.Fprintf(&buf, "next player: %c\n", s.player)
	fmt.Fprintf(&buf, "steps: %d\n", s.steps)
	return buf.String()
}

func randomPlayer() rune {
	if rand.Intn(2) == 0 {
		return 'X'
	}
	return 'O'
}
