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
	// fmt.Println(data)

	part1Result := part1(data)
	fmt.Println("Part 1:", part1Result)

	part2Result := part2(data)
	fmt.Println("Part 2:", part2Result)

}

type CalibrationEquation struct {
	answer int
	values []int
}

func passInput(data []string) []CalibrationEquation {
	calibrationEquations := []CalibrationEquation{}
	for _, line := range data {
		splitParts := strings.Split(line, ": ")
		answer, _ := strconv.Atoi(splitParts[0])
		values := make([]int, 0)
		for _, value := range strings.Split(splitParts[1], " ") {
			valueInt, _ := strconv.Atoi(value)
			values = append(values, valueInt)
		}
		calibrationEquations = append(calibrationEquations, CalibrationEquation{answer, values})
	}
	return calibrationEquations
}

func part1(data []string) int {
	calibrationEquations := passInput(data)
	total := 0
	for _, equation := range calibrationEquations {
		solution := equation.solution([]string{}, 1)
		if len(solution) > 0 {
			total += equation.answer
		}
	}
	// fmt.Println(calibrationEquations)
	return total
}

func (equation CalibrationEquation) solution(operators []string, part int) []string {
	if len(equation.values) == 1 {
		if equation.values[0] == equation.answer {
			return operators
		}
		return nil
	}
	lastValue := equation.values[len(equation.values)-1]
	newMulAnswer := equation.answer / lastValue
	remainder := equation.answer % lastValue
	if remainder == 0 {
		newEquation := CalibrationEquation{newMulAnswer, equation.values[:len(equation.values)-1]}
		newOperators := append([]string{"*"}, operators...)
		solution := newEquation.solution(newOperators, part)
		if solution != nil {
			return solution
		}
	}
	newEquation := CalibrationEquation{equation.answer - lastValue, equation.values[:len(equation.values)-1]}
	newOperators := append([]string{"+"}, operators...)
	plusSolution := newEquation.solution(newOperators, part)
	if plusSolution != nil {
		return plusSolution
	}
	if part == 1 {
		return nil
	}
	// if answer ends with last value as string
	answerStr := strconv.Itoa(equation.answer)
	if strings.HasSuffix(answerStr, strconv.Itoa(lastValue)) {
		newAnswer, _ := strconv.Atoi(answerStr[:len(answerStr)-len(strconv.Itoa(lastValue))])
		newEquation := CalibrationEquation{newAnswer, equation.values[:len(equation.values)-1]}
		newOperators := append([]string{"||"}, operators...)
		catSolution := newEquation.solution(newOperators, part)
		if catSolution != nil {
			return catSolution
		}
	}
	return nil
}

func part2(data []string) int {
	calibrationEquations := passInput(data)
	total := 0
	for _, equation := range calibrationEquations {
		solution := equation.solution([]string{}, 2)
		if len(solution) > 0 {
			fmt.Println("Equation:", equation, "Solution:", solution)
			total += equation.answer
		}
	}
	return total
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
