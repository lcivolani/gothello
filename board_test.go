package main

import (
	"fmt"
	"testing"
)

var testBoard Board

func init() {
	grid := `
. . . . . . . .
. . . . . . . .
. . . . . . . .
. . . X O . . .
. . . O X . . .
. . . . . . . .
. . . . . . . .
. . . . . . . .`
	testBoard = mustParseGrid(grid)
}

func TestCell(t *testing.T) {
	tests := []struct {
		row  int
		col  int
		want rune
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{3, 3, 'X'},
		{3, 4, 'O'},
		{4, 3, 'O'},
		{4, 4, 'X'},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("testBoard.Cell(%d, %d)", test.row, test.col)
		got := testBoard.Cell(test.row, test.col)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		mark rune
		want int
	}{
		{'X', 2},
		{'O', 2},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("testBoard.Count(%q)", test.mark)
		got := testBoard.Count(test.mark)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}
