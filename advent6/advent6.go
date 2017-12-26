package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const SIZE int = 16
var history map[[SIZE]int]int
var registers [SIZE]int

func readInput(line string) {
	banks := strings.Split(line, "\t")

	for i, el := range banks {
		value , _ := strconv.Atoi(el)
		registers[i] = value
	}

	history = make(map[[SIZE]int]int)
	history[registers] = 1
}

func findBiggest() int {
	max := 0
	maxPos := 0
	for i:=0;i<SIZE;i++ {
		if registers[i] > max {
			max = registers[i]
			maxPos = i
		}
	}
	
	return maxPos
}

func emptyBiggest(pos int) int {
	value := registers[pos]
	registers[pos] = 0

	return value
}

func distribute(amount, pos int) {

	pos++

	for i := 0; i < amount; i++ {
		registers[(pos +i) % SIZE]++
	}

}

func seenBeforeNTimes(n int) bool {
	if count, ok := history[registers]; ok == true {
		if count >= n {
			//fmt.Println(count)
			history[registers]++
			return true
		}
		history[registers]++
		return false
	} else {
		history[registers]++
		return false
	}
}

func process(n int) int {
	steps := 0

	for {
		//fmt.Println(registers)
		pos := findBiggest()
	 	amount := emptyBiggest(pos)
		distribute(amount, pos)
		steps++
		if seenBeforeNTimes(n) {
			return steps
		}
	 	
	}
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	fscanner.Scan()
	input := fscanner.Text()
	readInput(input)
	stepsA := process(1)

	fmt.Println("1st Answer:", stepsA)
	
	readInput(input)
	stepsB := process(2)
	fmt.Println("2nd Answer:", stepsB - stepsA)

}
