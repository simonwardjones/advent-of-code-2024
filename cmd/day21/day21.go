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

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

type Keypad interface {
	getPosition(target rune) [2]int
	getCurrentPosition() [2]int
	setPosition(position [2]int)
	avoidPoint() [2]int
	reset()
}

type NumericKeypad struct {
	currentPosition [2]int
}

func (k *NumericKeypad) getPosition(target rune) [2]int {
	numericKeypadMap := map[rune][2]int{
		'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
		'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
		'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
		'0': {3, 1}, 'A': {3, 2},
	}
	return numericKeypadMap[target]
}

func (k *NumericKeypad) getCurrentPosition() [2]int {
	return k.currentPosition
}

func (k *NumericKeypad) setPosition(position [2]int) {
	k.currentPosition = position
}

func (k *NumericKeypad) avoidPoint() [2]int {
	return [2]int{3, 0}
}

func (k *DirectionalKeypad) reset() {
	k.currentPosition = [2]int{0, 2}
}

func newDirectionalKeypad() *DirectionalKeypad {
	return &DirectionalKeypad{[2]int{0, 2}} // start in A
}

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

type DirectionalKeypad struct {
	currentPosition [2]int
}

func (k *DirectionalKeypad) getPosition(target rune) [2]int {
	directionalKeypadMap := map[rune][2]int{
		'^': {0, 1}, 'A': {0, 2},
		'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
	}
	return directionalKeypadMap[target]
}

func (k *DirectionalKeypad) getCurrentPosition() [2]int {
	return k.currentPosition
}

func (k *DirectionalKeypad) setPosition(position [2]int) {
	k.currentPosition = position
}

func (k *DirectionalKeypad) avoidPoint() [2]int {
	return [2]int{0, 0}
}

func (k *NumericKeypad) reset() {
	k.currentPosition = [2]int{3, 2}
}

func newNumericKeypad() *NumericKeypad {
	return &NumericKeypad{[2]int{3, 2}} // start in A
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getPath(k Keypad, targetPosition [2]int) string {
	currentPosition := k.getCurrentPosition()
	avoid := k.avoidPoint()
	dy := targetPosition[0] - currentPosition[0]
	dx := targetPosition[1] - currentPosition[1]
	xSteps := ""
	for i := 0; i < abs(dx); i++ {
		if dx > 0 {
			xSteps += ">"
		} else {
			xSteps += "<"
		}
	}
	ySteps := ""
	for i := 0; i < abs(dy); i++ {
		if dy > 0 {
			ySteps += "v"
		} else {
			ySteps += "^"
		}
	}
	if currentPosition[0] == avoid[0] && targetPosition[1] == avoid[1] {
		return ySteps + xSteps + "A" // avoid avoid!
	}
	if targetPosition[0] == avoid[0] && currentPosition[1] == avoid[1] {
		return xSteps + ySteps + "A"
	}
	if dx < 0 {
		return xSteps + ySteps + "A"
	}
	return ySteps + xSteps + "A"
}

func moveTo(keypad Keypad, target string) string {
	steps := ""
	for _, char := range target {
		targetPosition := keypad.getPosition(char)
		path := getPath(keypad, targetPosition)
		steps += path
		keypad.setPosition(targetPosition)
	}
	return steps
}

func part1(data []string) int {
	numericPad := newNumericKeypad()
	pads := []Keypad{}
	for range 2 {
		pads = append(pads, newDirectionalKeypad())
	}
	fmt.Println("Pads:", pads)
	mySteps := []string{}
	for _, line := range data {
		steps := moveTo(numericPad, line)
		fmt.Println("Target:", line)
		fmt.Println("Steps:", steps)
		for i, pad := range pads {
			steps = moveTo(pad, steps)
			fmt.Println("Steps from ", i, "are", steps)
		}
		mySteps = append(mySteps, steps)
		for _, pad := range pads {
			pad.reset()
		}
	}
	total := 0
	for i, steps := range mySteps {
		numberStr := strings.Trim(data[i], "A")
		numberStr = strings.TrimLeft(numberStr, "0")
		number, _ := strconv.Atoi(numberStr)
		complexity := len(steps) * number
		fmt.Println("Complexity:", complexity, "=", len(steps), "*", number)
		total += complexity
	}
	return total
}

type State struct {
	code  string
	depth int
}

var cache = map[State]int{}

func stepsRequired(code string, keypad *DirectionalKeypad, depth int) int {
	if depth == 0 {
		return len(code)
	}

	state := State{code, depth}
	if val, ok := cache[state]; ok {
		return val
	}

	steps := 0
	for _, CodeChunk := range strings.SplitAfter(code, "A") {
		fmt.Println("CodeChunk:", CodeChunk)
		codeM1 := moveTo(keypad, CodeChunk)
		steps += stepsRequired(codeM1, keypad, depth-1)
	}

	cache[state] = steps
	return steps
}

func part2(data []string) int {
	numericPad := newNumericKeypad()
	directionalPad := newDirectionalKeypad()
	mySteps := []int{}
	for _, line := range data {
		stepsLiteral := moveTo(numericPad, line)
		fmt.Println("Target:", line)
		fmt.Println("stepsLiteral:", stepsLiteral)
		steps := stepsRequired(stepsLiteral, directionalPad, 25)
		mySteps = append(mySteps, steps)
		fmt.Println("steps:", steps)
	}

	total := 0
	for i, s := range mySteps {
		numberStr := strings.Trim(data[i], "A")
		numberStr = strings.TrimLeft(numberStr, "0")
		number, _ := strconv.Atoi(numberStr)
		complexity := s * number
		fmt.Println("Complexity:", complexity, "=", s, "*", number)
		total += complexity
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

// tooo HIgh
// 300043033609508
// 263492840501566
