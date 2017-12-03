package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"math"
)

func getDistance(x, y int) int {
	dx := int(math.Abs(float64(x)))
	dy := int(math.Abs(float64(y)))
	return dx+dy
}

// r u l l d d r r r u u u l l l l d d d d r r r r r
func processLine(line string) int {

	var max int
	max, _ = strconv.Atoi(line)
	step := 0
	x, y := 0, 0
	direction := 0
	count := 1
	next_break := false

	for i := 0 ; ; i++ {
//		fmt.Printf("%3d %3d %3d (%d, %d) %d\n", i, step, direction, x, y, count)

		if i % 2 == 0 {
			step++
		}

		if (count + step > max) {
			step = max - count
			next_break = true;
		}

		count = count + step
		
		switch direction % 4 {
		case 0:
				x += step;
				break;
		case 1:
				y += step;
				break;
		case 2:
				x -= step;
				break;
		case 3:
				y -= step;
				break;
		}

		if next_break == true {
			return getDistance(x, y)
		}

		direction++
	}

}


func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	var value int
	for fscanner.Scan() {
		value = processLine(fscanner.Text())
	}

	fmt.Println(value)

}
