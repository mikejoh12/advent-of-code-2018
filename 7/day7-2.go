package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type worker struct {
	isWorking     bool
	workStep      rune
	jobFinishTime int
}

func arePriorStepsCompleted(finished map[rune]bool, steps []rune) bool {
	for _, r := range steps {
		if _, ok := finished[r]; !ok {
			return false
		}
	}
	return true
}

func getStepTime(r rune) int {
	return 60 + int(r) - 64
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

	workers := make([]worker, 5)
	stepsHaveBegan := make(map[rune]bool)
	var orderStr string
	var time int
	
	for {
		// check for and process complete jobs
		for i, w := range workers {
			if w.isWorking && w.jobFinishTime == time {
				orderStr += string(w.workStep)
				finished[w.workStep] = true
				workers[i] = worker{}
			}
		}

		// assign jobs to idle workers
		for i, w := range workers {
			if !w.isWorking {
				for _, r := range stepsSlice {

					if _, ok := finished[r]; !ok {
						needComplete, _ := stepOrder[r]
						_, isStarted := stepsHaveBegan[r]
						canDoStep := arePriorStepsCompleted(finished, needComplete)
						if !isStarted && canDoStep {
							newWorker := worker{
								isWorking:     true,
								workStep:      r,
								jobFinishTime: time + getStepTime(r),
							}
							stepsHaveBegan[r] = true
							workers[i] = newWorker
							break
						}
					}
				}
			}
		}

		if len(orderStr) >= len(stepsSlice) {
			break
		}

		time++
	}
	fmt.Println("time", time)
}