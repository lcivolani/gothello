package main

func main() {
	p1 := &ComputerPlayer{Name: "Alice"}
	p2 := &ComputerPlayer{Name: "Bob"}

	game := NewGame(p1, p2)
	game.Play()
}
