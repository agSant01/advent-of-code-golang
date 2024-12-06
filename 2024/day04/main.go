package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func makeData(line string) []string {
	return strings.Split(line, "")
}

func main() {
	println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	idx := 0
	var data [][]string
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		data = append(data, makeData(scanner.Text()))
		idx++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	// println("Result exercise 1: ", exercise1(data))
	println("Result exercise 2: ", exercise2(data))
}

type Node struct {
	X      int
	Y      int
	Value  []string
	DeltaX int
	DeltaY int
}

type Tuple struct {
	A int
	B int
}

var nextDict = map[string]string{
	"X": "M",
	"M": "A",
	"A": "S",
}

func checkNeighbor(x int, y int, data [][]string) int {
	to_visit := []Node{}
	total := 0

	for i := -1; i <= 1; i += 1 {
		for j := -1; j <= 1; j += 1 {
			if i == 0 && j == 0 {
				continue
			}
			to_visit = append(to_visit, Node{x, y, []string{data[y][x]}, i, j})
		}
	}
	fmt.Println(to_visit)

	for len(to_visit) > 0 {
		// Pop the first element
		current := to_visit[0]
		to_visit = to_visit[1:] // Remove the first element
		fmt.Println(current)

		if len(current.Value) >= 4 {
			fmt.Println("found full", current)
			total += 1
			continue
		}

		lastChar := current.Value[len(current.Value)-1]
		expectedNext := nextDict[lastChar]

		nextX := current.X + current.DeltaX
		nextY := current.Y + current.DeltaY

		if nextX < 0 || nextX >= len(data[0]) {
			continue
		}
		if nextY < 0 || nextY >= len(data) {
			continue
		}

		fmt.Println("neigh i", current.X, "j", current.Y, "exp", lastChar, expectedNext, data[nextY][nextX])
		if expectedNext == data[nextY][nextX] {
			current.Value = append(current.Value, data[nextY][nextX])
			current.X = nextX
			current.Y = nextY
			to_visit = append(to_visit, current)
		}
	}

	return total
}

func exercise1(data [][]string) int {
	total := 0
	// visited := make(map[Tuple]bool)

	for j, line := range data {
		for i, char := range line {
			if char == "X" {
				fmt.Println("analyzing", i, j)
				total += checkNeighbor(i, j, data)
			}
		}
	}

	return total
}

var referenceFrame = [][]string{
	{"M", ".", "M"},
	{".", "A", "."},
	{"S", ".", "S"},
}

func rotateFrame(frame [][]string) [][]string {
	n := len(frame)
	newFrame := make([][]string, n)
	for i := range newFrame {
		newFrame[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newFrame[j][n-i-1] = frame[i][j]
		}
	}
	return newFrame
}

func prettyPrint(matrix [][]string) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%s", value) // Adjust the width as needed (4 spaces here)
		}
		fmt.Println()
	}
	fmt.Println()

}

func checkXmasCross(x int, y int, data [][]string) int {
	f := rotateFrame(referenceFrame)
	for rotation := 0; rotation < 4; rotation++ {
		prettyPrint(f)
		if data[y-1][x-1] == f[0][0] && data[y+1][x+1] == f[2][2] && data[y-1][x+1] == f[0][2] && data[y+1][x-1] == f[2][0] {
			fmt.Println("found match")
			return 1
		}
		f = rotateFrame(f)
	}
	return 0
}

func exercise2(data [][]string) int {
	result := 0
	for j := 1; j < len(data)-1; j++ {
		for i := 1; i < len(data[0])-1; i++ {
			char := data[j][i]
			if char == "A" {
				result += checkXmasCross(i, j, data)
			}
		}
	}
	return result
}
