package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

// [A] sum of all digits that match the next digit in the list.
// [B] The list is circular, so the digit after the last digit is the first digit in the list.
func processLine(line string) {
	var code int
	length := len(line)

	for i := 0 ; i < length ; i++ {
		nextPos := i + 1

		// [B] next digit for last is the first in the list
		if (nextPos == length) {
			nextPos = 0
		}
		
		// skipped error checking
		current, _ := strconv.Atoi(string(line[i]))
		next, _ := strconv.Atoi(string(line[nextPos]))

		if current == next {
			// [A] sum
			code = code + current
		}
	}

	fmt.Println(code)
}


func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		processLine(fscanner.Text())
	}

}
