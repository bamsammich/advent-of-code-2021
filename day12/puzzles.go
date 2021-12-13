package day12

import (
	"advent-of-code-2021/util"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/echojc/aocutil"
)

type Node struct {
	Name            string
	RemainingVisits int
	Children        map[string]*Node
}

func parseInput(input []string) int {
	var fn func(n *Node) *Node

	fn = func(n *Node) *Node {
		var remainingRelations []string
		for i, rel := range input {
			if !strings.Contains(rel, n.Name) {
				remainingRelations = append(remainingRelations, rel)
				continue
			}
			for _, name := range strings.Split(rel, "-") {
				if name != n.Name {

					continue
				}
				var child = Node{
					Name:            strings.ReplaceAll(strings.ReplaceAll(rel, n.Name, ""), "-", ""),
					RemainingVisits: 1,
					Children:        make(map[string]*Node),
				}
				if regexp.MustCompile(`[A-Z]+`).MatchString(child.Name) {
					child.RemainingVisits = 2
				}
				input = append(input[:i], input[i+1:]...)
				n.Children[child.Name] = &child
			}
		}
		return n
	}
	startNode := fn(&Node{"start", 1, make(map[string]*Node)})
	fmt.Println(input)
	fmt.Println(startNode)
	return 0
}

func puzzle1(data []string) int {
	parseInput(data)
	return 0
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

	util.PrintResults(day, puzzle1(data), nil)
}
