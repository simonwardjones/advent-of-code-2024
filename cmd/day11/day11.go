package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

func part1(data []string) int {
	stones := passInput(data)
	for range 25 {
		newStones := []string{}
		for _, stone := range stones {
			stoneLen := len(stone)
			switch {
			case stone == "0":
				newStones = append(newStones, "1")
			case stoneLen%2 == 0:
				leftStone := stone[:stoneLen/2]
				rightStone := strings.TrimLeft(stone[stoneLen/2:], "0")
				if len(rightStone) == 0 {
					rightStone = "0"
				}
				newStones = append(newStones, leftStone)
				newStones = append(newStones, rightStone)
			default:
				stoneInt, _ := strconv.Atoi(stone)
				newStones = append(newStones, strconv.Itoa(stoneInt*2024))
			}
		}
		stones = newStones
		// fmt.Println(stones)
	}
	return len(stones)
}

func stonesAfterNBlinks(stone string, blinks int, cache *sync.Map) int {
	if blinks == 0 {
		return 1
	}
	// check cache
	cacheKey := fmt.Sprintf("%s:%d", stone, blinks)
	if val, ok := cache.Load(cacheKey); ok {
		return val.(int)
	}
	// calculate
	var total int
	stoneLen := len(stone)
	switch {
	case stone == "0":
		total = stonesAfterNBlinks("1", blinks-1, cache)
	case stoneLen%2 == 0:
		leftStone := stone[:stoneLen/2]
		rightStone := strings.TrimLeft(stone[stoneLen/2:], "0")
		if len(rightStone) == 0 {
			rightStone = "0"
		}
		total = stonesAfterNBlinks(leftStone, blinks-1, cache) + stonesAfterNBlinks(rightStone, blinks-1, cache)
	default:
		stoneInt, _ := strconv.Atoi(stone)
		total = stonesAfterNBlinks(strconv.Itoa(stoneInt*2024), blinks-1, cache)
	}
	// save to cache
	cache.Store(cacheKey, total)
	return total
}

func part2(data []string) int {
	stones := passInput(data)
	total := 0
	cache := &sync.Map{}
	for _, stone := range stones {
		total += stonesAfterNBlinks(stone, 75, cache)
	}
	return total
}

func passInput(data []string) []string {
	return strings.Split(data[0], " ")
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
