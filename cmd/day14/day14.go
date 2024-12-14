package main

import (
	"bufio"
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

func Mod(a int, b int) int {
	c := a % b
	if c < 0 {
		return b + c
	}
	return c
}

func part1(data []string) int {
	moves := 100
	// xMax := 11
	// yMax := 7

	xMax := 101
	yMax := 103

	robots := passInput(data)
	counts := map[[2]int]int{}
	for _, robot := range robots {
		x, y, dx, dy := robot[0], robot[1], robot[2], robot[3]
		x = Mod(x+moves*dx, xMax)
		y = Mod(y+moves*dy, yMax)
		counts[[2]int{x, y}]++
	}
	quadrantCounts := [4]int{}
	for cell, count := range counts {
		if cell[0] < xMax/2 && cell[1] < yMax/2 {
			quadrantCounts[0] += count
		} else if cell[0] > xMax/2 && cell[1] < yMax/2 {
			quadrantCounts[1] += count
		} else if cell[0] < xMax/2 && cell[1] > yMax/2 {
			quadrantCounts[2] += count
		} else if cell[0] > xMax/2 && cell[1] > yMax/2 {
			quadrantCounts[3] += count
		}
	}
	prod := 1
	for _, quad := range quadrantCounts {
		prod *= quad
	}
	fmt.Println(quadrantCounts)
	return prod
}

func getTree(counts map[[2]int]int, xMax int, yMax int) string {
    var sb strings.Builder
    sb.Grow(yMax * (xMax + 1)) // Preallocate memory for the string builder

    for i := 0; i < yMax; i++ {
        for j := 0; j < xMax; j++ {
            if counts[[2]int{j, i}] == 0 {
                sb.WriteByte(' ')
            } else {
                sb.WriteByte('#')
            }
        }
        sb.WriteByte('\n')
    }
    return sb.String()
}

func part2(data []string) int {
    xMax := 101
    yMax := 103
    robots := passInput(data)
    for moves := 1; true; moves++ {
        counts := map[[2]int]int{}
        for _, robot := range robots {
            x, y, dx, dy := robot[0], robot[1], robot[2], robot[3]
            x = Mod(x+moves*dx, xMax)
            y = Mod(y+moves*dy, yMax)
            counts[[2]int{x, y}]++
        }
        tree := getTree(counts, xMax, yMax)
        if strings.Contains(tree, "#####################") {
            fmt.Print(tree)
            fmt.Println(moves)
            return moves
        }
    }
    return 100
}

func passInput(data []string) [][4]int {
	// split by \n\n
	robots := [][4]int{}
	for _, line := range data {
		robotNumbers := regexp.MustCompile(`(-?\d+)`).FindAllString(line, -1)
		robot := [4]int{}
		for i, bot := range robotNumbers {
			botInt, _ := strconv.Atoi(bot)
			robot[i] = botInt
		}
		robots = append(robots, robot)
	}
	return robots
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
