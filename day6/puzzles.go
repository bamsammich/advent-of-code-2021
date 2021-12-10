package day6

import (
	"advent-of-code-2021/util"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func parseInput(input string) ([]int, error) {
	var out []int
	for _, fishStr := range strings.Split(input, ",") {
		fish, err := strconv.Atoi(fishStr)
		if err != nil {
			return out, err
		}
		out = append(out, fish)
	}
	return out, nil
}

// credit: https://github.com/gulliet
// gist: https://gist.github.com/gulliet/599a3e53ba5dc82f38bfa1268c7f637f
// Hadn't heard of memoization and took a long time to understand WHY this works
// before using it.
func memoizedCount() func(int, int) int {
	type fish struct {
		daysRemaining int // Time to Live: when 0 no more processing
		load          int // Internal timer: before spanning new fish
	}
	cache := make(map[fish]int)
	var fn func(int, int) int
	fn = func(daysRemaining int, load int) int {
		daysRemaining = daysRemaining - (load + 1)
		if daysRemaining < 0 {
			return 1
		}
		if _, ok := cache[fish{daysRemaining, load}]; !ok {
			cache[fish{daysRemaining, load}] = fn(daysRemaining, 6) + fn(daysRemaining, 8)
		}
		return cache[fish{daysRemaining, load}]
	}
	return fn
}

func puzzle1(school []int) int {
	var count int
	counter := memoizedCount()
	for _, fish := range school {
		count += counter(80, fish)
	}
	return count
}

func puzzle2(school []int) int {
	var count int
	counter := memoizedCount()
	for _, fish := range school {
		count += counter(256, fish)
	}
	return count
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 6)
	if err != nil {
		log.Fatal(err)
	}
	school, err := parseInput(data[0])
	if err != nil {
		log.Fatal(err)
	}

	util.PrintResults(6, puzzle1(school), puzzle2(school))
}
