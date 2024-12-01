package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Left  int
	Right int
}

func makePair(line string) Pair {
	trimmedStr := strings.TrimSpace(line)
	values := strings.Split(trimmedStr, "   ")
	intA, _ := strconv.Atoi(values[0])
	intB, _ := strconv.Atoi(values[1])
	return Pair{
		Left:  intA,
		Right: intB,
	}
}

func main() {
	// filename := "./data_test.txt"
	filename := "./data.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	idx := 0
	var data []Pair

	for scanner.Scan() {
		data = append(data, makePair(scanner.Text()))
		idx++
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	res1 := exercise1(data)
	fmt.Println("Result for part1: ", res1)

	res2 := exercise2(data)
	fmt.Println("Result for part2: ", res2)

}

// What is the total distance between your lists?
// total distance
func exercise1(data []Pair) int {
	var list1 []int
	var list2 []int

	for _, v := range data {
		list1 = append(list1, v.Left)
		list2 = append(list2, v.Right)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	diff := 0.0
	for i := 0; i < len(data); i++ {
		local_diff := list1[i] - list2[i]
		diff += math.Abs(float64(local_diff))
	}

	return int(diff)
}

// What is their similarity score?
// Figure out exactly how often each number from the left list appears in the right list.
func exercise2(data []Pair) int {
	counter := make(map[int]int)

	// check Left on Right counts
	for _, pair := range data {
		_, exist := counter[pair.Right]
		if exist {
			counter[pair.Right] += 1
		} else {
			counter[pair.Right] = 1
		}
	}

	similarityScore := 0
	for _, v := range data {
		occurrences := counter[v.Left]
		similarityScore += (v.Left * occurrences)
	}

	return similarityScore
}
