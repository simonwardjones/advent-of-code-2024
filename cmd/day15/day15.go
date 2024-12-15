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
func expandMap(warehouse [][]rune) [][]rune {
	newWarehouse := [][]rune{}
	for _, line := range warehouse {
		newLine := []rune{}
		for _, cell := range line {
			if cell == '.' {
				newLine = append(newLine, '.', '.')
			} else if cell == '#' {
				newLine = append(newLine, '#', '#')
			} else if cell == '@' {
				newLine = append(newLine, '@', '.')
			} else {
				newLine = append(newLine, '[', ']')
			}
		}
		newWarehouse = append(newWarehouse, newLine)
	}
	return newWarehouse
}

func printWarehouse(warehouse [][]rune) {
	for _, row := range warehouse {
		fmt.Println(string(row))
	}
}

func find(character rune, warehouse [][]rune) (int, int) {
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == character {
				return i, j
			}
		}
	}
	panic("Character not found")
}

func getStep(character rune) (int, int) {
	switch character {
	case '^':
		return -1, 0
	case 'v':
		return 1, 0
	case '<':
		return 0, -1
	case '>':
		return 0, 1
	}
	panic("Invalid step")
}

func part1(data []string) int {
	warehouse, instructions := passInput(data)
	printWarehouse(warehouse)
	row, col := find('@', warehouse)
	fmt.Println("Robot is at", row, col)
	for _, instruction := range instructions {
		dRow, dCol := getStep(instruction)
		// fmt.Println("Step is", dRow, dCol)
		// fmt.Println("Robot is at", row, col)
		nRow, nCol := row+dRow, col+dCol
		steps := 0
		for {
			newSquare := warehouse[nRow+steps*dRow][nCol+steps*dCol]
			// fmt.Println("New square is", string(newSquare))
			if newSquare == 'O' {
				steps++
			} else if newSquare == '.' {
				warehouse[row][col] = '.'
				warehouse[nRow][nCol] = '@'
				for i := 1; i <= steps; i++ {
					warehouse[nRow+i*dRow][nCol+i*dCol] = 'O'
				}
				row, col = nRow, nCol
				break
			} else if newSquare == '#' {
				break
			}
		}
	}
	printWarehouse(warehouse)
	total := 0
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == 'O' {
				total += i*100 + j
			}
		}
	}
	return total

}

func pointsToMove(warehouse [][]rune, instruction rune, row, col int) [][]int {
	dRow, dCol := getStep(instruction)
	newRow, newCol := row+dRow, col+dCol

	toCheck := [][]int{}
	if warehouse[newRow][newCol] == '#' {
		return [][]int{}
	} else if warehouse[newRow][newCol] == ']' {
		toCheck = append(toCheck, []int{newRow, newCol - 1, newCol})
	} else if warehouse[newRow][newCol] == '[' {
		toCheck = append(toCheck, []int{newRow, newCol, newCol + 1})
	} else if warehouse[newRow][newCol] == '.' {
		return [][]int{{row, col}}
	}

	blocks := [][]int{{row, col}}
	for len(toCheck) > 0 {
		block := toCheck[0]
		toCheck = toCheck[1:]
		if instruction == 'v' || instruction == '^' {
			if warehouse[block[0]+dRow][block[1]] == '#' || warehouse[block[0]+dRow][block[2]] == '#' {
				return [][]int{}
			}
			if warehouse[block[0]+dRow][block[1]] == ']' {
				toCheck = append(toCheck, []int{block[0] + dRow, block[1] - 1, block[1]})
			} else if warehouse[block[0]+dRow][block[1]] == '[' {
				toCheck = append(toCheck, []int{block[0] + dRow, block[1], block[2]})
			}
			if warehouse[block[0]+dRow][block[2]] == '[' {
				toCheck = append(toCheck, []int{block[0] + dRow, block[2], block[2] + 1})
			}
		} else if instruction == '<' {
			if warehouse[block[0]][block[1]-1] == '#' {
				return [][]int{}
			} else if warehouse[block[0]][block[1]-1] == ']' {
				toCheck = append(toCheck, []int{block[0], block[1] - 2, block[1] - 1})
			}
		} else if instruction == '>' {
			if warehouse[block[0]][block[2]+1] == '#' {
				return [][]int{}
			} else if warehouse[block[0]][block[2]+1] == '[' {
				toCheck = append(toCheck, []int{block[0], block[2] + 1, block[2] + 2})
			}
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func getGPS(warehouse [][]rune) int {
	total := 0
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == '[' {
				total += i*100 + j
			}
		}
	}
	return total
}

func movePoints(warehouse *[][]rune, points [][]int, instruction rune) *[][]rune {
	dRow, dCol := getStep(instruction)
	for n := len(points) - 1; n >= 0; n-- {
		point := points[n]
		// fmt.Println("Moving point", point)
		if len(point) == 2 {
			(*warehouse)[point[0]+dRow][point[1]+dCol] = '@'
			(*warehouse)[point[0]][point[1]] = '.'
		} else {
			(*warehouse)[point[0]][point[1]] = '.'
			(*warehouse)[point[0]][point[2]] = '.'
			(*warehouse)[point[0]+dRow][point[1]+dCol] = '['
			(*warehouse)[point[0]+dRow][point[2]+dCol] = ']'
		}

	}
	return warehouse
}

func part2(data []string) int {
	warehouse, instructions := passInput(data)
	warehouse = expandMap(warehouse)
	rx, ry := find('@', warehouse)
	for _, instruction := range instructions {
		points := pointsToMove(warehouse, instruction, rx, ry)
		movePoints(&warehouse, points, instruction)
		if len(points) > 0 {
			dx, dy := getStep(instruction)
			rx, ry = rx+dx, ry+dy
		}
	}
	return getGPS(warehouse)
}

func passInput(data []string) ([][]rune, []rune) {
	warehouse := [][]rune{}
	var j int
	for i, line := range data {
		j = i
		if line == "" {
			break
		}
		warehouse = append(warehouse, []rune(line))
	}
	var instructions []rune = []rune(strings.Join(data[j+1:], ""))
	return warehouse, instructions
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
