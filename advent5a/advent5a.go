package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var maze []int

func processLine(line string) {
	
	if val, err := strconv.Atoi(line); err != nil {
		fmt.Println(err)
		return
	 } else {
		maze = append(maze, val)
	 }

}

func solveMaze() int {
	step := 0
	pos := 0
	
	for {
		prev_pos := pos
		pos += maze[pos]

		maze[prev_pos] = maze[prev_pos] + 1

		step++

		if (pos >= len(maze) || pos < 0) {
			return step
		}

		
	}
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	maze = []int{}
	for fscanner.Scan() {
		processLine(fscanner.Text())
	}

	fmt.Println("Steps:", solveMaze())

}

