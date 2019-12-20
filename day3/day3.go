package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/max-b/advent-of-go/utils"
)

type position struct {
	x int
	y int
}

type rectangle struct {
	id       int
	position position
	width    int
	height   int
}

func parseLine(line string) (rectangle, error) {
	re := regexp.MustCompile(`#(\d+) @ (\d+,\d+): (\d+x\d+)`)
	matches := re.FindStringSubmatch(line)
	positions := strings.Split(matches[2], ",")
	areas := strings.Split(matches[3], "x")

	var id, x, y, width, height int
	var err error
	if id, err = strconv.Atoi(matches[1]); err != nil {
		return rectangle{}, err
	}
	if x, err = strconv.Atoi(positions[0]); err != nil {
		return rectangle{}, err
	}
	if y, err = strconv.Atoi(positions[1]); err != nil {
		return rectangle{}, err
	}
	if width, err = strconv.Atoi(areas[0]); err != nil {
		return rectangle{}, err
	}
	if height, err = strconv.Atoi(areas[1]); err != nil {
		return rectangle{}, err
	}

	return rectangle{
			id: id,
			position: position{
				x: x,
				y: y,
			},
			width:  width,
			height: height,
		},
		nil
}

// Run - yay I run
func Run() {
	fmt.Println("Day 3")
	fmt.Println("Part 1")
	lines, err := utils.ReadInputToLines("./inputs/day3.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	var rectangles []rectangle
	for _, line := range lines {
		rectangle, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
			return
		}
		rectangles = append(rectangles, rectangle)
	}

	points := make(map[position]int)

	for _, rectangle := range rectangles {
		for x := rectangle.position.x; x < rectangle.position.x+rectangle.width; x++ {
			for y := rectangle.position.y; y < rectangle.position.y+rectangle.height; y++ {
				p := position{x, y}
				points[p]++
			}
		}
	}

	total := 0

	for _, count := range points {
		if count >= 2 {
			total++
		}
	}
	fmt.Println(total)

	fmt.Println("Part 2")

	for _, rectangle := range rectangles {
		hasOverlap := false
		for x := rectangle.position.x; x < rectangle.position.x+rectangle.width; x++ {
			for y := rectangle.position.y; y < rectangle.position.y+rectangle.height; y++ {
				p := position{x, y}
				if points[p] > 1 {
					hasOverlap = true
				}
			}
		}
		if !hasOverlap {
			fmt.Println(rectangle.id)
		}
	}
}
