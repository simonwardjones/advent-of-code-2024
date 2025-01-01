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
)

func main() {
	inputFileName := getFileName()
	fmt.Println("Using input file:", inputFileName)

	data := loadInput(inputFileName)
	// fmt.Println(data)

	part1Result, part2Result := parts(data)
	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)

}

func mix(value int, secretNumber int) int {
	// bitwise xor
	return value ^ secretNumber
}

func prune(secretNumber int) int {
	// % 16777216
	return secretNumber & (16777216 - 1)
}

func process(secretNumber int) int {
	secretNumber = prune(mix(secretNumber<<6, secretNumber))
	secretNumber = prune(mix(secretNumber>>5, secretNumber))
	secretNumber = prune(mix(secretNumber<<11, secretNumber))
	return secretNumber

}

func parts(data []string) (int, int) {
	total := 0
	sums := map[[4]int]int{}
	for _, line := range data {
		seen := map[[4]int]bool{}
		secretNumber, _ := strconv.Atoi(line)
		price := secretNumber % 10
		sequence := []int{0, 0, 0, 0}
		for i := range 2000 {
			secretNumber = process(secretNumber)
			newPrice := secretNumber % 10
			diff := newPrice - price
			price = newPrice
			sequence = append(sequence[1:], diff)
			if _, ok := seen[[4]int(sequence)]; !ok && i >= 3  {
				sums[[4]int(sequence)] += newPrice
                seen[[4]int(sequence)] = true
			}
		}
		total += secretNumber
	}
	maxSum := 0
	maxSequence := [4]int{}
	for sequence, value := range sums {
		if value > maxSum {
			maxSequence = sequence
			maxSum = value
		}
	}
	fmt.Println(maxSequence, maxSum)
	return total, maxSum
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

func getFileName() string {
	inputFileName := "input.txt"
	for i, arg := range os.Args {
		if strings.Contains(arg, "test") || strings.Contains(arg, "-t") {
			for _, num := range "123456789" {
				if strings.Contains(os.Args[i], string(num)) ||
					(i+1 < len(os.Args) && strings.Contains(os.Args[i+1], string(num))) {
					return fmt.Sprintf("input_test_%s.txt", string(num))
				}
			}
			return "input_test_1.txt"
		}
	}
	return inputFileName
}
