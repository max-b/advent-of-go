package day1

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/max-b/advent-of-go/utils"
)

// Run - I'm a run function 😄
func Run() {
	fmt.Println("Day 1")
	total := 0

	seen := make(map[int]bool)

	foundIt := false

	lines, err := utils.ReadInputToLines("./inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	for !foundIt {
		for _, line := range lines {
			sign := line[0]
			number, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}
			if string(sign) == "-" {
				total -= number
			} else {
				total += number
			}
			_, ok := seen[total]
			if ok {
				fmt.Println(total)
				return
			}
			seen[total] = true
		}
	}
}
