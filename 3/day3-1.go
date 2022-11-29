package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coordCount map[[2]int]int

func addClaim(c coordCount, xOffset, yOffset, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c[[2]int{xOffset + x, yOffset + y}]++
		}
	}
}

func strNrSliceToInt(strNrs []string) (intNrs []int) {
	for _, nr := range strNrs {
		intNr, _ := strconv.Atoi(nr)
		intNrs = append(intNrs, intNr)
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	c := make(coordCount)

	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.FieldsFunc(s, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		claimData := strNrSliceToInt(parts)
		addClaim(c, claimData[1], claimData[2], claimData[3], claimData[4])
	}

	var overlaps int
	for _, count := range c {
		if count > 1 {
			overlaps++
		}
	}
	fmt.Println(overlaps)
}
