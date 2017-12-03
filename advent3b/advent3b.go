package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"math"
)

type Position struct {
	X       int
	Y       int
}

var path map[Position]int

func getDistance(x, y int) int {
	dx := int(math.Abs(float64(x)))
	dy := int(math.Abs(float64(y)))
	return dx+dy
}

func initMap() {
	var p Position
	p.X = 0
	p.Y = 0

	path[p] = 1
}

func storePoint(x, y int) int {
	var p Position
	p.X = x
	p.Y = y

	sum := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			var q Position
			q.X = x + i
			q.Y = y + j

			if _, err := path[p]; err == false {
				sum += path[q]
			}
		}
	}

	path[p] = sum

	return sum
}

func processLine(line string) int {

	var max int
	max, _ = strconv.Atoi(line)

	x, y := 0, 0
	direction := 0

	initMap()

	for i := 0 ; ; i++ {
		
		fmt.Printf("(%d, %d)\n", x, y)

		var p Position

		switch direction % 4 {
		case 0:
				x += 1;
				p.X = x
				p.Y = y + 1
				break;
		case 1:
				y += 1;
				p.X = x - 1
				p.Y = y
				break;
		case 2:
				x -= 1;
				p.X = x
				p.Y = y - 1
				break;
		case 3:
				y -= 1;
				p.X = x + 1
				p.Y = y
				break;
		}

		value := storePoint(x, y)

		if value > max {
			return value
		}
		
		if _, err := path[p]; err == false {
			direction++
			fmt.Println(x,y, path[p])
		}
		
	}

}


func main() {

	path = make(map[Position]int)
	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	var value int
	for fscanner.Scan() {
		value = processLine(fscanner.Text())
	}

	fmt.Println(value)

}

