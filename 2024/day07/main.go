package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func parseLine(line string) any {
	values := strings.Split(line, ":")
	testValue, _ := strconv.Atoi(values[0])
	numbers := []int{}
	strNums := strings.Split(values[1], " ")
	for i := 1; i < len(strNums); i++ {
		intVal, _ := strconv.Atoi(strNums[i])
		numbers = append(numbers, intVal)
	}

	return Equation{testValue, numbers}
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
	equations := []Equation{}
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		value, _ := parseLine(text).(Equation)
		equations = append(equations, value)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result exercise 1: ", exercise1(equations))
	fmt.Println("Result exercise 2: ", exercise2(equations))
}

func calculate(currentValue int, testValue int, values []int, valueIdx int) bool {
	if currentValue > testValue {
		return false
	}

	if currentValue == testValue {
		return true
	}

	if valueIdx >= len(values) {
		return false
	}

	currentValueSum := currentValue + values[valueIdx]
	if currentValue == 0 {
		currentValue = 1
	}
	currentValueMult := currentValue * values[valueIdx]

	return calculate(currentValueSum, testValue, values, valueIdx+1) || calculate(currentValueMult, testValue, values, valueIdx+1)
}

func exercise1(equations []Equation) int {
	total := 0

	for _, equation := range equations {
		tv := equation.TestValue

		if calculate(0, tv, equation.Numbers, 0) {
			total += tv
		}

	}

	return total
}

func calculateWithConcat(currentValue int, testValue int, values []int, valueIdx int) bool {
	if currentValue > testValue {
		return false
	}

	if currentValue == testValue {
		return true
	}

	if valueIdx >= len(values) {
		return false
	}

	currentValueSum := currentValue + values[valueIdx]
	if currentValue == 0 {
		currentValue = 1
	}
	currentValueMult := currentValue * values[valueIdx]

	concatValue := strconv.Itoa(currentValue) + strconv.Itoa(values[valueIdx])
	concatValueInt, _ := strconv.Atoi(concatValue)

	return calculateWithConcat(currentValueSum, testValue, values, valueIdx+1) || calculateWithConcat(currentValueMult, testValue, values, valueIdx+1) || calculateWithConcat(concatValueInt, testValue, values, valueIdx+1)
}

func exercise2(equations []Equation) int {
	total := 0
	for _, equation := range equations {
		tv := equation.TestValue

		if calculateWithConcat(0, tv, equation.Numbers, 0) {
			total += tv
		}

	}
	return total
}
