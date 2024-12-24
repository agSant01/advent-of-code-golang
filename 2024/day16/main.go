package main

import (
	"agSant01/aoc-golang/utils"
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"

	maze := [][]string{}
	start := utils.Coord{X: 0, Y: 0}

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
		line := scanner.Text()
		maze = append(maze, strings.Split(line, ""))
		if index := strings.Index(line, "S"); index >= 0 {
			start = utils.Coord{X: index, Y: id}
		}
		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(&maze, start))
	// fmt.Println("Result exercise 2: ", exercise2(&maze, start))
}

type Tuple struct {
	X int
	Y int
}

type Visit struct {
	Coord      utils.Coord
	Steps      int
	Turns      int
	Cost       int
	CurrentDir string
	Path       []utils.Coord
}

func isTurn(dir1 string, dir2 string) bool {
	// > & >
	// < & <
	// ^ & ^
	if dir1 == "" || dir2 == "" {
		return true
	}

	if dir1 == dir2 {
		return false
	}

	if dir1 == "^" && dir2 == "v" {
		return false
	}
	if dir1 == "v" && dir2 == "^" {
		return false
	}
	if dir1 == "<" && dir2 == ">" {
		return false
	}
	if dir1 == ">" && dir2 == "<" {
		return false
	}

	return true
}

var ARROW_DIRECTIONS map[string]utils.Coord = map[string]utils.Coord{
	">": {X: 1, Y: 0},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
	"^": {X: 0, Y: -1},
}

// A PriorityQueue implements heap.Interface for items
type PriorityQueue []*Visit

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority comes first
	return pq[i].Cost < pq[j].Cost
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Visit))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func findShortestPath(maze *[][]string, start utils.Coord) int {

	toVisit := &PriorityQueue{
		&Visit{Coord: start, Steps: 0, Turns: 0, Cost: 0, CurrentDir: ""},
	}

	heap.Init(toVisit)

	minCost := 200_000
	seen := map[utils.Coord]string{}
	pths := 0
	bestPaths := map[int][]Visit{}
	for toVisit.Len() > 0 {
		pths++

		current := heap.Pop(toVisit).(*Visit)

		if v, e := seen[current.Coord]; e && v == current.CurrentDir {
			continue
		}
		seen[current.Coord] = current.CurrentDir

		if current.Cost >= minCost {
			continue
		}

		if val, _ := utils.GetCoordValue(maze, current.Coord); val == "E" {
			minCost = int(math.Min(float64(minCost), float64(current.Cost)))
			// fmt.Println("end", current)
			if _, e := bestPaths[current.Cost]; !e {
				bestPaths[current.Cost] = []Visit{}
			}
			bestPaths[current.Cost] = append(bestPaths[current.Cost], *current)
			// continue
		}

		for arrow, direction := range ARROW_DIRECTIONS {
			// for _, arrow := range []string{"v", "<", ">", "^"} {
			// 	direction := ARROW_DIRECTIONS[arrow]
			if value, isBounds := utils.GetCoordValue(maze,
				utils.Coord{X: current.Coord.X + direction.X,
					Y: current.Coord.Y + direction.Y}); !isBounds || value == "#" {
				continue
			}

			// Create a new slice with the same length as the original
			copySlice := make([]utils.Coord, len(current.Path))

			// Copy the elements from the original slice
			copy(copySlice, current.Path)

			v := Visit{
				Coord: utils.Coord{
					X: current.Coord.X + direction.X,
					Y: current.Coord.Y + direction.Y,
				},
				Steps:      current.Steps + 1,
				Turns:      current.Turns + utils.Bool2Int(isTurn(current.CurrentDir, arrow)),
				CurrentDir: arrow,
				Path: append(copySlice, utils.Coord{
					X: current.Coord.X + direction.X,
					Y: current.Coord.Y + direction.Y,
				}),
			}
			v.Cost = v.Turns*1000 + v.Steps
			heap.Push(toVisit, &v)
		}
	}

	println("shortest", len(bestPaths[minCost]))
	paths := bestPaths[minCost]
	bestSeats := map[utils.Coord]bool{}
	for _, p := range paths {
		for _, c := range p.Path {
			bestSeats[c] = true
		}
	}
	println(len(bestSeats))

	return minCost
}

func exercise1(maze *[][]string, start utils.Coord) int {
	// fmt.Println(start)
	// utils.PrettyPrintGrid(&maze)
	return findShortestPath(maze, start)
}

func exercise2(maze *[][]string, start utils.Coord) int {
	fmt.Println(start)
	// utils.PrettyPrintGrid(&maze)
	min := math.MaxFloat64
	for i := 0; i < 2; i++ {
		min = math.Min(float64(findShortestPath(maze, start)), min)
	}

	return int(min)
}
