package day8

import (
	"advent-of-code-2021/util"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

/*
 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg
*/

func decode(patterns []string) map[string]string {
	var (
		output       = make(map[string]string)
		segmentCount = make(map[rune]int)
	)

	for _, p := range patterns {
		for _, seg := range p {
			segmentCount[seg]++
		}
	}
	for _, p := range patterns {
		switch len(p) {
		case 2: // 2 segments: 1
			output[util.SortString(p)] = "1"
		case 3: // 3 segments: 7
			output[util.SortString(p)] = "7"
		case 4: // 4 segments: 4
			output[util.SortString(p)] = "4"
		case 5: // 5 segments: 2,3,5
			for _, seg := range p {
				switch segmentCount[seg] {
				case 4:
					output[util.SortString(p)] = "2"
				case 6:
					output[util.SortString(p)] = "5"
				default:
					if _, ok := output[util.SortString(p)]; !ok {
						output[util.SortString(p)] = "3"
					}
				}
			}
		case 6: // 6 segments: 0,9,6
			var sixSegmentCount = make(map[int]int)
			for _, seg := range p {
				sixSegmentCount[segmentCount[seg]]++
			}
			if sixSegmentCount[4] == 0 {
				// only four numbers have the 'e' segment. The only 6 segment number missing it is 9.
				output[util.SortString(p)] = "9"
			} else if sixSegmentCount[7] == 1 {
				// zero is missing the 'd' segment, but has the 'g' segment. Meaning only on 7-count segment is present
				output[util.SortString(p)] = "0"
			} else {
				output[util.SortString(p)] = "6"
			}
		case 7: // 7: segments: 8
			output[util.SortString(p)] = "8"
		}
	}

	return output
}

func parseInput(input []string) map[string]string {
	var numbers = make(map[string]string)
	for _, entry := range input {
		var (
			entryParts = strings.Split(entry, " | ")
			patterns   = decode(strings.Fields(entryParts[0]))
		)
		for _, seq := range strings.Fields(entryParts[1]) {
			numbers[entryParts[1]] = strings.Join([]string{numbers[entryParts[1]], patterns[util.SortString(seq)]}, "")
		}
	}
	return numbers
}

func puzzle1(data []string) int {
	entries := parseInput(data)
	var count int
	for _, e := range entries {
		count += strings.Count(e, "1") + strings.Count(e, "4") + strings.Count(e, "7") + strings.Count(e, "8")
	}
	return count
}

func puzzle2(data []string) int {
	var sum int
	entries := parseInput(data)
	for _, e := range entries {
		val, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	}
	return sum
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 8)
	if err != nil {
		log.Fatal(err)
	}

	util.PrintResults(6, puzzle1(data), puzzle2(data))
}
