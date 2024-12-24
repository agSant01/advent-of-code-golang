package utils

import "fmt"

var ARROW_DIRECTIONS map[string]Coord = map[string]Coord{
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
	"^": {0, -1},
}

func PrettyPrintGrid(grid *[][]string) {
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

func CopyGrid[T1 any](grid *[][]T1) *[][]T1 {
	newGrid := make([][]T1, len(*grid))
	for i, inner := range *grid {
		newInner := make([]T1, len(inner))
		copy(newInner, inner) // Use the built-in copy function
		newGrid[i] = newInner
	}
	return &newGrid
}
