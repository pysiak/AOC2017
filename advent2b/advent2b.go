// build with go build advent1b.go
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
	var dividend, divisor int
	
	for _, element1 := range slice {
		for _, element2 := range slice {
			if element1 == element2 {
				continue
			}
			dividend, _ = strconv.Atoi(element1)
			divisor, _ = strconv.Atoi(element2)

			if dividend % divisor == 0 {
				return dividend / divisor
			}
			
		}

	}

	return 1

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
