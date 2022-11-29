package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type guardAction struct {
	id     int
	action string
	time   time.Time
}

type guardData struct {
	sleepTime      int
	sleepOnMinutes [60]int
}

var guardActions []guardAction

func findTopSleeper(g map[int]guardData) (id int) {
	var maxSleep int
	for guardId, gd := range g {
		if gd.sleepTime > maxSleep {
			maxSleep, id = gd.sleepTime, guardId
		}
	}
	return
}

func findMaxSleepMinute(gd guardData) (maxMinute int) {
	for i, count := range gd.sleepOnMinutes {
		if count > gd.sleepOnMinutes[maxMinute] {
			maxMinute = i
		}
	}
	return
}

func findMaxMinuteSleeperIDandMinute(g map[int]guardData) (maxGuardId, maxMinute int) {
	var maxSleepRepeat int
	for id, gd := range g {
		for curMinute, count := range gd.sleepOnMinutes {
			if count > maxSleepRepeat {
				maxGuardId, maxMinute, maxSleepRepeat = id, curMinute, count
			}
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		timeStr := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}`).FindString(s)
		const layout = "2006-01-02 15:04"
		t, _ := time.Parse(layout, timeStr)

		actionStr := regexp.MustCompile(`begins shift|falls asleep|wakes up`).FindString(s)
		action := guardAction{id: -1, time: t, action: actionStr}

		if actionStr == "begins shift" {
			idStr := regexp.MustCompile(`#\d*`).FindString(s)
			id, _ := strconv.Atoi(idStr[1:])
			action.id = id
		}

		guardActions = append(guardActions, action)
	}

	sort.Slice(guardActions, func(i, j int) bool {
		return guardActions[i].time.Before(guardActions[j].time)
	})

	allGuardData := make(map[int]guardData)

	var id, sleepStart int = -1, 0

	for _, guardAction := range guardActions {
		switch guardAction.action {
		case "begins shift":
			id = guardAction.id
			if _, ok := allGuardData[id]; !ok {
				allGuardData[id] = guardData{}
			}
		case "falls asleep":
			sleepStart = guardAction.time.Minute()
		case "wakes up":
			if guard, ok := allGuardData[id]; ok {
				guard.sleepTime += guardAction.time.Minute() - sleepStart
				for m := sleepStart; m < guardAction.time.Minute(); m++ {
					guard.sleepOnMinutes[m]++
				}
				allGuardData[id] = guard

			}
		}
	}

	sleeperId := findTopSleeper(allGuardData)
	fmt.Println("part one:", findMaxSleepMinute(allGuardData[sleeperId])*sleeperId)

	maxSleepId, minute := findMaxMinuteSleeperIDandMinute(allGuardData)
	fmt.Println("part two:", maxSleepId*minute)

}
