package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const size = 8

type Board struct {
	matrix [size][size]rune
}

func (b Board) Cell(i, j int) rune {
	// no need to check i and j: the runtime will do it for us, and panic if necessary
	return b.matrix[i][j]
}

func (b Board) SetCell(i, j int, mark rune) (Board, error) {
	if err := checkCoord(i, j); err != nil {
		return b, fmt.Errorf("cannot set cell (%d, %d): %v", i, j, err)
	}
	if b.matrix[i][j] != 0 {
		return b, fmt.Errorf("cannot set cell (%d, %d): cell is full", i, j)
	}
	b.matrix[i][j] = mark
	return b, nil
}

func (b Board) Flip(i, j int) Board {
	switch b.matrix[i][j] {
	case 'X':
		b.matrix[i][j] = 'O'
	case 'O':
		b.matrix[i][j] = 'X'
	default:
		// only the program can flip cells: an error here must be due to a bug
		panic("attempt to flip an empty cell")
	}
	return b
}

func (b Board) Count(mark rune) int {
	if mark != 'X' && mark != 'O' && mark != 0 {
		panic(fmt.Sprintf("unknown mark %q", mark))
	}
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.matrix[i][j] == mark {
				count++
			}
		}
	}
	return count
}

func (b Board) String() string {
	buf := bytes.Buffer{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j > 0 {
				fmt.Fprint(&buf, " ")
			}
			if b.matrix[i][j] == 0 {
				fmt.Fprint(&buf, ".")
			} else {
				fmt.Fprintf(&buf, "%c", b.matrix[i][j])
			}
		}
		fmt.Fprint(&buf, "\n")
	}
	return buf.String()
}

func checkCoord(i, j int) error {
	if i < 0 || i >= size {
		return fmt.Errorf("row %d out of range", i)
	}
	if j < 0 || j >= size {
		return fmt.Errorf("col %d out of range", j)
	}
	return nil
}

func parseGrid(grid string) (Board, error) {
	chars := make([]rune, 0, size*size)
	for _, c := range grid {
		if strings.ContainsRune("XO.", c) {
			chars = append(chars, c)
		}
	}
	if len(chars) != size*size {
		return Board{}, errors.New("parsing failed: malformed grid")
	}
	var matrix [size][size]rune
	for i, c := range chars {
		if c == '.' {
			c = 0
		}
		row := i / size
		col := i % size
		matrix[row][col] = c
	}
	return Board{matrix}, nil
}

func mustParseGrid(grid string) Board {
	b, err := parseGrid(grid)
	if err != nil {
		panic(err)
	}
	return b
}
