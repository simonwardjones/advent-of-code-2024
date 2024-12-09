package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	inputFileName := "input.txt"
	for _, arg := range os.Args {
		if arg == "--test" || arg == "test" {
			inputFileName = "input_test_1.txt"
			break
		} else if arg == "--test2" || arg == "test2" {
			inputFileName = "input_test_2.txt"
		}
	}

	data := loadInput(inputFileName)
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

func squashDisk(diskMap []int) []int {
	i, j := 0, len(diskMap)-1
	for i < j {
		if diskMap[i] == -1 && diskMap[j] != -1 {
			diskMap[i], diskMap[j] = diskMap[j], diskMap[i]
		}
		if diskMap[i] != -1 {
			i++
		}
		if diskMap[j] == -1 {
			j--
		}
	}
	return diskMap
}

func expandDisk(diskMap []int) []int {
	expandedMap := []int{}
	for i, value := range diskMap {
		id := i / 2
		for range value {
			if i%2 == 0 {
				expandedMap = append(expandedMap, id)
			} else {
				expandedMap = append(expandedMap, -1)
			}
		}
	}
	return expandedMap
}

func checkSum(squashedDisk []int) int {
	total := 0
	for i, value := range squashedDisk {
		if value == -1 {
			continue
		}
		total += value * i
	}
	return total
}

func part1(data []string) int {
	diskMap := passInput(data)
	expandedMap := expandDisk(diskMap)
	squashedDisk := squashDisk(expandedMap)
	return checkSum(squashedDisk)
}

func squashDiskFullFile(diskMap []int) []int {
	id := diskMap[len(diskMap)-1]
	for id > 0 {
		// find id group from right
		j := len(diskMap) - 1
		for diskMap[j] != id {
			j--
		}
		// find start of group
		i := j
		for diskMap[i-1] == diskMap[j] && i > 1 {
			i--
		}
		groupLength := j - i + 1
		// fmt.Println("Group:", diskMap[j], "Length:", groupLength)
		for k := 0; k < i; k++ {
			fits := true
			for delta := 0; delta < groupLength; delta++ {
				if diskMap[k+delta] != -1 {
					fits = false
					break
				}
			}
			if fits {
				for delta := 0; delta < groupLength; delta++ {
					diskMap[k+delta], diskMap[i+delta] = diskMap[i+delta], diskMap[k+delta]
				}
				break
			}
		}
		id--

	}
	// fmt.Println(diskMap)
	return diskMap
}

func part2(data []string) int {
	diskMap := passInput(data)
	expandedMap := expandDisk(diskMap)
	squashedDisk := squashDiskFullFile(expandedMap)
	return checkSum(squashedDisk)
}

func passInput(data []string) []int {
	results := []int{}
	for _, char := range data[0] {
		value, _ := strconv.Atoi(string(char))
		results = append(results, value)
	}
	return results
}

func loadInput(fileName string) []string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	inputFilePath := filepath.Join(currentDir, fileName)

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
