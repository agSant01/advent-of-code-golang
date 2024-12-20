package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"
	data := []Robot{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, err := regexp.Compile(`p=(-*\d+),(-*\d+) v=(-*\d+),(-*\d+)`)
	if err != nil {
		os.Exit(2) // Handle compilation error
	}

	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		line := scanner.Text()
		values := r.FindAllStringSubmatch(line, -1)
		pX, _ := strconv.Atoi(values[0][1])
		pY, _ := strconv.Atoi(values[0][2])
		vX, _ := strconv.Atoi(values[0][3])
		vY, _ := strconv.Atoi(values[0][4])

		robot := Robot{
			Id:       id,
			Position: Tuple{pX, pY},
			Velocity: Tuple{vX, vY},
		}

		data = append(data, robot)

		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data, len(data))
	// fmt.Println("Result exercise 1: ", exercise1(data))
	fmt.Println("Result exercise 2: ", exercise2(data))
}

type Tuple struct {
	X int
	Y int
}
type Robot struct {
	Id       int
	Position Tuple
	Velocity Tuple
}

func (r *Robot) move(width int, height int) {
	r.Position.X = PositiveModulo((r.Position.X + r.Velocity.X), width)
	r.Position.Y = PositiveModulo((r.Position.Y + r.Velocity.Y), height)
}

func runSimulation(robots *[]Robot, iter int, width int, height int) {
	for i := 0; i < iter; i++ {
		for i := range *robots {
			fmt.Println("m1", (*robots)[i])
			(*robots)[i].move(width, height)
			fmt.Println("m2", (*robots)[i])
		}
	}
}

func exercise1(data []Robot) int {
	// which is 101 tiles wide and 103 tiles tall
	width := 101
	height := 103

	horizontalLine := height / 2
	verticalLine := width / 2
	fmt.Println(data)

	runSimulation(&data, 100, width, height)

	occupied := map[Tuple]int{}
	for _, r := range data {
		occupied[r.Position] += 1
	}
	fmt.Println(occupied, verticalLine, horizontalLine)

	qVals := []int{0, 0, 0, 0}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x == verticalLine || y == horizontalLine {
				continue
			}
			q := 0
			if x < verticalLine && y < horizontalLine {
				q = 0
			} else if x < verticalLine && y > horizontalLine {
				q = 1
			} else if x > verticalLine && y < horizontalLine {
				q = 2
			} else if x > verticalLine && y > horizontalLine {
				q = 3
			} else {
				os.Exit(4)
			}

			if v, exists := occupied[Tuple{x, y}]; exists {
				qVals[q] += v
			}
		}
	}
	fmt.Println(qVals)
	res := 1
	for _, v := range qVals {
		res *= v
	}

	prettyPrint(occupied, width, height)
	return res
}

func prettyPrint(occupied map[Tuple]int, width int, height int) {
	// horizontalLine := height / 2
	// verticalLine := width / 2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// if x == verticalLine || y == horizontalLine {
			// 	fmt.Print(" ")
			// 	continue
			// }
			if v, exists := occupied[Tuple{x, y}]; exists {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		println()
	}
}

// PositiveModulo ensures the result of the modulo is always positive
func PositiveModulo(a, b int) int {
	result := a % b
	if result < 0 {
		result += b
	}
	return result
}

func exercise2(data []Robot) int {
	// which is 101 tiles wide and 103 tiles tall
	width := 101
	height := 103

	k := 0
	for {
		for i := 0; i < 1; i++ {
			for i := range data {
				data[i].move(width, height)
			}

		}

		occupied := map[Tuple]int{}
		max := 0
		for _, r := range data {
			occupied[r.Position] += 1
			max = int(math.Max(float64(max), float64(occupied[r.Position])))
		}
		k++
		if max == 1 {
			prettyPrint(occupied, width, height)
			time.Sleep(1 * time.Second)
			return k
		}

	}

	return k
}
