package day10

import (
	"fmt"
	"log"
	"sort"

	"github.com/bamsammich/advent-of-code-2021/util"

	"github.com/echojc/aocutil"
)

var closers = map[rune]rune{
	'(': ')', '[': ']', '{': '}', '<': '>',
}
var bracketErrorScore = map[rune]int{
	')': 3, ']': 57, '}': 1197, '>': 25137,
}
var bracketACScore = map[rune]int{
	')': 1, ']': 2, '}': 3, '>': 4,
}

func puzzle1(data []string) int {
	var errorScore = 0
	for _, chunks := range data {
		var opens = []rune{}
	BRACKET:
		for _, bracket := range chunks {
			switch bracket {
			case '(', '{', '<', '[':
				opens = append(opens, bracket)
			case ')', '}', '>', ']':
				if closers[opens[len(opens)-1]] == bracket {
					// Pop the stack
					opens = opens[:len(opens)-1]
				} else {
					errorScore += bracketErrorScore[bracket]
					break BRACKET
				}
			}
		}
	}
	return errorScore
}

func puzzle2(data []string) int {
	var acScores []int
CHUNK:
	for _, chunks := range data {
		var opens = []rune{}
		for _, bracket := range chunks {
			switch bracket {
			case '(', '{', '<', '[':
				opens = append(opens, bracket)
			case ')', '}', '>', ']':
				if closers[opens[len(opens)-1]] == bracket {
					// Pop the stack
					opens = opens[:len(opens)-1]
				} else {
					continue CHUNK
				}
			}
		}
		var acScore = 0
		for i := len(opens) - 1; i >= 0; i-- {
			closingBracket := closers[opens[i]]
			acScore = (acScore * 5) + bracketACScore[closingBracket]
		}
		acScores = append(acScores, acScore)
	}
	sort.Ints(acScores)
	fmt.Println(len(acScores))
	return acScores[(len(acScores)-1)/2]
}

func Run() {
	var day = 10
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
