package day11

import (
	"advent-of-code-2021/util"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

type Coordinate struct {
	X, Y int
}

type Squid struct {
	Charges    int
	HasFlashed bool
}

func parseInput(input []string) ([][]int, error) {
	var squidmap = make([][]int, len(input))
	for row, heights := range input {
		var rowColumns []int
		for _, num := range heights {
			val, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}

			rowColumns = append(rowColumns, val)
		}
		squidmap[row] = rowColumns
	}
	return squidmap, nil
}

func getAdjacents(loc Coordinate, squids [][]int) map[Coordinate]int {
	var (
		topmost    = loc.Y == 0
		leftmost   = loc.X == 0
		rightmost  = loc.X == len(squids)-1
		bottommost = loc.Y == len(squids[loc.X])-1
		adjacents  = make(map[Coordinate]int)
	)
	// Cardinals
	if !leftmost {
		adjacents[Coordinate{loc.X - 1, loc.Y}] = squids[loc.X-1][loc.Y]
	}
	if !topmost {
		adjacents[Coordinate{loc.X, loc.Y - 1}] = squids[loc.X][loc.Y-1]
	}
	if !rightmost {
		adjacents[Coordinate{loc.X + 1, loc.Y}] = squids[loc.X+1][loc.Y]
	}
	if !bottommost {
		adjacents[Coordinate{loc.X, loc.Y + 1}] = squids[loc.X][loc.Y+1]
	}
	// Diagonals
	if !topmost && !leftmost {
		adjacents[Coordinate{loc.X - 1, loc.Y - 1}] = squids[loc.X-1][loc.Y-1]
	}
	if !topmost && !rightmost {
		adjacents[Coordinate{loc.X + 1, loc.Y - 1}] = squids[loc.X+1][loc.Y-1]
	}
	if !bottommost && !leftmost {
		adjacents[Coordinate{loc.X - 1, loc.Y + 1}] = squids[loc.X-1][loc.Y+1]
	}
	if !bottommost && !rightmost {
		adjacents[Coordinate{loc.X + 1, loc.Y + 1}] = squids[loc.X+1][loc.Y+1]
	}

	return adjacents
}

func Step(data []string, squids [][]int) ([][]int, map[Coordinate]bool) {
	var (
		flashes = make(map[Coordinate]bool)
		fn      func(Coordinate, [][]int) [][]int
	)
	fn = func(loc Coordinate, squids [][]int) [][]int {
		if _, ok := flashes[loc]; ok {
			return squids
		}

		var squid = squids[loc.X][loc.Y]
		if squid < 9 {
			squids[loc.X][loc.Y] = squid + 1
			return squids
		}

		squids[loc.X][loc.Y] = 0
		flashes[loc] = true
		for coord := range getAdjacents(loc, squids) {
			squids = fn(coord, squids)
		}
		return squids
	}

	for row, cols := range squids {
		for col := range cols {
			squids = fn(Coordinate{row, col}, squids)
		}
	}
	return squids, flashes
}

func puzzle1(data []string) int {
	squidmap, err := parseInput(data)
	if err != nil {
		log.Fatal(err)
	}

	var (
		flashes    map[Coordinate]bool
		flashCount int
	)
	for step := 0; step < 100; step++ {
		squidmap, flashes = Step(data, squidmap)
		for _, v := range flashes {
			if v {
				flashCount++
			}
		}
	}
	return flashCount
}

func puzzle2(data []string) int {
	squidmap, err := parseInput(data)
	if err != nil {
		log.Fatal(err)
	}

	var (
		step   int
		synced = false
	)
SYNCING:
	for !synced {
		step++
		squidmap, _ = Step(data, squidmap)
		for row, cols := range squidmap {
			for col := range cols {
				if squidmap[row][col] != 0 {
					continue SYNCING
				}
			}
		}
		synced = true
	}
	return step
}

func Run() {
	var day = 11
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
