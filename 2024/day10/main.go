package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	X int
	Y int
}

type ToVisit struct {
	ToVisit Node
	Visited map[Node]bool
	Path    []Node
	// Visited int
}

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"
	data := [][]string{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		println(text)
		data = append(data, strings.Split(text, ""))
		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(data))
	fmt.Println("Result exercise 2: ", exercise2(data))
}

// Pop function to remove and return the last element
func pop(slice *[]ToVisit) ToVisit {
	if len(*slice) == 0 {
		panic("cannot pop from an empty slice")
	}

	// Retrieve the last element
	last := (*slice)[len(*slice)-1]

	// Reduce the length of the slice
	*slice = (*slice)[:len(*slice)-1]

	return last
}

func getNeighbors(node Node, width int, height int) []Node {
	toReturn := []Node{}
	for x := node.X - 1; x <= node.X+1; x++ {
		if node.X == x {
			continue
		}
		if x < 0 || x >= width {
			continue
		}
		toReturn = append(toReturn, Node{x, node.Y})
	}

	for y := node.Y - 1; y <= node.Y+1; y++ {
		if node.Y == y {
			continue
		}
		if y < 0 || y >= height {
			continue
		}
		toReturn = append(toReturn, Node{node.X, y})
	}

	return toReturn
}

func getValue(node Node, trails *[][]string) (string, bool) {
	width := len((*trails)[0])
	height := len(*trails)

	if node.X < 0 || node.Y < 0 || node.X >= width || node.Y >= height {
		return "", false
	}
	// println(node.X, node.Y)
	// println("aa", (*trails)[node.Y][node.X])
	return (*trails)[node.Y][node.X], true
}

func getPaths(trails [][]string, x int, y int, countUniqueTrail bool) int {
	start := ToVisit{ToVisit: Node{X: x, Y: y}, Visited: map[Node]bool{}, Path: []Node{}}
	toVisit := []ToVisit{start}
	width := len(trails[0])
	height := len(trails)

	// fmt.Println(start)

	visited := map[Node]bool{}
	paths := 0
	for len(toVisit) > 0 {
		current := pop(&toVisit)
		// fmt.Println("curr ", current)

		// visited := current.Visited
		if _, exists := visited[current.ToVisit]; exists {
			// fmt.Println("huh", current)
			continue
		}

		if !countUniqueTrail {
			visited[current.ToVisit] = true
		}

		if value, isBounds := getValue(current.ToVisit, &trails); isBounds && value == "9" {
			paths++
			// fmt.Println("arrived to 9", current)
			continue
		}

		for _, neigh := range getNeighbors(current.ToVisit, width, height) {
			nei, _ := getValue(neigh, &trails)
			curr, _ := getValue(current.ToVisit, &trails)
			iNei, _ := strconv.Atoi(nei)
			iCurr, _ := strconv.Atoi(curr)
			// fmt.Println("neight lvl", start, neigh, iCurr, iNei)
			if iNei-iCurr != 1 {
				continue
			}

			toVisit = append(toVisit, ToVisit{
				ToVisit: neigh,
				Path:    append(current.Path, neigh),
				Visited: current.Visited,
			})
		}

	}

	return paths
}

func exercise1(data [][]string) int {
	count := 0
	for y, line := range data {
		for x, v := range line {
			if v != "0" {
				continue
			}
			count += getPaths(data, x, y, false)
		}
	}
	return count
}

func exercise2(data [][]string) int {
	count := 0
	for y, line := range data {
		for x, v := range line {
			if v != "0" {
				continue
			}
			count += getPaths(data, x, y, true)
		}
	}
	return count
}
