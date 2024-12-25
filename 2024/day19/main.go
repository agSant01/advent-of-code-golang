package main

import (
	"agSant01/aoc-golang/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is main for 2024/day19")

	filename := "./test.in"
	if len(os.Args) > 1 {
		filename = os.Args[1]
		fmt.Println("IsTest:", strings.Contains(filename, "test"))
	}
	isTest := strings.Contains(filename, "test")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	patterns := map[string]bool{}
	for _, v := range strings.Split(scanner.Text(), ", ") {
		patterns[v] = true
	}

	desiredPatterns := []string{}
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		desiredPatterns = append(desiredPatterns, line)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("patterns", len(patterns))
	fmt.Println("desired", len(desiredPatterns))

	fmt.Println("Result part 1:", part1(desiredPatterns, patterns, isTest))
	fmt.Println("Result part 2:", part2(desiredPatterns, patterns, isTest))
}

var cache = map[string]bool{}

func isPossible(desired string, patterns map[string]bool, longestWindow int) bool {
	if len(desired) == 0 {
		return true
	}
	if v, e := cache[desired]; e {
		return v
	}

	for i := 0; i <= longestWindow && i < len(desired)+1; i++ {
		_, e := patterns[desired[:i]]
		if e && isPossible(desired[i:], patterns, longestWindow) {
			cache[desired] = true
			return true
		}
	}

	cache[desired] = false
	return false
}

var cacheCount = map[string]int{}

func isPossibleCount(desired string, patterns map[string]bool, longestWindow int) int {
	if len(desired) == 0 {
		return 1
	}
	if v, e := cacheCount[desired]; e {
		return v
	}
	count := 0
	for i := 0; i <= longestWindow && i < len(desired)+1; i++ {
		_, e := patterns[desired[:i]]
		if e {
			tmp := isPossibleCount(desired[i:], patterns, longestWindow)
			count += tmp
		}
	}
	cacheCount[desired] = count
	return count
}

func part1(desiredPatterns []string, patterns map[string]bool, isTest bool) int {
	cnt := 0

	longestWindow := 0
	for k := range patterns {
		longestWindow = max(longestWindow, len(k))
	}

	for _, v := range desiredPatterns {
		res := isPossible(v, patterns, longestWindow)
		cnt += utils.Bool2Int(res)
	}
	return cnt
}

func part2(desiredPatterns []string, patterns map[string]bool, isTest bool) int {
	cnt := 0

	longestWindow := 0
	for k := range patterns {
		longestWindow = max(longestWindow, len(k))
	}

	for _, v := range desiredPatterns {
		cnt += isPossibleCount(v, patterns, longestWindow)
	}
	return cnt
}
