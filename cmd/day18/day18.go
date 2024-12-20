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

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

type State struct {
	pos   [2]int
	steps int
}

func getMinScore(points map[[2]int]int) int {
	maxX := 70
	// maxX := 6
	frontier := []State{{[2]int{0, 0}, 0}}
	visited := make(map[[2]int]bool)
	for len(frontier) > 0 {
		current := frontier[0]
		pos := current.pos
		frontier = frontier[1:]
		// fmt.Println("Currently at", pos, " with steps ", current.steps)
		if pos == [2]int{maxX, maxX} {
			return current.steps
		}
		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			neighbor := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if neighbor[0] < 0 || neighbor[1] < 0 || neighbor[0] > maxX || neighbor[1] > maxX {
				continue
			}
			if visited[neighbor] {
				continue
			}
			if _, ok := points[neighbor]; ok {
				// cant visit # points
				continue
			}
			frontier = append(frontier, State{neighbor, current.steps + 1})
			visited[neighbor] = true
		}
	}
	return -1
}


func part1(data []string) int {
	points := passInput(data, 1024)
	minScore:= getMinScore(points)
	return minScore
}



func part2(data []string) int {
	i := 1024
	for {
		fmt.Println("Trying with ", i)
		points := passInput(data, i)
		minScore := getMinScore(points)
		if minScore == -1 {
			fmt.Println(data[i-1])
			return -100
		}
		i += 1
	}
}

func passInput(data []string, bytes int) map[[2]int]int  {
	points := make(map[[2]int]int)
	// bytes := 12
	for i, line := range data[:bytes] {
		raw_values := strings.Split(line, ",")
		pair := [2]int{}
		for j, raw_value := range raw_values {
			value, _ := strconv.Atoi(raw_value)
			pair[j] = value
		}
		points[pair] = i
	}
	fmt.Println(len(points))
	return points
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
