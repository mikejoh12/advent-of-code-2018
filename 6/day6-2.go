package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absDiffInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getManhattanDist(pos1, pos2 [2]int) int {
	return absDiffInt(pos1[0], pos2[0]) + absDiffInt(pos1[1], pos2[1])
}


func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maxX, maxY int
	coords := make([][2]int, 0)

	for scanner.Scan() {
		s := scanner.Text()
		var x, y int
		fmt.Sscanf(s, "%d, %d", &x, &y)
		coords = append(coords, [2]int{x, y})
		maxX, maxY = maxInt(maxX, x), maxInt(maxY, y)
	}

	grid := make([][]int, maxY+1)
	for y := range grid {
		grid[y] = make([]int, maxX+1)
	}

	var areaSize int

	for y, row := range grid {
		for x := range row {			
			var totDist int

			for _, coord := range coords {
				mhDist := getManhattanDist(coord, [2]int{x,y})
				totDist += mhDist
			}

			if totDist < 10000 {
				areaSize++
			}

		}
	}

	fmt.Println(areaSize)
}