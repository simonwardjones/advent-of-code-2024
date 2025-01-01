package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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

func part1(data []string) int {
	connectedTo := map[string][]string{}
	nodes := map[string]bool{}
	for _, line := range data {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		connectedTo[a] = append(connectedTo[a], b)
		connectedTo[b] = append(connectedTo[b], a)
		nodes[a] = true
		nodes[b] = true
	}
	// fmt.Println(connectedTo)
	threes := [][3]string{}
	for node := range nodes {
		for _, neighbour := range connectedTo[node] {
			for _, neighbour2 := range connectedTo[neighbour] {
				if neighbour2 != node && slices.Contains(connectedTo[neighbour2], node) {
					three := []string{node, neighbour, neighbour2}
					slices.Sort(three)
					threeT := [3]string(three)
					if !slices.Contains(threes, threeT) {
						threes = append(threes, threeT)
					}
				}
			}
		}
	}
	fmt.Println(len(threes))
	fmt.Println(threes[:3])
	total := 0
	for _, three := range threes {
		for _, node := range three {
			if node[0] == 't' {
				total++
				break
			}
		}
	}
	return total
}

func part2(data []string) int {
	connectedTo := map[string][]string{}
	nodes := map[string]bool{}
	for _, line := range data {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		connectedTo[a] = append(connectedTo[a], b)
		connectedTo[b] = append(connectedTo[b], a)
		nodes[a] = true
		nodes[b] = true
	}
	// fully connected groups
	fmt.Println(connectedTo)
	fullyConnectedGroups := [][]string{}
	for node := range nodes {
		fullyConnectedGroups = append(fullyConnectedGroups, []string{node})
	}
	i := 0
	for len(fullyConnectedGroups) > 1 {
		i++
		fmt.Println(i, len(fullyConnectedGroups))
		nextFullyConnectedGroups := [][]string{}
		for _, group := range fullyConnectedGroups {
			for _, new := range connectedTo[group[0]] {
				if slices.Contains(group, new) {
					continue
				}
				good := true
				for _, other := range group {
					if !slices.Contains(connectedTo[new], other) {
						good = false
						break
					}
				}
				if !good {
					continue
				}
				newGroup := make([]string, len(group)+1)
				copy(newGroup[1:], group)
				newGroup[0] = new
				slices.Sort(newGroup)
				contains := containsSlice(nextFullyConnectedGroups, newGroup)
				if !contains {
					nextFullyConnectedGroups = append(nextFullyConnectedGroups, newGroup)
				}
			}
		}
		fullyConnectedGroups = nextFullyConnectedGroups
	}
	answer := strings.Join(fullyConnectedGroups[0], ",")
	fmt.Println(answer)
	return len(fullyConnectedGroups[0])
}

func containsSlice(sliceOfSlices [][]string, slice []string) bool {
	for _, s := range sliceOfSlices {
		if slices.Equal(s, slice) {
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
