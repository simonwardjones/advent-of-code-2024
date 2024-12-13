package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

type Matrix [2][2]int

func (m *Matrix) Determinant() int {
	// ad - bc
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func (m *Matrix) Adjoint() *Matrix {
	return &Matrix{
		{m[1][1], -m[0][1]},
		{-m[1][0], m[0][0]},
	}
}

func (m *Matrix) Inverse() *Matrix {
	det := m.Determinant()
	if det == 0 {
		return nil
	}
	adj := m.Adjoint()
	return &Matrix{
		{adj[0][0] / det, adj[0][1] / det},
		{adj[1][0] / det, adj[1][1] / det},
	}
}

func (m *Matrix) Multiply(vector [2]int) [2]int {
	return [2]int{
		m[0][0]*vector[0] + m[0][1]*vector[1],
		m[1][0]*vector[0] + m[1][1]*vector[1],
	}
}

func getTotals(claws []Claw) int {
	total := 0
	for _, claw := range claws {
		// get the determinant
		det := claw.A.Determinant()
		if det == 0 {
			// fmt.Println("Claw", i, "has no inverse")
			continue
		}
		// get the adjoint
		adj := claw.A.Adjoint()
		preResult := adj.Multiply(claw.Prize)
		if preResult[0]%det != 0 || preResult[1]%det != 0 {
			// fmt.Println("Claw", i, "has no integer solution")
			continue
		}
		result := [2]int{preResult[0] / det, preResult[1] / det}
		total += result[0]*3 + result[1]
	}
	return total
}

func part1(data string) int {
	claws := passInput(data, 1)
	return getTotals(claws)
}

func part2(data string) int {
	claws := passInput(data, 2)
	return getTotals(claws)
}

type Claw struct {
	A     Matrix
	Prize [2]int
}

func passInput(data string, part int) []Claw {
	// split by \n\n
	claws := []Claw{}
	for _, chunk := range strings.Split(data, "\n\n") {
		var matrix Matrix = Matrix{}
		var prizes [2]int = [2]int{}
		for j, line := range strings.Split(chunk, "\n") {
			// extract two integers \d
			numbers := regexp.MustCompile(`\d+`).FindAllString(line, -1)
			for i, numStr := range numbers {
				// convert to int
				// store in matrix
				num, _ := strconv.Atoi(numStr)
				if i < 2 && j < 2 {
					matrix[i][j] = num
				} else {
					if part == 2 {
						num += 10000000000000
					}
					prizes[i] = num
				}

			}
		}
		claw := Claw{A: matrix, Prize: prizes}
		claws = append(claws, claw)
	}
	return claws
}

func loadInput(fileName string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFilePath)
	inputFilePath := filepath.Join(currentDir, fileName)

	// return whole file as a single string
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	return string(content)

}
