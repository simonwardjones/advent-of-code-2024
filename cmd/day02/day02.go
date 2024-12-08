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
	inputFileName := "input.txt"
	if testFile := false; testFile {
		inputFileName = "input_test_1.txt"
	}
	data := loadInput(inputFileName)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)
}

type Report []int

func parseLine(line string) (Report, error) {
	var reportData []int
	lineItems := strings.Fields(line)
	for _, item := range lineItems {
		value, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int: %v", err)
		}
		reportData = append(reportData, value)
	}
	return reportData, nil
}

func parseReports(data []string) ([]Report, error) {
	var reports []Report
	for _, line := range data {
		report, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func part1(data []string) int {
	reports, err := parseReports(data)
	if err != nil {
		log.Fatal(err)
	}

	var validCount int
	for _, report := range reports {
		if report.check() {
			validCount++
		}
	}
	return validCount
}

type Direction int

const (
	Unset Direction = iota
	Up
	Down
)

func (report Report) check() bool {
	var direction Direction = Unset
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]
		switch direction {
		case Unset:
			if difference > 0 && difference <= 3 {
				direction = Up
			} else if difference < 0 && difference >= -3 {
				direction = Down
			} else {
				return false
			}
		case Up:
			if difference <= 0 || difference > 3 {
				return false
			}
		case Down:
			if difference >= 0 || difference < -3 {
				return false
			}
		}
	}
	return true
}

func part2(data []string) int {
	reports, err := parseReports(data)
	if err != nil {
		log.Fatal(err)
	}

	var validCount int
	for _, report := range reports {
		if report.isValidWithPermutations() {
			validCount++
		}
	}
	return validCount
}

func (report Report) isValidWithPermutations() bool {
	if report.check() {
		return true
	}
	for i := range report {
		perm := make(Report, len(report)-1)
		copy(perm, report[:i])
		copy(perm[i:], report[i+1:])
		if perm.check() {
			return true
		}
	}
	return false
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
