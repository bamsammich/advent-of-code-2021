package day7

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/bamsammich/advent-of-code-2021/util"

	"github.com/echojc/aocutil"
)

func parseInput(xPositions string) ([]int, error) {
	var positions []int
	for _, xStr := range strings.Split(xPositions, ",") {
		x, err := strconv.Atoi(xStr)
		if err != nil {
			return positions, err
		}
		positions = append(positions, x)
	}
	return positions, nil
}
func avg(numbers []int) int {
	var sum = 0
	for _, n := range numbers {
		sum += n
	}
	return sum / len(numbers)
}
func puzzle1(positions []int) int {
	sort.Ints(positions)
	var cheapestCost = math.MaxInt64
	for i := 0; i <= positions[len(positions)-1]; i++ {
		var fuelCost = 0
		for _, x := range positions {
			fuelCost += int(math.Abs(float64(i - x)))
		}
		if fuelCost < cheapestCost {
			cheapestCost = fuelCost
		}
	}
	return cheapestCost
}

func puzzle2(positions []int) int {
	sort.Ints(positions)
	var cheapestCost = math.MaxInt64
	for i := 0; i <= positions[len(positions)-1]; i++ {
		var fuelCost = 0
		for _, x := range positions {
			unitsMoved := int(math.Abs(float64(i - x)))
			// Triangular Number
			fuelCost += (unitsMoved * (unitsMoved + 1)) / 2
		}
		if fuelCost < cheapestCost {
			cheapestCost = fuelCost
		}
	}
	return cheapestCost
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 7)
	if err != nil {
		log.Fatal(err)
	}
	xPositions, err := parseInput(data[0])
	if err != nil {
		log.Fatal(err)
	}
	util.PrintResults(7, puzzle1(xPositions), puzzle2(xPositions))
}
