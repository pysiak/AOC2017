package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAnagramOf(word, dictWord string) bool {
	lenA := len(word)
	lenB := len(dictWord)
	
	if lenA != lenB || word == dictWord {
		return false
	}
	
	for i := 0; i < lenA; i++ {
		occurencesA := strings.Count(word, string(word[i]))
		occurencesB := strings.Count(dictWord, string(word[i]))

		if occurencesA != occurencesB {
			return false
		}
	}

	return true
}

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

	if len(dict) > 1 {
		for dictWord1, _ := range dict {
			for dictWord2, _ := range dict {
				if isAnagramOf(dictWord1, dictWord2) {
					return false
				}
			}
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

