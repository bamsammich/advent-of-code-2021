package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func countOnes(data []string) map[int]int {
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

func countOnesInPosition(data []string, position int) int {
	return countOnes(data)[position]
}

func bitstringToInt(s string) int {
	v, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(v)
}

func puzzle1(input []string) int {
	var (
		gamma   = make([]byte, 12)
		epsilon = make([]byte, 12)
	)
	for pos, oneCount := range countOnes(input) {
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

	return bitstringToInt(string(gamma)) * bitstringToInt(string(epsilon))
}

func filterData(input []string, evalFunc func(ones int, comparator int) int) string {
	var filter string
	for i := 0; i < 12; i++ {
		posOnes := countOnesInPosition(input, i)
		filter += strconv.Itoa(evalFunc(posOnes, len(input)))
		var matches = []string{}
		for _, d := range input {
			if strings.HasPrefix(d, filter) {
				matches = append(matches, d)
			}
		}
		input = matches
		if len(input) == 1 {
			return input[0]
		}
	}
	return ""
}

func filterValues(data []string, position int, mostCommon bool) string {
	// Couldn 't figure this puzzle out, found help here:
	//		https://github.com/lynerist/Advent-of-code-2021-golang
	if len(data) == 1 {
		return data[0]
	}
	var bitMap = map[int][]string{
		0: make([]string, 0),
		1: make([]string, 0),
	}
	for _, val := range data {
		if val[position] == '0' {
			bitMap[0] = append(bitMap[0], val)
		} else {
			bitMap[1] = append(bitMap[1], val)
		}
	}
	if len(bitMap[1]) > len(bitMap[0]) == mostCommon {
		return filterValues(bitMap[1], position+1, mostCommon)
	}
	return filterValues(bitMap[0], position+1, mostCommon)
}

func puzzle2(input []string) int {
	o2Value := filterValues(input, 0, true)
	co2Value := filterValues(input, 0, false)
	return bitstringToInt(o2Value) * bitstringToInt(co2Value)
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
	fmt.Printf("Puzzle 2: %d\n", puzzle2(data))
	fmt.Println("------------------")
}