package day3

import (
	"log"
	"strconv"

	"github.com/bamsammich/advent-of-code-2021/util"

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

func filterValues(data []string, position int, mostCommon bool) string {
	if len(data) == 1 {
		return data[0]
	}
	var bitMap = map[rune][]string{
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

	if len(bitMap[1]) >= len(bitMap[0]) == mostCommon {
		return filterValues(bitMap[1], position+1, mostCommon)
	}
	if len(bitMap[1]) >= len(bitMap[0]) == !mostCommon {
		return filterValues(bitMap[0], position+1, mostCommon)
	}
	return ""
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
	util.PrintResults(3, puzzle1(data), puzzle2(data))
}
