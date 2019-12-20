package day2

import (
	"fmt"
	"log"

	"github.com/max-b/advent-of-go/utils"
)

func differsByOne(word1 string, word2 string) bool {
	diffs := 0

	for i := range word1 {
		if word2[i] != word1[i] {
			diffs++
		}
	}

	return diffs == 1
}

func hasSomeOfSame(word string) (bool, bool) {
	letters := make(map[string]int)

	has3 := false
	has2 := false

	for _, letter := range word {
		letters[string(letter)]++
	}

	for _, value := range letters {
		if value == 3 {
			has3 = true
		}
		if value == 2 {
			has2 = true
		}
	}
	return has2, has3
}

func has2OfSame(word string) bool {
	return true
}

// Run - I'm a run function 🎉
func Run() {
	fmt.Println("Day 2")
	has2Num := 0
	has3Num := 0

	lines, err := utils.ReadInputToLines("./inputs/day2.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range lines {

		has2, has3 := hasSomeOfSame(line)

		if has2 {
			has2Num++
		}
		if has3 {
			has3Num++
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(has2Num * has3Num)

	for _, line := range lines {
		for _, otherLine := range lines {
			if differsByOne(line, otherLine) {
				fmt.Println("Part 2:")
				fmt.Println(line)
				fmt.Println(otherLine)
				return
			}
		}
	}
}
