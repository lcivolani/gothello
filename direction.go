package main

type Direction struct {
	dx, dy int
}

func (d *Direction) Next(x, y int) (int, int) {
	return x + d.dx, y + d.dy
}

// TODO: the order in not relevant: consider changing to a set
var directions = []*Direction{
	&Direction{-1, 0}, // up
	&Direction{0, 1},  // right
	&Direction{1, 0},  // down
	&Direction{0, -1}, // left
}
