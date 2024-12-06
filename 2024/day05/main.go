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
	X      int
	Y      int
	Value  []string
	DeltaX int
	DeltaY int
}

type Tuple struct {
	A int
	B int
}

func rules(line string) []string {
	return strings.Split(line, "")
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
	mode := "rules"
	rules := [100][]int{{}}
	pageNumberUpdates := [][]int{}
	idx := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		if len(text) == 0 {
			mode = "print_order"
			continue
		}
		if mode == "rules" {
			beforeAfter := strings.Split(text, "|")
			before, _ := strconv.Atoi(beforeAfter[0])
			after, _ := strconv.Atoi(beforeAfter[1])
			rules[before] = append(rules[before], after)
		} else {
			values := strings.Split(text, ",")
			intValues := make([]int, len(values))
			for i, value := range values {
				intValues[i], _ = strconv.Atoi(value)
			}
			pageNumberUpdates = append(pageNumberUpdates, intValues)
		}
		idx++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// prettyPrint(rules)
	// fmt.Println(pageNumberUpdates)

	println("Result exercise 1: ", exercise1(rules, pageNumberUpdates))
	println("Result exercise 2: ", exercise2(rules, pageNumberUpdates))
}

func contains(arr []int, target int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func prettyPrint(rules [100][]int) {
	for i, rule := range rules {
		fmt.Print("i:", i, "=", rule)
		println()
	}
	println()
}

func exercise1(rules [100][]int, pageNumberUpdates [][]int) int {
	total := 0
	for _, singleUpdate := range pageNumberUpdates {
		validLine := true
		for i := 1; i < len(singleUpdate); i++ {
			if len(rules[singleUpdate[i]]) > 0 {
				mustBeBefore := rules[singleUpdate[i]]
				for j := i; j >= 0; j-- {
					if contains(mustBeBefore, singleUpdate[j]) {
						validLine = false
						// fmt.Println("XX not valid", singleUpdate)
						break
					}
				}
				if !validLine {
					break
				}
			}
		}

		if validLine {
			total += singleUpdate[len(singleUpdate)/2]
			// fmt.Println("YY => valid", singleUpdate, singleUpdate[len(singleUpdate)/2])
		}
	}

	return total
}

func swap(pageUpdate *[]int, i int, j int) {
	tmp := (*pageUpdate)[i]
	(*pageUpdate)[i] = (*pageUpdate)[j]
	(*pageUpdate)[j] = tmp
}

func exercise2(rules [100][]int, pageNumberUpdates [][]int) int {
	result := 0
	incorrectly := map[int]bool{}
	for upd := 0; upd < len(pageNumberUpdates); {
		singleUpdate := pageNumberUpdates[upd]
		validLine := true
		for i := 0; i < len(singleUpdate); i++ {
			if len(rules[singleUpdate[i]]) > 0 {
				mustBeBefore := rules[singleUpdate[i]]
				for j := i; j >= 0; j-- {
					if contains(mustBeBefore, singleUpdate[j]) {
						validLine = false
						incorrectly[upd] = true
						// fmt.Println("XX not valid", singleUpdate)
						swap(&singleUpdate, i, j)
						// fmt.Println("XX new", singleUpdate)
						break
					}
				}
				if !validLine {
					break
				}
			}
		}
		if validLine {
			upd++
		}
	}

	// fmt.Println(incorrectly)

	for k := range incorrectly {
		singleUpdate := pageNumberUpdates[k]
		result += singleUpdate[len(singleUpdate)/2]
	}

	return result
}
