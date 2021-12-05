package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const pathInput = "input.txt"
const pathTestInput = "test.txt"

type lineType int
type boardType [][]int

func (b boardType) Print() {
	for i := 0; i < len(b); i++ {
		str := ""
		for _, xVal := range b[i] {
			if xVal == 0 {
				str += "."
			} else {
				str += strconv.Itoa(xVal)
			}
		}
		log.Println(str)
	}
}

const (
	STRAIGHT lineType = iota
	DIAGONAL
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

type line struct {
	p1x int
	p1y int
	p2x int
	p2y int
	typ lineType
}

func getNumFromStr(str string) (number int) {
	res, err := strconv.ParseInt(str, 10, 64)
	logErr(err)

	return int(res)
}

func parseInput(spl []string) (lines []line) {
	for _, input := range spl {
		var line line
		sides := strings.Split(input, " -> ")
		p1 := strings.Split(sides[0], ",")
		p2 := strings.Split(sides[1], ",")

		line.p1x = getNumFromStr(p1[0])
		line.p1y = getNumFromStr(p1[1])
		line.p2x = getNumFromStr(p2[0])
		line.p2y = getNumFromStr(p2[1])

		lines = append(lines, line)
	}

	return lines
}

func filterStraightLines(lines []line) (filtered []line) {
	for _, line := range lines {
		if line.p1x == line.p2x || line.p1y == line.p2y {
			line.typ = STRAIGHT
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func getBoundaries(lines []line) (max int) {
	max = 0

	for _, line := range lines {
		var temp = line.p1x
		if line.p1y > temp {
			temp = line.p1y
		}
		if line.p2x > temp {
			temp = line.p2x
		}
		if line.p2y > temp {
			temp = line.p2y
		}

		if temp > max {
			max = temp
		}
	}
	return max
}

func buildBoard(lines []line, debug bool) (board boardType) {
	max := getBoundaries(lines)
	if debug {
		log.Println("Boundaries =>", max, "x", max)
	}

	// they are zero bound
	max++

	board = make(boardType, max)
	for i := range board {
		board[i] = make([]int, max)

		for j := 0; j < len(board[i]); j++ {
			board[i][j] = 0
		}
	}

	for _, line := range lines {
		switch line.typ {
		case STRAIGHT:
			// draw line on y-axis
			if line.p1x == line.p2x {
				start := 0
				end := 0

				if line.p1y > line.p2y {
					start = line.p2y
					end = line.p1y
				} else {
					start = line.p1y
					end = line.p2y
				}

				for i := start; i <= end; i++ {
					board[i][line.p1x]++
				}

				// draw line on y-axis
			} else {
				start := 0
				end := 0

				if line.p1x > line.p2x {
					start = line.p2x
					end = line.p1x
				} else {
					start = line.p1x
					end = line.p2x
				}

				for i := start; i <= end; i++ {
					board[line.p1y][i]++
				}
			}
			break
		case DIAGONAL:
			stepX := 0
			stepY := 0

			currX := line.p1x
			currY := line.p1y

			if line.p1x < line.p2x {
				stepX = 1
			} else if line.p1x > line.p2x {
				stepX = -1
			} else {
				log.Panicln("Equal x values", line)
			}

			if line.p1y < line.p2y {
				stepY = 1
			} else if line.p1y > line.p2y {
				stepY = -1
			} else {
				log.Panicln("Equal y values", line)
			}

			for i := 0; i <= abs(line.p1x-line.p2x); i++ {
				board[currY][currX]++

				currX += stepX
				currY += stepY
			}

			break
		default:
			log.Panicln("Unhandled line type =>", line.typ, line)
		}
	}

	return board
}

func countOccurrencesOverThreshold(board boardType, threshold int) (count int) {
	count = 0

	for _, yAxisCoords := range board {
		for _, coords := range yAxisCoords {
			if coords >= threshold {
				count++
			}
		}
	}

	return count
}

func part1(spl []string, debug bool) (res int) {
	lines := parseInput(spl)

	if debug {
		log.Println("lines =>", lines)
	}

	lines = filterStraightLines(lines)

	if debug {
		log.Println("Only straight lines =>", lines)
	}

	board := buildBoard(lines, debug)

	if debug {
		board.Print()
	}

	res = countOccurrencesOverThreshold(board, 2)

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func filterStraightAndDiagLines(lines []line) (filtered []line) {
	filtered = filterStraightLines(lines)

	for _, line := range lines {
		if abs(line.p1x-line.p2x) == abs(line.p1y-line.p2y) {
			line.typ = DIAGONAL
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func part2(spl []string, debug bool) (res int) {
	lines := parseInput(spl)

	if debug {
		log.Println("lines =>", lines)
	}

	lines = filterStraightAndDiagLines(lines)

	if debug {
		log.Println("Straight and diagonal lines =>", lines)
	}

	board := buildBoard(lines, debug)

	if debug {
		board.Print()
	}

	res = countOccurrencesOverThreshold(board, 2)

	return res
}

func main() {
	str := readFile(pathInput)
	splice := strings.Split(str, "\n")

	prod1 := part1(splice, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(splice, true)
	log.Println("Part2 result =>", prod2)
}
