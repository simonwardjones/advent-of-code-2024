package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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

func getNumber(wires []Wire, key byte) int {
	result := 0
	filteredWires := []Wire{}
	for _, wire := range wires {
		if wire.key[0] == key {
			filteredWires = append(filteredWires, wire)
		}
	}
	fmt.Println(filteredWires)
	for i, wire := range filteredWires {
		if wire.value == 1 {
			result += (1 << i)
		}
	}
	return result
}

func part1(data []string) int {
	wires, gates := passInputs(data)
	// fmt.Println(wires)
	// fmt.Println(gates)
	wireId := 0
	for wireId < len(wires) {
		gateId := 0
		wire := &wires[wireId]
		for gateId < len(gates) {
			gate := &gates[gateId]
			if out, outValue, ok := gate.trySetValue(wire.key, wire.value); ok {
				wires = append(wires, Wire{key: out, value: outValue})
			}
			gateId++
		}
		wireId++
	}
	slices.SortFunc(wires, sortFunc)
	z := getNumber(wires, 'z')
	return z
}

func convertFromBinary(binaryStr string) int {
	result := 0
	for i, char := range binaryStr {
		if char == '1' {
			result += 1 << i
		}
	}
	return result
}

func convertGatesToDOT(gates []Gate) string {
	dot := "digraph G {\n"
	for _, gate := range gates {
		color := ""
		switch gate.op {
		case "AND":
			color = "green"
		case "OR":
			color = "blue"
		case "XOR":
			color = "red"
		default:
			color = "black"
		}
		dot += fmt.Sprintf("  %s -> %s [color=%s]\n", gate.a, gate.out, color)
		dot += fmt.Sprintf("  %s -> %s [color=%s]\n", gate.b, gate.out, color)
	}
	dot += "}"
	return dot
}

// z37 is blue but should be red
// z19 is green but should be red
// z11 is green but should be red
// jqf,jqf blues and reds should merge

var swaps = map[string]string{
	"z37": "wts",
	"wts": "z37",
	"z11": "wpd",
	"wpd": "z11",
	"z19": "mdd",
	"mdd": "z19",
	"jqf": "skh",
	"skh": "jqf",
}

func swapGate(gate *Gate) {
	if newOut, ok := swaps[gate.out]; ok {
		gate.out = newOut
	}
}

func part2(data []string) int {
	wires, gates := passInputs(data)
	for i := range gates {
		swapGate(&gates[i])
	}
	// fmt.Println(wires)
	// fmt.Println(gates)
	wireId := 0
	for wireId < len(wires) {
		gateId := 0
		wire := &wires[wireId]
		for gateId < len(gates) {
			gate := &gates[gateId]
			if out, outValue, ok := gate.trySetValue(wire.key, wire.value); ok {
				wires = append(wires, Wire{key: out, value: outValue})
			}
			gateId++
		}
		wireId++
	}
	slices.SortFunc(wires, sortFunc)
	z := getNumber(wires, 'z')
	x := getNumber(wires, 'x')
	y := getNumber(wires, 'y')
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	xy := x + y
	fmt.Println("xy:", xy)
	fmt.Println("z:", z)
	fmt.Println("z == xy:", z == xy)
	dot := convertGatesToDOT(gates)
	// write do to file
	file, err := os.Create("cmd/day24/gates.dot")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()
	file.WriteString(dot)

	// print sorted swap keys
	swapKeys := []string{}
	for key := range swaps {
		swapKeys = append(swapKeys, key)
	}
	slices.Sort(swapKeys)
	swapStr := strings.Join(swapKeys, ",")
	fmt.Println("swap keys:", swapStr)
	return z
}

type Gate struct {
	a, b       string
	aVal, bVal int
	aSet, bSet bool
	op         string
	out        string
	outValue   int
	computed   bool
}

func (g *Gate) ready() bool {
	return g.aSet && g.bSet
}

func (g *Gate) trySetValue(key string, value int) (out string, outValue int, ok bool) {
	if key == g.a {
		g.aVal = value
		g.aSet = true
	} else if key == g.b {
		g.bVal = value
		g.bSet = true
	}
	if g.ready() && !g.computed {
		g.compute()
		return g.out, g.outValue, true
	}
	return g.out, g.outValue, false
}

func (g *Gate) compute() {
	switch g.op {
	case "AND":
		g.outValue = g.aVal & g.bVal
	case "OR":
		g.outValue = g.aVal | g.bVal
	case "XOR":
		g.outValue = g.aVal ^ g.bVal
	default:
		panic("unknown operation")
	}
	g.computed = true
}

type Wire struct {
	key   string
	value int
}

func sortFunc(wire1, wire2 Wire) int {
	return strings.Compare(wire1.key, wire2.key)
}

func passInputs(data []string) ([]Wire, []Gate) {
	wires := []Wire{}
	gates := []Gate{}
	i := 0
	for _, line := range data {
		if line == "" {
			break
		}
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])
		wires = append(wires, Wire{key: parts[0], value: value})
		i++
	}
	var a, b, op, out string
	for _, line := range data[i+1:] {
		// match format "x00 AND y00 -> z00"
		fmt.Sscanf(line, "%s %s %s -> %s", &a, &op, &b, &out)
		gates = append(gates, Gate{a: a, b: b, op: op, out: out})
	}
	return wires, gates
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
