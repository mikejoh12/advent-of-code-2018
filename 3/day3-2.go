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

func findNoOverlap(c coordCount, id, xOffset, yOffset, width, height int) (int, bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if count, _ := c[[2]int{xOffset + x, yOffset + y}]; count > 1 {
				return -1, false
			}
		}
	}
	return id, true
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	c := make(coordCount)
	claims := make([][]int, 0)

	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.FieldsFunc(s, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		claimData := strNrSliceToInt(parts)
		claims = append(claims, claimData)
		addClaim(c, claimData[1], claimData[2], claimData[3], claimData[4])
	}

	for _, claim := range claims {
		id, ok := findNoOverlap(c, claim[0], claim[1], claim[2], claim[3], claim[4])
		if ok {
			fmt.Println(id)
			break
		}
	}

}
