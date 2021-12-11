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

const flashThreshold = 10
const part1Flashes = 100

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

type boardEntry struct {
	val     int
	flashed bool
}
type boardMapT map[string]boardEntry

type board struct {
	m      boardMapT
	width  int
	height int
}

var surr []point

func (b *board) addEnergy(x, y int) {
	if (x < 0) || (y < 0) || (x > b.width) || (y > b.height) {
		return
	}

	id := getMapIdx(x, y)
	v, ok := b.m[id]

	if !ok {
		return
	}

	v.val += 1

	if v.val >= flashThreshold && !v.flashed {
		v.flashed = true
		b.m[id] = v
		for _, p := range surr {
			gBoard.addEnergy(x+p.x, y+p.y)
		}
	} else {
		b.m[id] = v
	}
}

func (b *board) countAndResetFlashes() int {
	count := 0

	for k, v := range b.m {
		if v.val >= flashThreshold {
			count++
			v.val = 0
			v.flashed = false
			b.m[k] = v
		}
	}

	return count
}

func (b *board) print() {
	fmt.Printf("==== Board ====\n")
	for y := 1; y <= gBoard.height; y++ {
		for x := 1; x <= gBoard.width; x++ {
			idx := getMapIdx(x, y)
			fmt.Printf("%d", gBoard.m[idx].val)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("===============\n")
}

var gBoard board

type point struct {
	x int
	y int
}

func getMapIdx(x, y int) (idx string) {
	return fmt.Sprintf("%d-%d", x, y)
}

func initSurrPoints() {
	surr = []point{
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: -1, y: 1},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
	}
}

func parseInput(str string) {
	lines := strings.Split(str, "\n")

	gBoard.width = len(lines[0])
	gBoard.height = len(lines)
	gBoard.m = make(boardMapT)

	for y, line := range lines {
		log.Println("line", line)
		for x, char := range line {
			num, err := strconv.ParseInt(string(char), 10, 64)
			logErr(err)
			idx := getMapIdx(x+1, y+1)
			gBoard.m[idx] = boardEntry{
				val:     int(num),
				flashed: false,
			}
		}
	}
}

func simulateStep(step int, debug bool) int {
	for y := 1; y <= gBoard.height; y++ {
		for x := 1; x <= gBoard.width; x++ {
			gBoard.addEnergy(x, y)
		}
	}

	flashes := gBoard.countAndResetFlashes()

	if debug {
		log.Println("step:", step, "flashes", flashes)
	}

	if debug && step == 1 {
		gBoard.print()
	}

	return flashes
}

func simulateAllSteps(steps int, debug bool) int {
	flashes := 0

	for i := 1; i <= steps; i++ {
		flashes += simulateStep(i, debug)
		if debug && i == 10 {
			gBoard.print()
		}
	}

	return flashes
}

func part1(str string, debug bool, part1Flashes int) (res int) {
	initSurrPoints()
	parseInput(str)

	res = simulateAllSteps(part1Flashes, debug)

	return res
}

func findStepSimultan(debug bool) (step int) {
	for i := 1; i <= 99999999; i++ {
		flashes := simulateStep(i, debug)
		if flashes == gBoard.width*gBoard.height {
			return i
		}
	}

	log.Panicln("not found")
	return -1
}

func part2(str string, debug bool) (res int) {
	initSurrPoints()
	parseInput(str)

	res = findStepSimultan(debug)

	return res
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, true, part1Flashes)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, true)
	log.Println("Part2 result =>", prod2)

}
