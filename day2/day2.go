package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Part 1")

	opcodeArray, err := generateOpcodes()

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	result := getResult(opcodeArray, 12, 2)

	fmt.Println(result)

	fmt.Println("Part 2")
	noun, verb, err := findNounVerb(opcodeArray)

	fmt.Println(noun, verb)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	fmt.Println(100*noun + verb)
}

func findNounVerb(opcodeArray []int) (int, int, error) {
	for i := 0; i < 100; i++ {
		for v := 0; v < 100; v++ {
			opcodeArray, err := generateOpcodes()
			if err != nil {
				log.Fatal(err)
			}
			result := getResult(opcodeArray, i, v)
			if result == 19690720 {
				return i, v, nil
			}
		}
	}
	return 0, 0, nil
}

func generateOpcodes() ([]int, error) {

	content, err := ioutil.ReadFile("./inputs/day2.txt")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var opcodeArray []int
	opcodeStrings := strings.Split(strings.TrimSpace(string(content)), ",")

	for _, v := range opcodeStrings {
		number, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		opcodeArray = append(opcodeArray, number)
	}

	return opcodeArray, nil
}

func getResult(opcodeArray []int, noun int, verb int) int {
	opcodeArray[1] = noun
	opcodeArray[2] = verb

	startPos := 0
	for {
		if opcodeArray[startPos] == 99 {
			break
		}
		storageField := opcodeArray[startPos+3]
		first := opcodeArray[opcodeArray[startPos+1]]
		second := opcodeArray[opcodeArray[startPos+2]]
		if opcodeArray[startPos] == 1 {
			opcodeArray[storageField] = first + second
		} else if opcodeArray[startPos] == 2 {
			opcodeArray[storageField] = first * second
		}
		startPos += 4
	}
	return opcodeArray[0]
}
