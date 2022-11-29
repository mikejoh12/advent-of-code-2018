package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var pairs, triples int

	for scanner.Scan() {
		countMap := make(map[rune]int)
		s := scanner.Text()
		for _, r := range s {
			countMap[r]++
		}

		var pair, triple int
		for _, count := range countMap {
			if count == 2 {
				pair = 1
			} else if count == 3 {
				triple = 1
			}
		}
		pairs += pair
		triples += triple
	}
	fmt.Println(pairs * triples)
}
