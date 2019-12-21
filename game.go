package main

import "fmt"

func NewGame(p1, p2 Player) *Game {
	return &Game{
		players: map[rune]Player{
			'X': p1,
			'O': p2,
		},
		state: InitialState(),
	}
}

type Game struct {
	players map[rune]Player
	state   *State
}

func (g *Game) Play() {
	for {
		if g.state.Terminal() {
			break
		}

		player := g.players[g.state.Player()]
		action := player.Play(g.state)

		nstate, err := g.state.Result(action)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		g.state = nstate
	}

	fmt.Println(g.state)
	fmt.Println("game over!")
	fmt.Println("X's score:", g.state.Utility('X'))
	fmt.Println("O's score:", g.state.Utility('O'))
}
