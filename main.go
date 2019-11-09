package main

func main() {
	p1 := &HumanPlayer{"Aldo"}
	p2 := &HumanPlayer{"Ugo"}
	g := NewGame(p1, p2)
	g.Play()
}
