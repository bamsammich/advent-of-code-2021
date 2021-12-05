package day3

import (
	"fmt"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

func countOnesInPosition(data []string) map[int]int {
	var out = make(map[int]int)
	for _, val := range data {
		for i, c := range val {
			if c == '1' {
				out[i]++
			}
		}
	}
	return out
}

func count(match int, data []int) int {
	var out int
	for _, num := range data {
		if num == match {
			out++
		}
	}
	return out
}

func puzzle1(input []string) int {
	var (
		gamma   = make([]byte, 12)
		epsilon = make([]byte, 12)
	)
	for pos, oneCount := range countOnesInPosition(input) {
		var (
			gammaByte   = '1'
			epsilonByte = '0'
		)
		if oneCount < len(input)/2 {
			gammaByte = '0'
			epsilonByte = '1'
		}
		gamma[pos] = byte(gammaByte)
		epsilon[pos] = byte(epsilonByte)
	}

	gammaVal, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonVal, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(gammaVal) * int(epsilonVal)
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 3:")
	fmt.Println("------------------")
	fmt.Printf("Puzzle 1: %d\n", puzzle1(data))
	// fmt.Printf("Puzzle 2: %d\n", puzzle2(data))
	fmt.Println("------------------")
}
