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
func part1(data []string) int {
	register, commands := passInput(data)
	pointer := 0
	var out []int
	output := []int{}
	for pointer >= 0 && pointer < len(commands)-1 {
		opcode := commands[pointer]
		operand := commands[pointer+1]
		register, pointer, out = evaluate(register, opcode, operand, pointer)
		if out != nil {
			output = append(output, out...)
		}
	}

	fmt.Println(output)
	return 10
}

func part2(data []string) int {
	register, commands := passInput(data)
	a := 1
	nc := len(commands)
	for {
		(*register)["A"] = a
		(*register)["B"] = 0
		(*register)["C"] = 0
		pointer := 0
		var out []int
		output := []int{}
		for pointer >= 0 && pointer < len(commands)-1 {
			opcode := commands[pointer]
			operand := commands[pointer+1]
			// fmt.Println(register, getName(opcode), operand)
			register, pointer, out = evaluate(register, opcode, operand, pointer)
			if out != nil {
				output = append(output, out...)
			}
		}
		no := len(output)
		correct := true
		for i := range output {
			if commands[nc-1-i] != output[no-1-i] {
				correct = false
				break
			}
		}
		if correct {
            if no == nc {
                return a
            }
			a = a << 3
		} else {
			a++
		}
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func getName(opcode int) string {
	switch opcode {
	case 0:
		return "adv"
	case 1:
		return "bxl"
	case 2:
		return "bst"
	case 3:
		return "jnz"
	case 4:
		return "bxc"
	case 5:
		return "out"
	case 6:
		return "bdv"
	case 7:
		return "cdv"
	}
	return "unknown"
}

func evaluate(register *map[string]int, opcode int, operand int, pointer int) (*map[string]int, int, []int) {
	var comboOperand int = operand
	var out []int
	if operand == 4 {
		comboOperand = (*register)["A"]
	} else if operand == 5 {
		comboOperand = (*register)["B"]
	} else if operand == 6 {
		comboOperand = (*register)["C"]
	} else if operand == 7 {
		panic("operand 7!!!")
	}
	switch opcode {
	case 0: // adv
		(*register)["A"] = ((*register)["A"] / IntPow(2, comboOperand))
	case 1: // bxl
		(*register)["B"] = (*register)["B"] ^ operand
	case 2: // bst
		(*register)["B"] = comboOperand % 8
	case 3: // jnz
		if (*register)["A"] == 0 {
			break
		}
		pointer = operand - 2
	case 4: // bxc
		(*register)["B"] = (*register)["B"] ^ (*register)["C"]
	case 5: // out
		out = append(out, comboOperand%8)
	case 6: // bdv
		(*register)["B"] = (*register)["A"] / IntPow(2, comboOperand)
	case 7: // cdv
		(*register)["C"] = (*register)["A"] / IntPow(2, comboOperand)
	}

	pointer += 2
	return register, pointer, out
}

func passInput(data []string) (*map[string]int, []int) {
	A, _ := strconv.Atoi(data[0][12:])
	B, _ := strconv.Atoi(data[1][:12])
	C, _ := strconv.Atoi(data[2][:12])
	register := map[string]int{}
	register["A"] = A
	register["B"] = B
	register["C"] = C
	commands := []int{}
	for _, command := range strings.Split(data[4][9:], ",") {
		command, _ := strconv.Atoi(command)
		commands = append(commands, command)
	}
	return &register, commands
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
