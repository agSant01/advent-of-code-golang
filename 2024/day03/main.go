package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	operation string
	valueA    int
	valueB    int
}

var r = regexp.MustCompile(`(mul)\((\d+),(\d+)\)|don\'t|do`)

func makeData(line string) []Instruction {
	matches := r.FindAllStringSubmatch(line, -1)
	result := []Instruction{}

	for _, m := range matches {
		intA, _ := strconv.Atoi(m[2])
		intB, _ := strconv.Atoi(m[3])
		if m[1] == "mul" {
			result = append(result, Instruction{
				operation: m[1],
				valueA:    intA,
				valueB:    intB,
			})
		} else {
			result = append(result, Instruction{
				operation: m[0],
				valueA:    0,
				valueB:    0,
			})
		}
	}

	return result
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
	var data []Instruction
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one instruction
		data = append(data, makeData(scanner.Text())...)
		idx++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	println("Result exercise 1: ", exercise1(data))
	println("Result exercise 2: ", exercise2(data))
}

func exercise1(data []Instruction) int {
	result := 0
	for _, inst := range data {
		if inst.operation == "mul" {
			result += (inst.valueA * inst.valueB)
		}
	}
	return result
}

func exercise2(data []Instruction) int {
	result := 0
	enable := true
	for _, inst := range data {
		if inst.operation == "mul" && enable {
			result += (inst.valueA * inst.valueB)
		} else if inst.operation == "do" {
			enable = true
		} else if inst.operation == "don't" {
			enable = false
		}
	}
	return result
}
