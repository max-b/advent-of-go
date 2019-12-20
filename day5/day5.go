package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func removeReactions(polymer []rune) []rune {
	i := 0

	for i < len(polymer)-1 {
		first := polymer[i]
		second := polymer[i+1]
		if unicode.ToUpper(first) == unicode.ToUpper(second) && first != second {
			var newPolymer []rune
			newPolymer = append(newPolymer, polymer[0:i]...)
			newPolymer = append(newPolymer, polymer[i+2:]...)
			return removeReactions(newPolymer)
		}
		i++
	}

	return polymer
}

// Run I'm a run func
func Run() {
	fmt.Println("Day 5")
	fmt.Println("Part 1")
	content, err := ioutil.ReadFile("./inputs/day5.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	text := string(content)
	runes := []rune(text)

	reacted := removeReactions(runes)

	fmt.Println(len(reacted))
}
