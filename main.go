package main

func main() {
	p1 := &RandomPlayer{"Aldo"}
	p2 := &RandomPlayer{"Ugo"}
	g := NewGame(p1, p2)
	g.Play()
}
