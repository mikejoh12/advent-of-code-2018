package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var fDiffs []int
	var f int
	fMap := map[int]bool{0: true}

	for scanner.Scan() {
		fDiff, _ := strconv.Atoi(scanner.Text())
		fDiffs = append(fDiffs, fDiff)
	}

	for {
		for _, fDiff := range fDiffs {
			f += fDiff
			if _, ok := fMap[f]; ok {
				fmt.Println(f)
				os.Exit(0)
			}
			fMap[f] = true
		}
	}
}
