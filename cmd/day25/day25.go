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

}

func part1(data []string) int {
	boxes := passInput(data)
	isLock := []bool{}
	for _, box := range boxes {
		isLock = append(isLock, box[0] == "#####")
	}
	transposed := [][]string{}
	for _, box := range boxes {
		transposed = append(transposed, transpose(box))
	}
	conversions := [][]int{}
	for _, box := range transposed {
		conversions = append(conversions, convertToInts(box))
	}
	// fmt.Println(isLock)
	// fmt.Println(boxes)
	// fmt.Println(transposed)
	// fmt.Println(conversions)
	total := 0
	for i, lock := range conversions {
		for j, lock2 := range conversions {
			good := true
			if isLock[i] && !isLock[j] {
				for k := range lock {
					if lock[k]+lock2[k] > 7 {
						good = false
						break
					}
				}
				if good {
					total += 1
				}
			}
		}
	}
	return total
}

func convertToInts(box []string) []int {
	out := []int{}
	for _, line := range box {
		out = append(out, strings.Count(line, "#"))
	}
	return out
}


func transpose(box []string) []string {
	transposed := []string{}
	for i := 0; i < len(box[0]); i++ {
		line := ""
		for j := 0; j < len(box); j++ {
			line += string(box[j][i])
		}
		transposed = append(transposed, line)
	}
	return transposed
}

func passInput(data []string) [][]string {
	boxes := [][]string{}
	box := []string{}
	for _, line := range data {
		if line == "" {
			boxes = append(boxes, box)
			box = []string{}
			continue
		}
		box = append(box, line)
	}
	boxes = append(boxes, box)
	return boxes
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
