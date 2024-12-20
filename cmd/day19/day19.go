package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	inputFileName := getFileName()
	fmt.Println("Using input file:", inputFileName)

	data := loadInput(inputFileName)
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

func getMaxTowelLen(baseTowels map[string]bool) int {
	maxBaseTowel := 0
	for towel := range baseTowels {
		if len(towel) > maxBaseTowel {
			maxBaseTowel = len(towel)
		}
	}
	return maxBaseTowel
}

func waysToMake(baseTowels map[string]bool, design string, maxTowelLen int) int {
	waysTo := map[string]int{"": 1}
	for i := 1; i <= len(design); i++ {
		for j := 1; j <= maxTowelLen && j <= i; j++ {
			if baseTowels[design[i-j:i]] {
				waysTo[design[:i]] += waysTo[design[:i-j]]
			}
		}
	}
	return waysTo[design]
}

func part1(data []string) int {
	total := 0
	baseTowels, designs := passInput(data)
	maxTowelLen := getMaxTowelLen(baseTowels)
	for _, design := range designs {
		if waysToMake(baseTowels, design, maxTowelLen) > 0 {
			total++
		}
	}
	return total
}

func part2(data []string) int {
	total := 0
	baseTowels, designs := passInput(data)
	maxTowelLen := getMaxTowelLen(baseTowels)
	for _, design := range designs {
		total += waysToMake(baseTowels, design, maxTowelLen)
	}
	return total
}

func passInput(data []string) (map[string]bool, []string) {
	baseTowels := make(map[string]bool)
	for _, towel := range strings.Split(data[0], ", ") {
		baseTowels[towel] = true
	}
	designs := data[2:]
	return baseTowels, designs
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
