package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	Name  string
	Delay time.Duration
}

// Play generates a random move and returns it to the caller.
func (p *RandomPlayer) Play(s *State) Action {
	fmt.Println(s)
	actions := s.Actions()
	time.Sleep(p.Delay * time.Second)
	i := rand.Intn(len(actions))
	fmt.Printf("%s picks %v randomly from %v\n", p.Name, actions[i], actions)
	return actions[i]
}

func (p *RandomPlayer) String() string {
	return fmt.Sprintf("random player %s", p.Name)
}

type ComputerPlayer struct {
	Name string
	mark rune
}

// Play generates a clever move and returns it to the caller.
func (p *ComputerPlayer) Play(s *State) Action {
	fmt.Println(s)
	p.mark = s.Player()
	return p.minimax(s)
}

func (p *ComputerPlayer) minimax(s *State) Action {
	actions := s.Actions()
	best := 0
	bestVal := math.MinInt32
	for i := range s.Actions() {
		ns, _ := s.Result(actions[i])
		val := p.minValue(ns)
		fmt.Println("action", actions[i], "would lead to", val)
		if val > bestVal {
			best = i
			bestVal = val
		}
	}
	return actions[best]
}

func (p *ComputerPlayer) minValue(s *State) int {
	if s.Terminal() {
		return s.Utility(p.mark)
	}
	min := math.MaxInt32
	for _, a := range s.Actions() {
		ns, _ := s.Result(a)
		if val := p.maxValue(ns); val < min {
			min = val
		}
	}
	return min
}

func (p *ComputerPlayer) maxValue(s *State) int {
	if s.Terminal() {
		return s.Utility(p.mark)
	}
	max := math.MinInt32
	for _, a := range s.Actions() {
		ns, _ := s.Result(a)
		if val := p.minValue(ns); val > max {
			max = val
		}
	}
	return max
}
