package day1

import (
	"fmt"
	"log"

	"github.com/echojc/aocutil"
)

func puzzle1(input []int) int {
	var (
		increases = 0
		depth     = input[0]
	)
	for i := 1; i < len(input); i++ {
		if input[i] > depth {
			increases++
		}
		depth = input[i]
	}
	return increases
}

func puzzle2(input []int) int {
	var (
		increases = 0
		depth     = sum(input[0:3])
	)
	for i := 1; i < len(input)-2; i++ {
		window := sum(input[i : i+3])
		if window > depth {
			increases++
		}
		depth = window
	}

	return increases
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Ints(2021, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 1:")
	fmt.Println("------------------")
	fmt.Printf("Puzzle 1: %d\n", puzzle1(data))
	fmt.Printf("Puzzle 2: %d\n", puzzle2(data))
	fmt.Println("------------------")
}
