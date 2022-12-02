package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func areAllStepsCompleted(finished map[rune]bool, steps []rune) bool {
	for _, r := range steps {
		if _, ok := finished[r]; !ok {
			return false
		}
	}
	return true
}

func main() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
  
	stepOrder := make(map[rune][]rune)
	steps := make(map[rune]bool)
	finished := make(map[rune]bool)

    for scanner.Scan() {
		s := scanner.Text()
		var before, after string
		fmt.Sscanf(s, "Step %s must be finished before step %s can begin.", &before, &after)
		stepOrder[rune(after[0])] = append(stepOrder[rune(after[0])], rune(before[0]))
		steps[rune(before[0])], steps[rune(after[0])] = true, true
    }

	var stepsSlice []rune
	for r := range steps {
		stepsSlice = append(stepsSlice, r)
	}
	sort.Slice(stepsSlice, func(i, j int) bool {
		return stepsSlice[i] < stepsSlice[j]
	})

	var orderStr string
	for len(orderStr) < len(stepsSlice) {
		for _, r := range stepsSlice {
			if _, ok := finished[r]; !ok {
				needComplete, ok := stepOrder[r]
				canDoStep := areAllStepsCompleted(finished, needComplete)
				if !ok || canDoStep {
					orderStr += string(r)
					finished[r] = true
					break
				}
			}
		}
	}
	fmt.Println(orderStr)
}