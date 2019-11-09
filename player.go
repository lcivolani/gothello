package main

import (
	"fmt"
	"math/rand"
)

// Player is an interface that wraps the basic Play method.
type Player interface {
	Play(s *State) Action
}

// HumanPlayer plays by asking for a move to the user through the command line.
type HumanPlayer struct {
	Name string
}

// Play asks for a move through the command line and returns it to the caller.
func (p *HumanPlayer) Play(s *State) Action {
	fmt.Println(s)
	var row, col int
	for {
		fmt.Printf("%s (%q), insert move (row col): ", p.Name, s.Player())
		n, _ := fmt.Scanf("%d %d", &row, &col)
		if n == 2 {
			break
		}
	}
	return Action{s.Player(), row, col}
}

func (p *HumanPlayer) String() string {
	return fmt.Sprintf("human player %s", p.Name)
}

// RandomPlayer plays by issuing random moves.
type RandomPlayer struct {
	Name string
}

// Play generates a random move and returns it to the caller.
func (p *RandomPlayer) Play(s *State) Action {
	fmt.Println(s)
	actions := s.Actions()
	i := rand.Intn(len(actions))
	fmt.Printf("%s picks %v randomly from %v\n", p.Name, actions[i], actions)
	return actions[i]
}

func (p *RandomPlayer) String() string {
	return fmt.Sprintf("random player %s", p.Name)
}
