package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("This is main for 2024/day02")
	filename := "./data_test.txt"
	// data := [19]int{}
	filename = "./data.txt"
	data := [19999]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	id := 0
	for scanner.Scan() {
		// extend data slice
		// one line contains more than one Line
		text := scanner.Text()
		data[id], _ = strconv.Atoi(text)
		id++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result exercise 1: ", exercise1(data))
	fmt.Println("Result exercise 2: ", exercise2(&data))
}

func exercise1(data [19999]int) int {
	// func exercise1(data *[23]int) int {
	left := 1
	checkSum := 0
	blockId := 0
	mem := [500_000]int{}
	leftPtr := 0

	for i := len(data) - 1; i >= 1; {
		if (data)[i] <= 0 {
			i -= 2
			continue
		}

		for (data)[leftPtr] > 0 {
			// println("e blockId", blockId, "cS=", blockId, "*", (leftPtr / 2))
			checkSum += blockId * (leftPtr / 2)
			mem[blockId] = (leftPtr / 2)
			blockId++
			(data)[leftPtr]--
		}
		leftPtr += 2

		if left > i {
			// println("left vs i", left, i)
			break
		}

		for (data)[left] > 0 {
			for i >= 0 && (data)[i] == 0 {
				i -= 2
			}

			if i <= 0 {
				break
			}
			// add moved checksum
			checkSum += blockId * (i / 2)
			mem[blockId] = (i / 2)
			// println("o blockId", blockId, "cS=", blockId, "*", (i / 2))
			blockId++
			(data)[i]--
			(data)[left]--
		}
		left += 2

	}

	return checkSum
}

func getNextFreeBlock(disk *[]int16, start int) int {
	for i := 0; i < len(*disk); i++ {
		if start+i < len(*disk) && (*disk)[start+i] == -1 {
			return start + i
		}
	}
	return -1
}

func exercise2(data *[19999]int) int {
	checkSum := 0

	files := 0
	for i := 0; i < len(*data); i++ {
		files += (*data)[i]
	}
	disk := make([]int16, files)

	block := 0
	fileId := int16(0)
	for i := 0; i < len(*data); i++ {
		if i%2 == 0 {
			for j := 0; j < (*data)[i]; j++ {
				disk[block] = fileId
				block++
			}
			fileId++
		} else {
			// free
			for j := 0; j < (*data)[i]; j++ {
				disk[block] = -1
				block++
			}
		}
	}

	totalFiles := int16(len(*data) / 2)
	lastFile := len(disk) - 1
	for totalFiles > 0 {
		fileSize := 0
		for lastFile > 0 && disk[lastFile] == totalFiles {
			fileSize++
			lastFile--
		}
		freeSize := 0
		freeBlock := 0
		for freeBlock < len(disk) && freeSize < fileSize {
			freeSize = 0
			freeBlock = getNextFreeBlock(&disk, freeBlock+1)
			if freeBlock < 0 {
				break
			}
			for s := 0; freeBlock+s < len(disk) && disk[freeBlock+s] == -1; s++ {
				freeSize++
			}
		}

		if fileSize <= freeSize && freeBlock < lastFile {
			for fileSize > 0 {
				disk[freeBlock] = int16(totalFiles)
				disk[lastFile+fileSize] = -1
				freeBlock++
				fileSize--
			}
		}
		totalFiles--
		for lastFile > 0 && disk[lastFile] != totalFiles {
			lastFile--
		}
	}

	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		checkSum += i * int(disk[i])
	}
	return checkSum
}
