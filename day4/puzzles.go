package day4

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/bamsammich/advent-of-code-2021/util"

	"github.com/echojc/aocutil"
)

type Game struct {
	Sequence map[int]int
	Boards   []Board
}

type Win struct {
	Board       Board
	WinningMark int
}

func (g *Game) Winners() map[int][]Win {
	var wins = make(map[int][]Win)
	for seq := 0; seq < len(g.Sequence); seq++ {
		for i := 0; i < len(g.Boards); i++ {
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					board := g.Boards[i]
					if board[row][col] == g.Sequence[seq] {
						board[row][col] = -1
						g.Boards[i] = board
						if board.Won() {
							g.Boards[i] = Board{}
							wins[seq] = append(wins[seq], Win{board, g.Sequence[seq]})
						}
					}
				}
			}
		}
	}
	return wins
}

func (g *Game) FirstWinner() []Win {
	winners := g.Winners()
	for i := 0; i < len(winners); i++ {
		if wins, ok := winners[i]; ok {
			return wins
		}
	}
	return nil
}

func (g *Game) LastWinner() []Win {
	winners := g.Winners()
	var turns []int
	for turn := range winners {
		turns = append(turns, turn)
	}
	sort.Ints(turns)
	return winners[turns[len(turns)-1]]
}

type Board [5][5]int

func sum(numbers [5]int) int {
	var out = 0
	for _, v := range numbers {
		out += v
	}
	return out
}

func (b Board) Score() int {
	var sum = 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] != -1 {
				sum += b[i][j]
			}
		}
	}
	return sum
}

func (b Board) Won() bool {
	return b.ColumnsWin() || b.RowsWin()
}

func (b Board) ColumnsWin() bool {
	for col := 0; col < 5; col++ {
		var colVals = [5]int{}
		for row := 0; row < 5; row++ {
			colVals[row] = b[row][col]
		}
		if sum(colVals) == -5 {
			return true
		}
	}
	return false
}

func (b Board) RowsWin() bool {
	for i := 0; i < 5; i++ {
		if sum(b[i]) == -5 {
			return true
		}
	}
	return false
}

func parseInput(input []string) Game {
	var game = Game{
		Sequence: make(map[int]int),
		Boards:   []Board{},
	}
	for i, draw := range strings.Split(input[0], ",") {
		drawNum, _ := strconv.Atoi(draw)
		game.Sequence[i] = drawNum
	}
	var (
		row          = 0
		currentBoard = Board{}
	)
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			game.Boards = append(game.Boards, currentBoard)
			currentBoard = Board{}
			row = 0
			continue
		}

		for col, val := range strings.Fields(input[i]) {
			valNum, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			currentBoard[row][col] = valNum
		}
		row++
	}
	return game
}

func puzzle1(input []string) int {
	game := parseInput(input)
	winners := game.FirstWinner()
	return winners[0].Board.Score() * winners[0].WinningMark
}

func puzzle2(input []string) int {
	game := parseInput(input)
	winners := game.LastWinner()
	return winners[0].Board.Score() * winners[0].WinningMark
}

func Run() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 4)
	if err != nil {
		log.Fatal(err)
	}
	util.PrintResults(4, puzzle1(data), puzzle2(data))
}
