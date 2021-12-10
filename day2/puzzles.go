package day2

import (
	"advent-of-code-2021/util"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func puzzle1(input []string) int {
	var x, y = 0, 0
	for _, move := range input {
		moveparts := strings.Split(move, " ")
		distance, _ := strconv.Atoi(moveparts[1])
		switch moveparts[0] {
		case "forward":
			x += distance
		case "up":
			y -= distance
		case "down":
			y += distance
		}
	}
	return (x * y)
}

func puzzle2(input []string) int {
	var x, y, aim = 0, 0, 0
	for _, move := range input {
		moveparts := strings.Split(move, " ")
		distance, _ := strconv.Atoi(moveparts[1])
		switch moveparts[0] {
		case "forward":
			x += distance
			y += distance * aim
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return (x * y)
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 2)
	if err != nil {
		log.Fatal(err)
	}
	util.PrintResults(2, puzzle1(data), puzzle2(data))
}
