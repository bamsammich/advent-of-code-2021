package day5

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

type Graph [1000][1000]int

func (g Graph) CountIntersections() int {
	var intersections int
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] > 1 {
				intersections++
			}
		}
	}
	return intersections
}

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Line struct {
	A, B *Coordinate
}

func (l Line) String() string {
	return fmt.Sprintf("%s -> %s", l.A, l.B)
}

func (l *Line) IsVertical() bool {
	return l.A.X == l.B.X
}

func (l *Line) IsHorizontal() bool {
	return l.A.Y == l.B.Y
}
func makeRange(min, max int) []int {
	var a []int
	if max < min {
		a = make([]int, min-max+1)
	} else {
		a = make([]int, max-min+1)
	}
	for i := range a {
		if max < min {
			a[i] = min - i
		} else {
			a[i] = min + i
		}
	}
	return a
}

func (l *Line) Plot(graph *Graph) {
	if l.IsVertical() {
		// Then x doesn't change
		x := l.A.X

		for _, y := range makeRange(l.A.Y, l.B.Y) {
			graph[x][y]++
		}
	}
	if l.IsHorizontal() {
		// Then y doesn't change
		y := l.A.Y
		for _, x := range makeRange(l.A.X, l.B.X) {
			graph[x][y]++
		}
	}
}

func printGraph(graph Graph) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == 0 {
				fmt.Printf("%s", ".")
				continue
			}
			fmt.Printf("%d", graph[i][j])
		}
		fmt.Print("\n")
	}
}

func parseCoordinate(rawCoord string) (*Coordinate, error) {
	if !regexp.MustCompile(`\d+,\d+`).MatchString(rawCoord) {
		return nil, errors.New("can't parse coordinate; not in the format \"X,Y\" ")
	}
	coordStrings := strings.Split(rawCoord, ",")
	x, err := strconv.Atoi(coordStrings[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(coordStrings[1])
	if err != nil {
		return nil, err
	}

	return &Coordinate{x, y}, nil
}

func parseInput(input []string) ([]Line, error) {
	var lines []Line
	for _, rawLine := range input {
		matches := regexp.MustCompile(`(\d+,\d+) -> (\d+,\d+)`).FindAllStringSubmatch(rawLine, -1)
		a, err := parseCoordinate(matches[0][1])
		if err != nil {
			return lines, err
		}
		b, err := parseCoordinate(matches[0][2])
		if err != nil {
			return lines, err
		}
		lines = append(lines, Line{a, b})
	}
	return lines, nil
}

func puzzle1(input []string) int {
	lines, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	var graph Graph
	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() {
			line.Plot(&graph)
		}
	}

	return graph.CountIntersections()
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 5:")
	fmt.Println("------------------")
	fmt.Printf("Puzzle 1: %v\n", puzzle1(data))
	// fmt.Printf("Puzzle 2: %d\n", puzzle2(data))
	fmt.Println("------------------")
}
