package main

import (
	"agSant01/aoc-golang/utils"
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is main for 2024/day18")
	filename := "./data_test.in"
	filename = "./data.in"

	coordinates := []utils.Coord{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		line := scanner.Text()
		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		coordinates = append(coordinates, utils.Coord{X: x, Y: y})
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(coordinates)

	fmt.Println("Result exercise 1 with HEAP: ", exercise1(&coordinates))
	fmt.Println("Result exercise 2: ", exercise2(&coordinates))
}

func simulate(coords *[]utils.Coord, bts int, size int) *[][]int8 {
	memoryMap := make([][]int8, size)
	for k := range size {
		memoryMap[k] = make([]int8, size)
		for x := range size {
			memoryMap[k][x] = 0
		}
	}
	for sim := range bts {
		c := (*coords)[sim]
		memoryMap[c.Y][c.X] = 1
	}
	return &memoryMap
}

type ToVisit struct {
	Coord utils.Coord
	Steps int
}

// ToVisitHeap defines a min-heap of integers
type ToVisitHeap []ToVisit

// Len is the number of elements in the heap
func (h ToVisitHeap) Len() int { return len(h) }

// Less defines the heap property (min-heap)
func (h ToVisitHeap) Less(i, j int) bool { return h[i].Steps < h[j].Steps }

// Swap swaps elements at indexes i and j
func (h ToVisitHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *ToVisitHeap) Push(x any) {
	*h = append(*h, x.(ToVisit))
}

// Pop removes the smallest element from the heap
func (h *ToVisitHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findPath(memory *[][]int8, start utils.Coord, target utils.Coord, size int, cntIters bool) int {
	tv := ToVisit{
		Coord: start,
		Steps: 0,
	}
	pending := &ToVisitHeap{tv}
	heap.Init(pending)

	seen := map[utils.Coord]bool{}
	seen[start] = true

	iters := 0
	for pending.Len() > 0 {
		iters++
		curr := heap.Pop(pending).(ToVisit)
		if curr.Coord == target {
			if cntIters {
				fmt.Println(" - iters", iters)
			}
			return curr.Steps
		}
		for _, nb := range utils.GetNeighbors(curr.Coord, size, size) {
			value, _ := utils.GetCoordValueInt8(memory, nb)
			if value == 1 {
				continue
			}
			ntv := ToVisit{
				Coord: nb,
				Steps: curr.Steps + 1,
			}
			if _, e := seen[nb]; e {
				continue
			}
			seen[nb] = true
			heap.Push(pending, ntv)
		}
	}

	if cntIters {
		fmt.Println(" - iters", iters)
	}
	return -1
}

func exercise1(coordinates *[]utils.Coord) int {

	start := utils.Coord{X: 0, Y: 0}

	test := false
	size := 7
	target := utils.Coord{
		X: 6,
		Y: 6,
	}
	simBytes := 12
	if !test {
		simBytes = 1024
		size = 71
		target = utils.Coord{
			X: 70,
			Y: 70,
		}
	}

	memoryMap := simulate(coordinates, simBytes, size)
	utils.PrettyPrintGrid(memoryMap)
	return findPath(memoryMap, start, target, size, true)
}

func exercise2(coordinates *[]utils.Coord) string {
	start := utils.Coord{X: 0, Y: 0}

	test := false
	size := 7
	target := utils.Coord{
		X: 6,
		Y: 6,
	}
	if !test {
		size = 71
		target = utils.Coord{
			X: 70,
			Y: 70,
		}
	}

	// start with the end state of the memory
	i := len(*coordinates) - 1

	// simulate all bytes
	memoryMap := *simulate(coordinates, i, size)

	// go backwards, logic says that there is more than 50% chance the Blocker is closer to the end
	for i >= 0 {
		simBytes := (*coordinates)[i]

		// toggle off the byte
		memoryMap[simBytes.Y][simBytes.X] = 0
		i--

		r := findPath(&memoryMap, start, target, size, false)
		if r != -1 {
			return fmt.Sprintf("%+v %d", simBytes, i)
		}
	}

	return "<>"
}
