// Day 1 advent of code 2020

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
	for _, arg := range os.Args {
		if arg == "--test" || arg == "test" {
			inputFileName = "input_test_1.txt"
			break
		} else if arg == "--test2" || arg == "test2" {
			inputFileName = "input_test_2.txt"
		}
	}

	data := loadInput(inputFileName)
	pagesBefore, updates := loadRulesAndUpdates(data)
	fmt.Println(pagesBefore)
	fmt.Println(updates)

	part1Result := part1(pagesBefore, updates)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(pagesBefore, updates)
	fmt.Println("Part 2:", part2Result)

}

func checkUpdate(pagesBefore map[int]map[int]bool, update []int) bool {
	// for each pos for all pos after check if there is a wrong order.
	n := len(update)
	for i, page := range update {
		for j := i + 1; j < n; j++ {
			wrong := pagesBefore[page][update[j]]
			// fmt.Println("Checking:", page, update[j], "wrong:", wrong)
			if wrong {
				return false
			}
		}
	}
	return true
}

func part1(pagesBefore map[int]map[int]bool, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if checkUpdate(pagesBefore, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func part2(pagesBefore map[int]map[int]bool, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		valid := checkUpdate(pagesBefore, update)
		if !valid {
			update = reorderUpdate(pagesBefore, update)
			sum += update[len(update)/2]
		}
	}
	return sum
}

func reorderUpdate(pagesBefore map[int]map[int]bool, update []int) []int {
	n := len(update)
	for i := range len(update) {
		for j := 0; j < n-i-1; j++ {
			if pagesBefore[update[j]][update[j+1]] {
				update[j], update[j+1] = update[j+1], update[j]
			}
		}
	}
	return update

}

func loadRulesAndUpdates(data []string) (pagesBefore map[int]map[int]bool, updates [][]int) {
	pagesBefore = make(map[int]map[int]bool)
	updates = [][]int{}
	for _, line := range data {
		switch {
		case strings.Contains(line, "|"):
			pages := strings.Split(line, "|")
			pageBefore, _ := strconv.Atoi(pages[0])
			pageAfter, _ := strconv.Atoi(pages[1])
			if pagesBefore[pageAfter] == nil {
				pagesBefore[pageAfter] = make(map[int]bool)
			}
			pagesBefore[pageAfter][pageBefore] = true
		case strings.Contains(line, ","):
			updateStr := strings.Split(line, ",")
			update := []int{}
			for _, updatePageStr := range updateStr {
				updatePage, _ := strconv.Atoi(updatePageStr)
				update = append(update, updatePage)
			}
			updates = append(updates, update)
		}
	}
	return pagesBefore, updates
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
