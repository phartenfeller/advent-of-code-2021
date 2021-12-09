package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const pathInput = "input.txt"
const pathTestInput = "test.txt"

const riskExtra = 1
const valNotFound = 999999999

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

type boardMapT map[string]int

type board struct {
	m      boardMapT
	width  int
	height int
}

type point struct {
	x          int
	y          int
	isLowPoint bool
	val        int
}

func getMapIdx(x, y int) (idx string) {
	return fmt.Sprintf("%d-%d", x, y)
}

func getMapVal(board board, x int, y int) (val int) {
	idx := getMapIdx(x, y)
	if val, ok := board.m[idx]; ok {
		return val
	} else {
		return valNotFound
	}
}

func parseInput(str string) (board board) {
	lines := strings.Split(str, "\n")

	board.width = len(lines[0])
	board.height = len(lines)
	board.m = make(boardMapT)

	for y, line := range lines {
		log.Println("line", line)
		for x, char := range line {
			num, err := strconv.ParseInt(string(char), 10, 64)
			logErr(err)
			board.m[getMapIdx(x+1, y+1)] = int(num)
		}
	}

	return board
}

func checkIsLowPoint(board board, x int, y int) (p point) {
	idx := getMapIdx(x, y)

	p.x = x
	p.y = y
	p.isLowPoint = false
	p.val = board.m[idx]

	// top
	if getMapVal(board, x-1, y) <= p.val {
		return p
	}

	// right
	if getMapVal(board, x, y+1) <= p.val {
		return p
	}

	// bottom
	if getMapVal(board, x+1, y) <= p.val {
		return p
	}

	// left
	if getMapVal(board, x, y-1) <= p.val {
		return p
	}

	p.isLowPoint = true

	// no lower neighbour found
	return p
}

func part1(str string, debug bool) (res int) {
	board := parseInput(str)

	if debug {
		log.Println("board =>", board)
		for y := 1; y <= board.height; y++ {
			for x := 1; x <= board.width; x++ {
				idx := getMapIdx(x, y)
				fmt.Printf("%d", board.m[idx])
			}
			fmt.Printf("\n")
		}
	}

	res = 0

	for y := 1; y <= board.height; y++ {
		for x := 1; x <= board.width; x++ {
			p := checkIsLowPoint(board, x, y)
			if p.isLowPoint {
				res += p.val + riskExtra
				if debug {
					fmt.Printf("Val %d found at %d,%d\n", p.val+riskExtra, x, y)
				}
			}
		}
	}

	return res
}

func getBasinSize(lowPoint point, bo board, foundMap boardMapT, debug bool) int {
	log.Println("Len", len(foundMap))

	count := 1 // incl point itself

	alreadyFound := func(x int, y int) bool {
		idx := getMapIdx(x, y)
		_, ok := foundMap[idx]

		//fmt.Printf("Already found %d,%d => %t\n", x, y, ok)

		return ok
	}

	// declare function before definition to allow recursion
	var searchSurrounding func(x int, y int)

	searchSurrounding = func(x int, y int) {
		t := point{x: x + 1, y: y}
		r := point{x: x, y: y + 1}
		b := point{x: x - 1, y: y}
		l := point{x: x, y: y - 1}

		surr := []point{t, r, b, l}

		for _, p := range surr {
			// don't get out of bounds and skip already found ones
			if p.x < 1 || p.y < 1 || p.x > bo.width || p.y > bo.height || alreadyFound(p.x, p.y) {
				continue
			}

			val := getMapVal(bo, p.x, p.y)
			if val < 9 {
				idx := getMapIdx(p.x, p.y)
				foundMap[idx] = 1
				count++

				searchSurrounding(p.x, p.y)
			}
		}
	}

	searchSurrounding(lowPoint.x, lowPoint.y)

	if debug {
		fmt.Printf("== Found %d values surrounding %d,%d\n", count, lowPoint.x, lowPoint.y)
	}

	return count
}

func part2(str string, debug bool) (res int) {
	board := parseInput(str)

	if debug {
		log.Println("board =>", board)
		for y := 1; y <= board.height; y++ {
			for x := 1; x <= board.width; x++ {
				idx := getMapIdx(x, y)
				fmt.Printf("%d", board.m[idx])
			}
			fmt.Printf("\n")
		}
	}

	var lowPoints []point
	foundMap := make(boardMapT)

	for y := 1; y <= board.height; y++ {
		for x := 1; x <= board.width; x++ {
			p := checkIsLowPoint(board, x, y)
			if p.isLowPoint {
				lowPoints = append(lowPoints, p)

				idx := getMapIdx(p.x, p.y)
				foundMap[idx] = 1

				if debug {
					fmt.Printf("Val %d found at %d,%d\n", p.val, x, y)
				}
			}
		}
	}

	var basinSizes []int

	for _, p := range lowPoints {
		var size int
		size = getBasinSize(p, board, foundMap, debug)
		basinSizes = append(basinSizes, size)
	}

	sort.Slice(basinSizes, func(a, b int) bool {
		return basinSizes[a] > basinSizes[b]
	})

	if debug {
		log.Println("Basins sorted by size", basinSizes)
	}

	res = 1

	for i := 0; i < 3; i++ {
		res = res * basinSizes[i]
	}

	return res
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, true)
	log.Println("Part2 result =>", prod2)
}
