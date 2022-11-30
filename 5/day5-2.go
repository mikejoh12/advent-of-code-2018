package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPolymerLength(polymer []byte) int {
	processing := true
	for processing {
		var isModified bool
		var newPolymer []byte
		for i := 0; i < len(polymer)-1; i++ {
			if isCharPair(polymer[i], polymer[i+1]) {
				i++
				isModified = true
			} else {
				newPolymer = append(newPolymer, polymer[i])
			}
		}
		polymer = append(newPolymer, polymer[len(polymer)-1])
		processing = isModified
	}
	return len(polymer)
}

func getPolymerWithoutChar(p []byte, nonCapChar rune) []byte {
	var newPolymer []byte
	for _, char := range p {
		if char != byte(nonCapChar) && char != byte(nonCapChar)-32 {
			newPolymer = append(newPolymer, char)
		}
	}
	return newPolymer
}

func isCharPair(char1, char2 byte) bool {
	return char1-char2 == 32 || char2-char1 == 32  
}


func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Bytes()

	var minLength = len(polymer)

	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		length := getPolymerLength(getPolymerWithoutChar(polymer, r))
		if length < minLength {
			minLength = length
		}
	}

	fmt.Println(minLength)
}