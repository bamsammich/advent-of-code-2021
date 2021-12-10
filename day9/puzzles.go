package day9

import (
	"advent-of-code-2021/util"
	"log"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

type Coordinate struct {
	X, Y int
}

func parseInput(input []string) ([][]int, error) {
	var heatmap = make([][]int, len(input))
	for row, heights := range input {
		var rowColumns []int
		for _, num := range heights {
			val, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}

			rowColumns = append(rowColumns, val)
		}
		heatmap[row] = rowColumns
	}
	return heatmap, nil
}

func getAdjacents(coord Coordinate, heatmap [][]int) map[Coordinate]int {
	var adjacents = make(map[Coordinate]int)
	if coord.X > 0 {
		adjacents[Coordinate{coord.X - 1, coord.Y}] = heatmap[coord.X-1][coord.Y]
	}
	if coord.Y > 0 {
		adjacents[Coordinate{coord.X, coord.Y - 1}] = heatmap[coord.X][coord.Y-1]
	}
	if coord.X < len(heatmap)-1 {
		adjacents[Coordinate{coord.X + 1, coord.Y}] = heatmap[coord.X+1][coord.Y]
	}
	if coord.Y < len(heatmap[coord.X])-1 {
		adjacents[Coordinate{coord.X, coord.Y + 1}] = heatmap[coord.X][coord.Y+1]
	}
	return adjacents
}

func lowPoints(heatmap [][]int) map[Coordinate]int {
	var lowPoints = make(map[Coordinate]int)
	for row, cols := range heatmap {
		for col, val := range cols {
			var adjacentHeights []int
			var coord = Coordinate{row, col}
			for _, h := range getAdjacents(coord, heatmap) {
				adjacentHeights = append(adjacentHeights, h)
			}
			sort.Ints(adjacentHeights)
			if adjacentHeights[0] > val {
				lowPoints[coord] = val
			}
		}
	}
	return lowPoints
}

func getBasin(coord Coordinate, heatmap [][]int) map[Coordinate]int {
	var (
		location  = heatmap[coord.X][coord.Y]
		adjacents = getAdjacents(coord, heatmap)
		basin     = map[Coordinate]int{
			coord: location,
		}
	)
	for adjacentCoord, h := range adjacents {

		if h > location && h < 9 {
			for k, v := range getBasin(adjacentCoord, heatmap) {
				basin[k] = v
			}
		}
	}
	return basin
}

func puzzle1(data []string) int {
	heatmap, err := parseInput(data)
	if err != nil {
		log.Fatal(err)
	}
	var sum int
	for _, v := range lowPoints(heatmap) {
		sum += v + 1
	}
	return sum
}

func puzzle2(data []string) int {
	heatmap, err := parseInput(data)
	if err != nil {
		log.Fatal(err)
	}
	var (
		lows       = lowPoints(heatmap)
		basinSizes []int
	)

	for coord := range lows {
		basinSizes = append(basinSizes, len(getBasin(coord, heatmap)))
	}
	sort.Ints(basinSizes)
	var product = basinSizes[len(basinSizes)-3]
	for _, v := range basinSizes[len(basinSizes)-2:] {
		product *= v
	}
	return product
}

func Run() {
	var day = 9
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
