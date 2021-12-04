package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const BOARD_SIDE_LEN = 5

const pathInput = "input.txt"
const pathTestInput = "test.txt"

type splitMode int

const (
	WHITESPACE splitMode = 0
	COMMAS               = 1
)

func logErr(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func readFile(path string) (str string) {
	fp, err := filepath.Abs(path)
	logErr(err)
	dat, err := os.ReadFile(fp)
	logErr(err)

	str = string(dat)
	return str
}

type bingoBoardLine struct {
	values []int
	found  int
}

type bingoBoard struct {
	solved bool
	lines  []bingoBoardLine
}

func splitStringsToIntSlice(line string, mode splitMode) (bingoNums []int) {
	var tempSpl []string

	switch mode {
	case WHITESPACE:
		tempSpl = strings.Fields(line)
		break
	case COMMAS:
		tempSpl = strings.Split(line, ",")
		break
	default:
		log.Panicln("Unknown mode =>", mode)

	}

	for _, val := range tempSpl {
		newNum, err := strconv.ParseInt(val, 10, 64)
		logErr(err)

		bingoNums = append(bingoNums, int(newNum))
	}

	return bingoNums
}

func parseBingoBoard(lines []string) (board bingoBoard) {
	var rows [][]int

	board.solved = false

	for _, line := range lines {
		newNums := splitStringsToIntSlice(line, WHITESPACE)
		rows = append(rows, newNums)
	}

	for i := 0; i < len(rows); i++ {
		newRowLine := bingoBoardLine{found: 0, values: rows[i]}
		board.lines = append(board.lines, newRowLine)
	}

	for i := 0; i < len(rows); i++ {
		colNums := make([]int, len(rows))
		for j := 0; j < len(rows); j++ {
			colNums[j] = rows[j][i]
		}

		newColLine := bingoBoardLine{found: 0, values: colNums}
		board.lines = append(board.lines, newColLine)
	}

	if len(board.lines) != 10 {
		log.Panicln("Not 10 lines in board =>", len(board.lines))
	}

	return board
}

func parseInput(spl []string) (bingoNums []int, boards []bingoBoard) {
	var currentBoard []string

	for i, line := range spl {
		if i == 0 {
			bingoNums = splitStringsToIntSlice(line, COMMAS)

		} else if len(line) == 0 {
			if len(currentBoard) != BOARD_SIDE_LEN && i != 1 {
				log.Panicln("currentBoard has a len of =>", len(currentBoard))

			} else if i != 1 {
				newBoard := parseBingoBoard(currentBoard)
				boards = append(boards, newBoard)

				// reset current
				currentBoard = make([]string, 0)
			}

		} else {
			currentBoard = append(currentBoard, line)
		}
	}

	// for last board
	newBoard := parseBingoBoard(currentBoard)
	boards = append(boards, newBoard)

	return bingoNums, boards
}

func lineContainsNum(line []int, num int) bool {
	for _, lineVal := range line {
		if lineVal == num {
			return true
		}
	}

	return false
}

func getWinnerBoardIndex(bingoNums []int, boards []bingoBoard, debug bool) (boardIdx int, numIdx int) {
	// for every drawn number ...
	for i := 0; i < len(bingoNums); i++ {
		num := bingoNums[i]

		// ... check all boards ...
		for j := 0; j < len(boards); j++ {
			// ... and all lines in them
			for k := 0; k < len(boards[j].lines); k++ {
				line := boards[j].lines[k]
				if lineContainsNum(line.values, num) {
					// need to access over loop nums to not update copy
					boards[j].lines[k].found = line.found + 1

					if boards[j].lines[k].found == BOARD_SIDE_LEN {
						if debug {
							log.Println("Winner line in board", j, "=>", line.values)
						}

						return j, i
					}
				}
			}
		}
	}

	// debug output
	for i, board := range boards {
		log.Println("Lines for board", i)
		for j, line := range board.lines {
			log.Println(j, "found => ", line.found, " values => ", line.values)
		}
	}

	log.Panicln("No winner found!")
	return 0, 0
}

func calcResult(winnerBoard bingoBoard, drawnNums []int, debug bool) (res int) {
	// only take rows as with cols values would be duplicated
	relevantRows := winnerBoard.lines[:BOARD_SIDE_LEN]

	if debug {
		log.Println("drawnNumbers", drawnNums)
	}

	sumUnused := 0

	for _, num := range drawnNums {
		for i := 0; i < len(relevantRows); i++ {
			for j := 0; j < len(relevantRows[i].values); j++ {
				if relevantRows[i].values[j] == num {
					relevantRows[i].values[j] = 0
				}
			}
		}
	}

	for _, row := range relevantRows {
		for _, val := range row.values {
			sumUnused += val
		}
	}

	lastDrawn := drawnNums[len(drawnNums)-1]

	if debug {
		log.Println("sumUnused =>", sumUnused)
		log.Println("Last drawn =>", lastDrawn)
	}

	return sumUnused * lastDrawn
}

func part1(spl []string, debug bool) (res int, numBoards int) {
	bingoNums, boards := parseInput(spl)

	if debug {
		log.Println("bingoNums =>", bingoNums)
		log.Println("board count =>", len(boards))
	}

	boardIdx, numIdx := getWinnerBoardIndex(bingoNums, boards, debug)

	res = calcResult(boards[boardIdx], bingoNums[:numIdx+1], debug)

	return res, len(boards)
}

func main() {
	str := readFile(pathInput)
	splice := strings.Split(str, "\n")

	prod1, _ := part1(splice, true)
	log.Println("Part1 result =>", prod1)

	//prod2 := part2(splice, true)
	//log.Println("Part2 result =>", prod2)
}
