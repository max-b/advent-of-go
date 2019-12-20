package day3

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type wireCrossings struct {
	wire1Steps int
	wire2Steps int
}

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// Run - yay I run
func Run() {
	fmt.Println("Day 3")
	fmt.Println("Part 1")
	content, err := ioutil.ReadFile("./inputs/day3.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	wiresStrings := strings.Split(string(content), "\n")
	wire1Codes := strings.Split(strings.TrimSpace(string(wiresStrings[0])), ",")
	wire2Codes := strings.Split(strings.TrimSpace(string(wiresStrings[1])), ",")

	wireMap := make(map[position]wireCrossings)

	updateWireMapFromCodes(wire1Codes, wireMap, 1)
	updateWireMapFromCodes(wire2Codes, wireMap, 2)

	minDistance := maxInt
	for position, value := range wireMap {
		if position.x == 0 && position.y == 0 {
			continue
		}
		if value.wire1Steps > 0 && value.wire2Steps > 0 {
			distance := absInt(position.x) + absInt(position.y)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}
	fmt.Println(minDistance)

	fmt.Println("Part 2")

	minSteps := maxInt
	for position, value := range wireMap {
		if position.x == 0 && position.y == 0 {
			continue
		}
		if value.wire1Steps > 0 && value.wire2Steps > 0 {
			steps := value.wire1Steps + value.wire2Steps
			if steps < minSteps {
				minSteps = steps
			}
		}
	}
	fmt.Println(minSteps)
	// drawWireMap(wireMap)
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func setMap(wireMap map[position]wireCrossings, pos position, iteration int, value int) {
	if iteration == 1 {
		if wireMap[pos].wire1Steps == 0 {
			wireMap[pos] = wireCrossings{
				wire2Steps: wireMap[pos].wire2Steps,
				wire1Steps: value,
			}
		}

	} else if iteration == 2 {
		if wireMap[pos].wire2Steps == 0 {
			wireMap[pos] = wireCrossings{
				wire1Steps: wireMap[pos].wire1Steps,
				wire2Steps: value,
			}
		}
	}
}

func updateWireMapFromCodes(wireCodes []string, wireMap map[position]wireCrossings, iteration int) {
	currentPosition := position{
		x: 0,
		y: 0,
	}

	totalSteps := 0
	for _, code := range wireCodes {
		direction := code[0]
		distance, err := strconv.Atoi(code[1:])
		if err != nil {
			log.Fatal(err)
			return
		}

		switch string(direction) {
		case "U":
			top := currentPosition.y + distance
			for i := currentPosition.y; i <= top; i++ {
				currentPosition = position{
					x: currentPosition.x,
					y: i,
				}
				totalSteps++
				setMap(wireMap, currentPosition, iteration, totalSteps)
			}
		case "D":
			bottom := currentPosition.y - distance
			for i := currentPosition.y; i >= bottom; i-- {
				currentPosition = position{
					x: currentPosition.x,
					y: i,
				}
				totalSteps++
				setMap(wireMap, currentPosition, iteration, totalSteps)
			}
		case "R":
			right := currentPosition.x + distance
			for i := currentPosition.x; i <= right; i++ {
				currentPosition = position{
					y: currentPosition.y,
					x: i,
				}
				totalSteps++
				setMap(wireMap, currentPosition, iteration, totalSteps)
			}
		case "L":
			left := currentPosition.x - distance
			for i := currentPosition.x; i >= left; i-- {
				currentPosition = position{
					y: currentPosition.y,
					x: i,
				}
				totalSteps++
				setMap(wireMap, currentPosition, iteration, totalSteps)
			}
		}
	}
}

func drawWireMap(wireMap map[position]wireCrossings) {

	var img = image.NewRGBA(image.Rect(-4000, -4000, 5000, 4000))

	for position, wireCrossing := range wireMap {
		if wireCrossing.wire1Steps > 0 {
			col1 := color.RGBA{255, 0, 0, 255} // Red
			img.Set(position.x, position.y, col1)
		}
		if wireCrossing.wire2Steps > 0 {
			col2 := color.RGBA{0, 255, 0, 255} // Green
			img.Set(position.x, position.y, col2)
		}
	}

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
