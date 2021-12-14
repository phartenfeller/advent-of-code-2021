package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const pathInput = "input.txt"
const pathTestInput = "test.txt"

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

type boardMapT map[string]bool

type fold struct {
	axis rune
	val  int
}

type board struct {
	m      boardMapT
	width  int
	height int
	folds  []fold
}

type point struct {
	x int
	y int
}

func (b *board) getCount() int {
	count := 0
	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			idx := getMapIdx(x, y)
			if b.m[idx] {
				count++
			}
		}
	}

	return count
}

func (b *board) getVal(idx string) bool {
	val, ok := b.m[idx]
	if !ok {
		log.Panicln("Access to", idx, "does not work")
	}

	return val
}

func (b *board) print() {
	fmt.Printf("==== Board ====\n")
	fmt.Printf("Width: %d, Height: %d, Points: %d\n", b.width, b.height, b.getCount())

	for _, fold := range b.folds {
		fmt.Printf("Fold: %c, %d\n", fold.axis, fold.val)
	}

	if b.height <= 120 && b.width <= 120 {
		for y := 0; y <= b.height; y++ {
			for x := 0; x <= b.width; x++ {
				idx := getMapIdx(x, y)
				if b.m[idx] {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("===============\n")
}

var gBoard board

func strToInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	logErr(err)

	return int(num)
}

func getMapIdx(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (p point) getMapIdx() string {
	return getMapIdx(p.x, p.y)
}

func parseFold(line string) {
	var axis rune
	var val int

	_, err := fmt.Sscanf(line, "fold along %c=%d", &axis, &val)
	logErr(err)

	gBoard.folds = append(gBoard.folds, fold{axis, val})
}

func parseInput(str string) {
	lines := strings.Split(str, "\n")

	gBoard.width = -1
	gBoard.height = -1
	gBoard.m = make(boardMapT)
	gBoard.folds = make([]fold, 0)

	var points []point

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == 'f' {
			parseFold(line)
		} else {
			p := point{}

			for i, s := range strings.Split(line, ",") {
				if i == 0 {
					p.x = strToInt(s)
					if p.x > gBoard.width {
						gBoard.width = p.x
					}
				} else {
					p.y = strToInt(s)
					if p.y > gBoard.height {
						gBoard.height = p.y
					}
				}
			}

			points = append(points, p)
		}
	}

	// because it is zero based
	gBoard.width++
	gBoard.height++

	// init all values with false
	for y := 0; y <= gBoard.height; y++ {
		for x := 0; x <= gBoard.width; x++ {
			idx := getMapIdx(x, y)
			gBoard.m[idx] = false
		}
	}

	// set pts to true
	for _, p := range points {
		idx := p.getMapIdx()
		gBoard.m[idx] = true
	}
}

func foldBoard(onlyFirst bool, debug bool) {
	for i, f := range gBoard.folds {
		if onlyFirst && i > 0 {
			return
		}

		if debug {
			fmt.Printf("Axis: %c, Val: %d\n", f.axis, f.val)
		}

		nBoard := board{}
		nBoard.folds = gBoard.folds[1:]
		nBoard.m = make(boardMapT)

		if f.axis == 'y' {
			nBoard.height = f.val
			nBoard.width = gBoard.width

			// keep top half
			for y := 0; y < nBoard.height; y++ {
				for x := 0; x < nBoard.width; x++ {
					idx := getMapIdx(x, y)
					nBoard.m[idx] = gBoard.getVal(idx)
				}
			}

			// fold lower half
			for y := gBoard.height - 1; y > nBoard.height; y-- {
				for x := 0; x < nBoard.width; x++ {
					oldIdx := getMapIdx(x, y)
					newIdx := getMapIdx(x, nBoard.height-(y-nBoard.height))
					if gBoard.getVal(oldIdx) {
						nBoard.m[newIdx] = true
					}
				}
			}
		} else if f.axis == 'x' {
			nBoard.height = gBoard.height
			nBoard.width = f.val

			// keep left half
			for y := 0; y < nBoard.height; y++ {
				for x := 0; x < nBoard.width; x++ {
					idx := getMapIdx(x, y)
					nBoard.m[idx] = gBoard.getVal(idx)
				}
			}

			// fold right half
			for y := 0; y < nBoard.height; y++ {
				for x := gBoard.width - 1; x > nBoard.width; x-- {
					oldIdx := getMapIdx(x, y)
					newIdx := getMapIdx(nBoard.width-(x-nBoard.width), y)
					if gBoard.getVal(oldIdx) {
						nBoard.m[newIdx] = true
					}
				}

			}
		} else {
			log.Panicln("Unknown char", f.axis)
		}

		if debug {
			nBoard.print()
		}
		gBoard = nBoard

	}
}

func part1(str string, debug bool) (res int) {
	parseInput(str)

	if debug {
		gBoard.print()
	}

	foldBoard(true, true)

	return gBoard.getCount()
}

func part2(str string, debug bool) {
	parseInput(str)

	if debug {
		gBoard.print()
	}

	foldBoard(false, true)
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, false)
	log.Println("Part1 result =>", prod1)

	part2(str, true)
}
