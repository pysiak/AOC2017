package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func validLine(line string) bool {

	dict := make(map[string]bool)

	words := strings.Split(line, " ")

	for _, word := range words {
		if _, ok := dict[word]; ok {
			return false
		} else {
			dict[word] = true
		}
	}
	
	return true

}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	var count int
	for fscanner.Scan() {
		if validLine(fscanner.Text()) == true {
			count++
		}
	}

	fmt.Println(count)

}

