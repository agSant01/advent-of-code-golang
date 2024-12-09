package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func parseLine(line string) []string {
	array := strings.Split(line, "")
	return array
}

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	labMap := [][]string{}
	startCoord := Coord{}
	y := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		if idx := strings.Index(text, "^"); idx >= 0 {
			startCoord = Coord{X: idx, Y: y}
		}

		labMap = append(labMap, parseLine(text))
		y++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(labMap, startCoord))
	fmt.Println("Result exercise 2: ", exercise2(labMap, startCoord))
}

var DIRECTION_MAPS = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

var NEXT_STEP = map[string]Coord{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}

func getMapValue(labMap [][]string, current Coord) string {
	return labMap[current.Y][current.X]
}

func exercise1(labMap [][]string, startCoord Coord) int {
	total := 0

	currentCoord := startCoord
	currentStep := NEXT_STEP[getMapValue(labMap, currentCoord)]
	currentDirection := getMapValue(labMap, currentCoord)

	width := len(labMap[0])
	height := len(labMap)
	visited := map[Coord]string{}
	visited[startCoord] = currentDirection
	for {
		nextCoord := Coord{
			X: currentCoord.X + currentStep.X,
			Y: currentCoord.Y + currentStep.Y,
		}
		visited[currentCoord] = currentDirection

		if nextCoord.X < 0 || nextCoord.Y < 0 || nextCoord.X >= width || nextCoord.Y >= height {
			break
		}
		value := getMapValue(labMap, nextCoord)

		if value == "#" {
			currentDirection = DIRECTION_MAPS[currentDirection]
			currentStep = NEXT_STEP[currentDirection]
			nextCoord = Coord{
				X: currentCoord.X + currentStep.X,
				Y: currentCoord.Y + currentStep.Y,
			}
		}

		currentCoord = nextCoord
		total += 1
	}

	return len(visited)
}

type SeenCoord struct {
	X         int
	Y         int
	Direction string
}

func isLoop(labMap [][]string, startCoord Coord, currentDirection string) bool {
	height := len(labMap)
	width := len(labMap[0])

	currentCoord := Coord{startCoord.X, startCoord.Y}
	currentStep := NEXT_STEP["^"]

	seen := map[SeenCoord]bool{}
	seen[SeenCoord{currentCoord.X, currentCoord.Y, currentDirection}] = true

	for {
		nextCoord := Coord{
			X: currentCoord.X + currentStep.X,
			Y: currentCoord.Y + currentStep.Y,
		}

		if nextCoord.X >= width || nextCoord.Y >= height || nextCoord.X < 0 || nextCoord.Y < 0 {
			return false
		}

		if getMapValue(labMap, nextCoord) == "#" {
			currentDirection = DIRECTION_MAPS[currentDirection]
			currentStep = NEXT_STEP[currentDirection]
		} else {
			currentCoord = Coord{
				X: currentCoord.X + currentStep.X,
				Y: currentCoord.Y + currentStep.Y,
			}

			_, exist := seen[SeenCoord{currentCoord.X, currentCoord.Y, currentDirection}]
			if exist {
				return true
			}

			seen[SeenCoord{currentCoord.X, currentCoord.Y, currentDirection}] = true
		}
	}
}

func exercise2(labMap [][]string, startCoord Coord) int {
	height := len(labMap)
	width := len(labMap[0])
	fmt.Println("Start=", startCoord)
	p2 := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			original := labMap[y][x]
			if original == "#" {
				continue
			}
			labMap[y][x] = "#"
			if isLoop(labMap, startCoord, "^") {
				p2 += 1
			}
			labMap[y][x] = original
		}
	}
	return p2
}
