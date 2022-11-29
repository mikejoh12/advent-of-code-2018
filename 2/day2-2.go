package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkStringsForDiff1(s1, s2 string) bool {
	var diff int
	for i, r := range s1 {
		if s2[i] != byte(r) {
			diff++
		}
	}
	return diff == 1
}

func getCommonLetters(s1, s2 string) (result string) {
	for i, r := range s1 {
		if r == rune(s2[i]) {
			result += string(r)
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var boxIDs []string

	for scanner.Scan() {
		boxID := scanner.Text()
		boxIDs = append(boxIDs, boxID)
	}

	for i := 0; i < len(boxIDs)-1; i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			if checkStringsForDiff1(boxIDs[i], boxIDs[j]) {
				fmt.Println(getCommonLetters(boxIDs[i], boxIDs[j]))
				os.Exit(0)
			}
		}
	}
}
