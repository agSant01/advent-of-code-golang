package main

import (
	"agSant01/aoc-golang/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	// filename := "./data_test.txt"
	filename := "./data_test2.txt"
	filename = "./data.txt"

	data := [][]string{}
	directions := ""
	start := utils.Coord{X: 0, Y: 0}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 0
	isMap := true
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		line := scanner.Text()

		if len(line) == 0 {
			isMap = false
			continue
		}

		if isMap {
			data = append(data, strings.Split(line, ""))
			if index := strings.Index(line, "@"); index >= 0 {
				start = utils.Coord{X: index, Y: id}
			}
		} else {
			directions += line
		}
		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	warehouseCopy := *utils.CopyGrid(&data)

	fmt.Println("Result exercise 1: ", exercise1(data, strings.Split(directions, ""), start, false))
	fmt.Println("Result exercise 2: ", exercise2(warehouseCopy, strings.Split(directions, "")))
}

func getGPSDistanceSum(warehouse *[][]string) int {
	result := 0
	// 100 times its distance from the top edge of the map
	// plus its distance from the left edge of the map.
	// (This process does not stop at wall tiles; measure all the way to the edges of the map.)
	for y, line := range *warehouse {
		for x, v := range line {
			// O for part1;
			// ] for part2;
			if v == "O" || v == "[" {
				result += 100*y + x
			}
		}
	}
	return result
}

func exercise1(warehouse [][]string, directions []string, start utils.Coord, printGrid bool) int {
	curr := start.Copy()
	iter := 0
	for len(directions) > 0 {
		c := utils.RemoveFront(&directions)
		if c == nil {
			break
		}

		t := utils.ARROW_DIRECTIONS[*c]
		tmp := curr
		moves := true

		coords := []utils.Coord{}
		for {
			value, isBounds := utils.GetCoordValue(&warehouse, tmp)

			if !isBounds {
				break
			} else if value == "#" {
				moves = false
				break
			} else if value == "." {
				break
			}
			coords = append(coords, tmp)

			tmp.Displaces(t)
		}

		for moves && len(coords) > 0 {
			n := utils.Pop(&coords)
			b := n.Add(t)
			// fmt.Println("swapping ", n, b)
			utils.SwapGridValues(&warehouse, n, b)
			curr = b
		}

		if printGrid {
			fmt.Printf("Iter %d:\n", iter)
			utils.PrettyPrintGrid(&warehouse)
			iter++
		}
	}

	return getGPSDistanceSum(&warehouse)
}

func getExpandedWarehouse(warehouse [][]string) ([][]string, utils.Coord) {
	expandedWarehouse := [][]string{}

	for _, line := range warehouse {
		newLine := ""
		for _, val := range line {
			if val == "#" {
				newLine += "##"
			} else if val == "@" {
				newLine += "@."
			} else if val == "O" {
				newLine += "[]"
			} else if val == "." {
				newLine += ".."
			}
		}
		expandedWarehouse = append(expandedWarehouse, strings.Split(newLine, ""))
	}

	newStart := utils.Coord{}

	for y, line := range expandedWarehouse {
		for x, v := range line {
			if v == "@" {
				newStart = utils.Coord{X: x, Y: y}
				break
			}
		}
	}

	return expandedWarehouse, newStart
}

func getBlocksToMove(warehouse *[][]string, direction string, rippleStart utils.Coord, blocks *[]utils.Coord) bool {
	// fmt.Println("analyzing", rippleStart)
	if slices.Contains(*blocks, rippleStart) {
		return true
	}

	dv := utils.ARROW_DIRECTIONS[direction]
	tmp := rippleStart.Add(dv)
	value, isBounds := utils.GetCoordValue(warehouse, tmp)

	if !isBounds {
		return false
	}
	if value == "#" {
		return false
	}

	*blocks = append(*blocks, rippleStart)

	if value == "." {
		return true
	}

	otherBracket := utils.Coord{}
	if value == "[" {
		otherBracket = utils.Coord{X: tmp.X + 1, Y: tmp.Y}
	}
	if value == "]" {
		otherBracket = utils.Coord{X: tmp.X - 1, Y: tmp.Y}
	}

	return getBlocksToMove(warehouse, direction, tmp, blocks) && getBlocksToMove(warehouse, direction, otherBracket, blocks)
}

func exercise2(warehouse [][]string, directions []string) int {
	expandedWarehouse, newStart := getExpandedWarehouse(warehouse)
	utils.PrettyPrintGrid(&expandedWarehouse)

	curr := newStart.Copy()
	// iter := 0
	oldValues := map[utils.Coord]string{}
	for len(directions) > 0 {
		directionString := utils.RemoveFront(&directions)
		if directionString == nil {
			break
		}
		// fmt.Printf("Iter %d:\n", iter)
		// fmt.Printf("Direction %s:\n", *directionString)

		coords := []utils.Coord{}
		moves := getBlocksToMove(&expandedWarehouse, *directionString, curr, &coords)

		if moves {
			t := utils.ARROW_DIRECTIONS[*directionString]
			for _, c := range coords {
				oldValues[c] = expandedWarehouse[c.Y][c.X]
				expandedWarehouse[c.Y][c.X] = "."
			}
			for _, c := range coords {
				expandedWarehouse[c.Y+t.Y][c.X+t.X] = oldValues[c]
			}
			curr.Displaces(t)
		}
		// utils.PrettyPrintGrid(&expandedWarehouse)
	}
	// utils.PrettyPrintGrid(&expandedWarehouse)

	return getGPSDistanceSum(&expandedWarehouse)
}
