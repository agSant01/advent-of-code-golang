package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Tuple struct {
	X int
	Y int
}

type Machine struct {
	ButtonA Tuple
	ButtonB Tuple
	Prize   Tuple
}

// func (m *Machine) pushA() {}

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"
	data := []Machine{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.Sca)

	r, err := regexp.Compile(`\d+`)
	if err != nil {
		// Handle compilation error
		os.Exit(2)
	}

	id := 0
	// data.PushFront("1") // dummy
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		buttonA := scanner.Text()
		valuesA := r.FindAllString(buttonA, -1)
		aX, _ := strconv.Atoi(valuesA[0])
		aY, _ := strconv.Atoi(valuesA[1])
		scanner.Scan()

		buttonB := scanner.Text()
		bValues := r.FindAllString(buttonB, -1)
		bX, _ := strconv.Atoi(bValues[0])
		bY, _ := strconv.Atoi(bValues[1])
		scanner.Scan()

		prize := scanner.Text()
		pValues := r.FindAllString(prize, -1)
		pX, _ := strconv.Atoi(pValues[0])
		pY, _ := strconv.Atoi(pValues[1])
		scanner.Scan()

		machine := Machine{
			ButtonA: Tuple{aX, aY},
			ButtonB: Tuple{bX, bY},
			Prize:   Tuple{pX, pY},
		}

		data = append(data, machine)

		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(data))
	fmt.Println("Result exercise 2: ", exercise2(data))
}

type Sim struct {
	X      int
	Y      int
	Cost   int
	aPress int
	bPress int
}

// Pop function to remove and return the last element
func pop(slice *[]Sim) Sim {
	if len(*slice) == 0 {
		panic("cannot pop from an empty slice")
	}

	// Retrieve the last element
	last := (*slice)[len(*slice)-1]

	// Reduce the length of the slice
	*slice = (*slice)[:len(*slice)-1]

	return last
}

// RemoveFront removes the first element from a slice and returns the removed element and the updated slice
func RemoveFront(slice *[]Sim) Sim {
	if len(*slice) == 0 {
		// Handle empty slice case
		return Sim{} // Return default zero value for int and the empty slice
	}
	removed := (*slice)[0]

	*slice = (*slice)[1:]

	return removed
}

func runSim(machine Machine) int {
	aCost := 3
	bCost := 1
	start := Sim{0, 0, 0, 0, 0}
	toVisit := []Sim{start}
	seen := map[Tuple]bool{}
	for len(toVisit) > 0 {
		current := pop(&toVisit)

		if _, exists := seen[Tuple{current.aPress, current.bPress}]; exists {
			continue
		}
		seen[Tuple{current.aPress, current.bPress}] = true

		if current.X > machine.Prize.X || current.Y > machine.Prize.Y {
			continue
		}

		if current.X == machine.Prize.X && current.Y == machine.Prize.Y {
			return current.Cost
		}

		if current.X+machine.ButtonB.X <= machine.Prize.X &&
			current.Y+machine.ButtonB.Y <= machine.Prize.Y {
			toVisit = append(toVisit, Sim{
				aPress: current.aPress,
				bPress: current.bPress + 1,
				Cost:   current.Cost + bCost,
				X:      current.X + machine.ButtonB.X,
				Y:      current.Y + machine.ButtonB.Y,
			})
		}

		if current.X+machine.ButtonA.X <= machine.Prize.X &&
			current.Y+machine.ButtonA.Y <= machine.Prize.Y {
			toVisit = append(toVisit, Sim{
				aPress: current.aPress + 1,
				bPress: current.bPress,
				Cost:   current.Cost + aCost,
				X:      current.X + machine.ButtonA.X,
				Y:      current.Y + machine.ButtonA.Y,
			})
		}
	}

	return 0
}

func exercise1(data []Machine) int {
	result := 0
	for _, machine := range data {
		result += runSim(machine)

	}

	return result
}

func calculate(machine Machine) int {
	aPresses := float64((machine.Prize.X*machine.ButtonB.Y)-(machine.Prize.Y*machine.ButtonB.X)) / float64(machine.ButtonA.X*machine.ButtonB.Y-machine.ButtonA.Y*machine.ButtonB.X)
	bPresses := float64(float64(machine.Prize.X)-float64(machine.ButtonA.X)*aPresses) / float64(machine.ButtonB.X)
	if math.Mod(aPresses, 1) != 0 || math.Mod(bPresses, 1) != 0 {
		return 0
	}
	return int(aPresses*3 + bPresses)
}

func exercise2(data []Machine) int {
	result := 0
	for _, machine := range data {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
		result += calculate(machine)
	}

	return int(result)
}
