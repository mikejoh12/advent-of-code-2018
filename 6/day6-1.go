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

func getMaxFiniteArea(grid [][]int, distCounts map[int]int) (area int) {
	for x := 0; x < len(grid[0]); x++ {
		delete(distCounts, grid[0][x])
		delete(distCounts,  grid[len(grid)-1][x])
	}
	for y := 0; y < len(grid); y++ {
		delete(distCounts, grid[y][0])
		delete(distCounts, grid[y][len(grid[0])-1])
	}
	for _, size := range distCounts {
		if size > area {
			area = size
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var maxX, maxY int
	coords := make([][2]int, 0)
	distCounts := make(map[int]int, 0)

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

	for y, row := range grid {
		for x := range row {			
			var minDist, minCoordIdx, coordIdx int = len(grid) + len(grid[0]), -1, -1
			var isTie bool

			for idx, coord := range coords {
				mhDist := getManhattanDist(coord, [2]int{x,y})
				if mhDist < minDist {
					isTie, coordIdx = false, idx
					minDist = mhDist
					minCoordIdx = idx
				} else if mhDist == minDist {
					isTie = true
				}
			}
			if isTie {
				grid[y][x] = -1
			} else {
				grid[y][x] = minCoordIdx
				distCounts[coordIdx]++
			}
		}
	}

	fmt.Println(getMaxFiniteArea(grid, distCounts))
}