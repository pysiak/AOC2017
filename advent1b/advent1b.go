package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

// [A] sum of all digits that match the next digit in the list.
// [B] consider the digit halfway around the circular list.
func processLine(line string) {
	var code int
	length := len(line)
	halfWay := length / 2

	for i := 0 ; i < length ; i++ {
		// [B] next digit is halfWay away
		nextPos := i + halfWay

		if (nextPos == length) {
			nextPos = 0
		}
		
		// skipped error checking
		current, _ := strconv.Atoi(string(line[i]))
		next, _ := strconv.Atoi(string(line[nextPos % length]))

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
