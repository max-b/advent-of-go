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
	fmt.Println("Part 1")
	total := 0

	lines, err := utils.ReadInputToLines("./inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		total += (number / 3) - 2
	}

	fmt.Println(total)

	fmt.Println("Part 2")

	part2Total := 0
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		part2Total += esotericCalculation(number)
	}

	fmt.Println(part2Total)
}

func esotericCalculation(amount int) int {
	total := 0
	counter := amount
	for counter > 0 {
		counter = (counter / 3) - 2
		if counter > 0 {
			total += counter
		}
	}

	return total
}
