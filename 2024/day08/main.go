package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Antenna struct {
	X         int
	Y         int
	Frequency string
}

func (a1 Antenna) diff(a2 Antenna) Antenna {
	return Antenna{a1.X - a2.X, a1.Y - a2.Y, a1.Frequency}
}

func (a1 Antenna) Add(a2 Antenna) Antenna {
	return Antenna{
		X:         a1.X + a2.X,
		Y:         a1.Y + a2.Y,
		Frequency: a2.Frequency,
	}
}

func parseLine(line string) []string {
	return strings.Split(line, "")
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
	fieldMap := [][]string{}
	antennasMap := map[string][]Antenna{}
	y := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()

		line := parseLine(text)

		for x, val := range line {
			if val != "." {
				if _, exists := antennasMap[val]; !exists {
					antennasMap[val] = []Antenna{}
				}
				antennasMap[val] = append(antennasMap[val], Antenna{x, y, val})
			}
		}

		fieldMap = append(fieldMap, line)
		y++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(antennasMap)

	fmt.Println("Result exercise 1: ", exercise1(fieldMap, antennasMap))
	fmt.Println("Result exercise 2: ", exercise2(fieldMap, antennasMap))
}

func withinBounds(antiNode Antenna, width int, height int) bool {
	return antiNode.X >= 0 && antiNode.Y >= 0 && antiNode.X < width && antiNode.Y < height
}

func printField(fieldMap [][]string, antiNodes map[Antenna]bool) {
	for y, line := range fieldMap {
		for x, val := range line {
			if val != "." {
				fmt.Print(val)
			} else if _, exists := antiNodes[Antenna{x, y, "#"}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func exercise1(fieldMap [][]string, antennaMap map[string][]Antenna) int {
	height := len(fieldMap)
	width := len(fieldMap[0])

	antiNodeMap := map[Antenna]bool{}
	for freq, antennas := range antennaMap {
		fc := 0
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}

				diffVec := a2.diff(a1)
				diffVec.Frequency = "#"
				diffVec = diffVec.Add(diffVec)
				absAntiNode := a1.Add(diffVec)

				if withinBounds(absAntiNode, width, height) {
					antiNodeMap[absAntiNode] = true
				}
			}
		}
		fmt.Println(freq, len(antennas), "antiNodes=", fc)
	}

	printField(fieldMap, antiNodeMap)

	return len(antiNodeMap)
}

func exercise2(fieldMap [][]string, antennaMap map[string][]Antenna) int {
	height := len(fieldMap)
	width := len(fieldMap[0])

	antiNodeMap := map[Antenna]bool{}
	for _, antennas := range antennaMap {
		for _, a1 := range antennas {
			for _, a2 := range antennas {
				if a1 == a2 {
					continue
				}

				diffVec := a2.diff(a1)
				diffVec.Frequency = "#"
				absAntiNode := a1.Add(diffVec)
				for {
					if withinBounds(absAntiNode, width, height) {
						antiNodeMap[absAntiNode] = true
					} else {
						break
					}
					absAntiNode = absAntiNode.Add(diffVec)
				}

			}
		}
	}

	// printField(fieldMap, antiNodeMap)

	return len(antiNodeMap)
}
