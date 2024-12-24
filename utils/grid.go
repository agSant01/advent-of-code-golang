package utils

import (
	"fmt"
)

type Coord struct {
	X int
	Y int
}

func (c *Coord) Displaces(vector Coord) {
	c.X += vector.X
	c.Y += vector.Y
}

func (c *Coord) Add(vector Coord) Coord {
	return Coord{
		X: c.X + vector.X,
		Y: c.Y + vector.Y,
	}
}

func (c *Coord) Copy() Coord {
	return Coord{
		X: c.X,
		Y: c.Y,
	}
}

var ARROW_DIRECTIONS map[string]Coord = map[string]Coord{
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
	"^": {0, -1},
}

func PrettyPrintGrid[T1 any](grid *[][]T1) {
	height := len(*grid)
	width := len((*grid)[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print((*grid)[y][x])
		}
		fmt.Println()
	}
}

func SwapGridValues(warehouse *[][]string, a Coord, b Coord) {
	tmp := (*warehouse)[a.Y][a.X]
	(*warehouse)[a.Y][a.X] = (*warehouse)[b.Y][b.X]
	(*warehouse)[b.Y][b.X] = tmp
}

func GetCoordValue(territory *[][]string, coord Coord) (string, bool) {
	width := len((*territory)[0])
	height := len(*territory)

	if coord.X >= 0 && coord.Y >= 0 && coord.X < width && coord.Y < height {
		return (*territory)[coord.Y][coord.X], true
	}

	return "", false
}

func GetCoordValueInt8(territory *[][]int8, coord Coord) (int8, bool) {
	width := len((*territory)[0])
	height := len(*territory)

	if coord.X >= 0 && coord.Y >= 0 && coord.X < width && coord.Y < height {
		return (*territory)[coord.Y][coord.X], true
	}

	return -1, false
}

func CopyGrid[T1 any](grid *[][]T1) *[][]T1 {
	newGrid := make([][]T1, len(*grid))
	for i, inner := range *grid {
		newInner := make([]T1, len(inner))
		copy(newInner, inner) // Use the built-in copy function
		newGrid[i] = newInner
	}
	return &newGrid
}

func GetNeighbors(node Coord, width int, height int) []Coord {
	toReturn := []Coord{}
	for _, v := range []Coord{{node.X - 1, node.Y}, {node.X, node.Y - 1}, {node.X + 1, node.Y}, {node.X, node.Y + 1}} {
		if v.X < 0 || v.Y < 0 || v.X >= width || v.Y >= height {
			continue
		}
		toReturn = append(toReturn, v)
	}
	return toReturn
}
