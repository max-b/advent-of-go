package day4

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/max-b/advent-of-go/utils"
)

type record struct {
	guardID      int
	time         time.Time
	action       string
	actionString string
}

func parseLine(line string) (record, error) {
	timeRe := regexp.MustCompile(`\[(.*?)\]`)
	timeString := timeRe.FindStringSubmatch(line)[1]
	layout := "2006-01-02 15:04"
	recordTime, err := time.Parse(layout, timeString)
	if err != nil {
		return record{}, err
	}

	actionRe := regexp.MustCompile(`\[.*?\] (.*?)$`)
	actionString := actionRe.FindStringSubmatch(line)[1]
	return record{
		time:         recordTime,
		actionString: actionString,
	}, nil
}

func annotateRecords(records []record) ([]record, error) {
	guardRe := regexp.MustCompile(`Guard #(\d+)`)
	guardID := 0
	for i, record := range records {
		match := guardRe.FindStringSubmatch(record.actionString)

		if len(match) == 2 {
			var err error
			if guardID, err = strconv.Atoi(match[1]); err != nil {
				return records, err
			}
			records[i].action = "starts shift"
		} else {
			records[i].action = record.actionString
		}
		records[i].guardID = guardID
	}

	return records, nil
}

type recordOfSleep struct {
	amount       int
	minutesSlept map[int]int
}

func countSleep(records []record) map[int]*recordOfSleep {
	sleepRecord := make(map[int]*recordOfSleep)
	guardID := 0
	var start time.Time
	for _, record := range records {
		if record.action == "falls asleep" {
			start = record.time
		} else if record.action == "wakes up" {
			startMin := int(start.Minute())
			endMin := int(record.time.Minute())
			diff := endMin - startMin

			if sleepRecord[guardID] == nil {
				sleepRecord[guardID] = &recordOfSleep{
					amount:       0,
					minutesSlept: make(map[int]int),
				}
			}

			sleepRecord[guardID].amount += diff
			for i := startMin; i < endMin; i++ {
				sleepRecord[guardID].minutesSlept[i]++
			}
		} else {
			guardID = record.guardID
		}
	}

	return sleepRecord
}

func findMaxMinute(slept map[int]int) int {
	minute := 0
	max := 0
	for i, v := range slept {
		if v > max {
			max = v
			minute = i
		}
	}
	return minute
}

// Returns a tuple (guardID, minute)
func findGuardWithMaxSleepMinute(sleepRecords map[int]*recordOfSleep) (int, int) {
	maxMinute := 0
	maxAmount := 0
	guardID := 0
	for id, sleepRecord := range sleepRecords {
		for minute, amount := range sleepRecord.minutesSlept {
			if amount > maxAmount {
				maxMinute = minute
				maxAmount = amount
				guardID = id
			}
		}
	}

	return guardID, maxMinute
}

// Run - yay I run
func Run() {
	fmt.Println("Day 4")
	fmt.Println("Part 1")
	lines, err := utils.ReadInputToLines("./inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	var records []record
	for _, line := range lines {
		record, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
			return
		}
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].time.Before(records[j].time) })

	records, err = annotateRecords(records)

	if err != nil {
		log.Fatal(err)
		return
	}

	sleepCount := countSleep(records)

	maxMinutesSlept := 0
	sleepiestGuardID := 0
	for i, v := range sleepCount {
		if v.amount > maxMinutesSlept {
			maxMinutesSlept = v.amount
			sleepiestGuardID = i
		}
	}

	minute := findMaxMinute(sleepCount[sleepiestGuardID].minutesSlept)

	fmt.Println(sleepiestGuardID * minute)

	fmt.Println("Part 2")

	guardID, maxMinute := findGuardWithMaxSleepMinute(sleepCount)
	fmt.Println(guardID * maxMinute)
}
