package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Values         []int
	OriginalValues []int
	Len            int
}

func makeReport(line string) Report {
	values := strings.Split(line, " ")
	intValues := make([]int, len(values))
	copiedSlice := make([]int, len(values))

	// Copy the elements using the built-in copy function
	copy(copiedSlice, intValues)

	for i, val := range values {
		intValues[i], _ = strconv.Atoi(val)
		copiedSlice[i], _ = strconv.Atoi(val)
	}

	return Report{
		Values:         intValues,
		OriginalValues: copiedSlice,
		Len:            len(intValues),
	}
}

func main() {
	println("This is main for 2024/day02")
	filename := "./data_test.txt"
	filename = "./data.txt"
	// filename = "./weird.txt"
	// filename = "./diff.out"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	idx := 0
	var data []Report

	for scanner.Scan() {
		data = append(data, makeReport(scanner.Text()))
		idx++
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	println("Result exercise 1: ", exercise1(data))
	println("Result exercise 2: ", exercise2(data))

}

func isSafeReport(report Report) bool {
	safe := true

	direction := "increasing"
	if report.Values[0] > report.Values[1] {
		direction = "decreasing"
	}

	idx := 0
	for idx < len(report.Values)-1 {
		if report.Values[idx] == report.Values[idx+1] {
			safe = false
			break
		}
		// if diff is more than 3: Unsafe
		diff := report.Values[idx] - report.Values[idx+1]
		if math.Abs(float64(diff)) > 3 {
			safe = false
			break
		}

		if direction == "decreasing" && diff < 0 {
			safe = false
			break
		} else if direction == "increasing" && diff > 0 {
			safe = false
			break
		}

		idx++
	}

	return safe
}

func remove(slice []int, index int) []int {
	newArr := make([]int, len(slice))
	copy(newArr, slice)
	return append(newArr[:index], newArr[index+1:]...)
}

func exercise1(data []Report) int {
	safeReports := 0
	for _, report := range data {
		safe := isSafeReport(report)
		if safe {
			safeReports += 1
		}
	}

	return safeReports
}

func exercise2(data []Report) int {
	safeReports := 0
	for _, report := range data {
		// if default report is not Safe, calculate the sub sets
		if !isSafeReport(report) {
			for i := 0; i < len(report.OriginalValues); i++ {
				report.Values = remove(report.OriginalValues, i)
				if isSafeReport(report) {
					safeReports += 1
					break
				}
			}
		} else {
			safeReports += 1
		}
	}

	return safeReports
}
