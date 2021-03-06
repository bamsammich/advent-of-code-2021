package day12

import (
	"log"
	"regexp"
	"strings"

	"github.com/bamsammich/advent-of-code-2021/util"

	"github.com/echojc/aocutil"
)

func parseInput(input []string) map[string][]string {
	var relMap = make(map[string][]string)
	for _, relation := range input {
		nodes := strings.Split(relation, "-")
		if nodes[0] == "start" || nodes[1] == "end" {
			relMap[nodes[0]] = append(relMap[nodes[0]], nodes[1])
			continue
		}
		if nodes[1] == "start" || nodes[0] == "end" {
			relMap[nodes[1]] = append(relMap[nodes[1]], nodes[0])
			continue
		}

		relMap[nodes[0]] = append(relMap[nodes[0]], nodes[1])
		relMap[nodes[1]] = append(relMap[nodes[1]], nodes[0])
	}
	return relMap
}

func Traverse(from string, relations map[string][]string, visits map[string]int, forbidVisitingSmallCaveTwice bool) [][]string {
	var (
		neighbors = relations[from]
		output    [][]string
	)
	for _, n := range neighbors {
		if n == "end" {
			output = append(output, []string{n})
			continue
		}
		smallCaveVisitedTwice := forbidVisitingSmallCaveTwice
		if !regexp.MustCompile(`[A-Z]+`).MatchString(n) && visits[n] > 0 {
			if !smallCaveVisitedTwice {
				smallCaveVisitedTwice = true
			} else {
				continue
			}
		}
		visitsCheckpoint := make(map[string]int)
		for k, v := range visits {
			visitsCheckpoint[k] = v
		}
		visitsCheckpoint[n]++
		// not the end or an already-visited small cave
		for _, p := range Traverse(n, relations, visitsCheckpoint, smallCaveVisitedTwice) {
			nodes := append([]string{n}, p...)
			output = append(output, nodes)
		}

	}
	return output
}

func puzzle1(data []string) int {
	return len(Traverse("start", parseInput(data), map[string]int{}, true))
}
func puzzle2(data []string) int {
	return len(Traverse("start", parseInput(data), map[string]int{}, false))
}

func Run() {
	var day = 12
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, day)
	if err != nil {
		log.Fatal(err)
	}

	util.PrintResults(day, puzzle1(data), puzzle2(data))
}
