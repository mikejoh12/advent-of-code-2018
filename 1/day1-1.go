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

	var f int

	for scanner.Scan() {
		fDiff, _ := strconv.Atoi(scanner.Text())
		f += fDiff
	}

	fmt.Println(f)
}
