package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"
	data := list.New()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	id := 0
	// data.PushFront("1") // dummy
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		// println(text)
		data.PushBack(text)
		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
	fmt.Println("Result exercise 1: ", exercise1(data))
	fmt.Println("Result exercise 2: ", exercise2(data))
}

func simulate(stones *list.List, iterations int) int {

	stonesSlice := list.New()
	for e := stones.Front(); e != nil; e = e.Next() {
		content := fmt.Sprintf("%s", e.Value)
		stonesSlice.PushBack(content)
	}

	for i := 0; i < iterations; i++ {
		for e := stonesSlice.Front(); e != nil; e = e.Next() {
			content := fmt.Sprintf("%s", e.Value)
			if len(content)%2 == 0 {
				left := content[0 : len(content)/2]
				left = strings.TrimLeft(left, "0")
				if left == "" {
					left = "0"
				}
				right := content[len(content)/2:]
				right = strings.TrimLeft(right, "0")
				if right == "" {
					right = "0"
				}
				e.Value = left
				e = stonesSlice.InsertAfter(right, e)
			} else if content == "0" {
				e.Value = "1"
			} else {
				value, _ := strconv.Atoi(content)
				e.Value = strconv.Itoa(value * 2024)
			}
		}
	}
	return stonesSlice.Len()
}

func exercise1(stones *list.List) int {
	// Iterate over the list
	result := simulate(stones, 25)
	return result
}

type Tuple struct {
	V    string
	Iter int
}

type ToCalculate struct {
	V    string
	Iter int
	Len  int
}

var seenPatterns map[Tuple]int = map[Tuple]int{}

func simulateDynamicRec(stone string, iterations int) int {
	result := 0

	// fmt.Println(current)
	// fmt.Println(seenPatterns)
	if _, exists := seenPatterns[Tuple{stone, iterations}]; exists {
		result += seenPatterns[Tuple{stone, iterations}]
	} else if iterations == 0 {
		return 1
	} else if stone == "0" {
		return simulateDynamicRec("1", iterations-1)
	} else if len(stone)%2 == 0 {
		left := stone[0 : len(stone)/2]
		left = strings.TrimLeft(left, "0")
		if left == "" {
			left = "0"
		}
		right := stone[len(stone)/2:]
		right = strings.TrimLeft(right, "0")
		if right == "" {
			right = "0"
		}
		result = simulateDynamicRec(left, iterations-1) + simulateDynamicRec(right, iterations-1)
	} else {
		value, _ := strconv.Atoi(stone)
		strValue := strconv.Itoa(value * 2024)
		result = simulateDynamicRec(strValue, iterations-1)
	}

	seenPatterns[Tuple{stone, iterations}] = result
	return result
}

func exercise2(stones *list.List) int {

	stonesSlice := []string{}
	for e := stones.Front(); e != nil; e = e.Next() {
		content := fmt.Sprintf("%s", e.Value)
		stonesSlice = append(stonesSlice, content)
	}

	result := 0
	for _, stone := range stonesSlice {
		result += simulateDynamicRec(stone, 75)
	}

	return result
}
