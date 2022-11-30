package main

import (
	"bufio"
	"fmt"
	"os"
)

func isCharPair(char1, char2 byte) bool {
	return char1-char2 == 32 || char2-char1 == 32  
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Bytes()

	processing := true
	for processing {
		var isModified bool
		for i := 0; i < len(polymer)-1; i++ {
			if isCharPair(polymer[i], polymer[i+1]) {
				polymer = append(polymer[0:i], polymer[i+2:]...)
				isModified = true
				break
			}
		}
		processing = isModified
	}
	fmt.Println(len(polymer))
}