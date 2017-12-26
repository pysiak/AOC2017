package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var registers map[string]int
var largestValue int

func processLine(line string) {
	split := strings.Split(line, " ")

	reg := split[0]
	op := split[1]
	val, _ := strconv.Atoi(split[2])
	target_reg := split[4]
	target_cond := split[5]
	target_val, _ := strconv.Atoi(split[6])

	if _, ok := registers[target_reg]; !ok {
		registers[target_reg] = 0
	}

	if _, ok := registers[reg]; !ok {
		registers[reg] = 0
	}

	if op == "dec" {
		val = -val
	}

	switch target_cond {
	case ">":
		if registers[target_reg] > target_val {
			registers[reg] += val
		}
	case "<":
		if registers[target_reg] < target_val {
			registers[reg] += val
	}
	case ">=":
		if registers[target_reg] >= target_val {
			registers[reg] += val
	}
	case "<=":
		if registers[target_reg] <= target_val {
			registers[reg] += val
	}
	case "!=":
	if registers[target_reg] != target_val {
		registers[reg] += val
	}
	case "==":
		if registers[target_reg] == target_val {
			registers[reg] += val
	}
	}

	currentLargest := findLargestVal()

	if currentLargest > largestValue {
		largestValue = currentLargest
	}
}

func findLargestVal() int {
	max := 0
	
	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	registers = make(map[string]int)
	for fscanner.Scan() {
		processLine(fscanner.Text())
	}

	fmt.Println("1st Answer:", findLargestVal())
	fmt.Println("2nd Answer:", largestValue)
}
