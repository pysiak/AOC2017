package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)


// [B] consider the digit halfway around the circular list.
func processLine(line string) int {

	slice := strings.Split(line, "\t")
	var value, min, max int
	
	for index, element := range slice {
		value, _ = strconv.Atoi(element)

		if index == 0 {
			min = value
			max = value
		}

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}

	}

	return max - min

}


func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	sum := 0
	for fscanner.Scan() {
		value := processLine(fscanner.Text())
		sum += value
	}

	fmt.Println(sum)

}
