package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const pathInput = "input.txt"
const pathTestInput = "test.txt"

type calcMode int

const (
	CalcPart1 calcMode = iota
	CalcPart2
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

func parseInput(str string) (startPos []int) {
	spl := strings.Split(str, ",")
	for _, valStr := range spl {
		num, err := strconv.ParseInt(valStr, 10, 64)
		logErr(err)

		startPos = append(startPos, int(num))
	}

	sort.Slice(startPos, func(a, b int) bool {
		return startPos[a] < startPos[b]
	})

	return startPos
}

// int median so floating values are ignored
func getMedian(positions []int) (median int) {
	halfIdx := len(positions) / 2
	return positions[halfIdx]
}

func calcSingleScenario(positions []int, inputHeight int, mode calcMode) (result int) {
	result = 0

	for _, h := range positions {
		diff := float64(inputHeight - h)
		absChange := int(math.Abs(diff))
		switch mode {
		case CalcPart1:
			result += absChange
			break
		case CalcPart2:
			for i := 1; i <= absChange; i++ {
				result += i
			}
			break
		default:
			log.Panicln("Unknown mode", mode)
		}

	}

	return result
}

func findOptimum(positions []int, debug bool, mode calcMode) (optimum int) {
	optimize := func(position []int, startInput int, currBest int, modifier int) (optimum int) {
		currInput := startInput

		for {
			currInput = currInput + modifier
			testRes := calcSingleScenario(positions, currInput, mode)

			if testRes >= currBest {
				return currBest
			} else {
				if debug {
					fmt.Printf("Further optimized: %d -> %d\n", currBest, testRes)
				}
				currBest = testRes
			}
		}
	}

	median := getMedian(positions)

	currInput := 0

	medianRes := calcSingleScenario(positions, median, mode)

	// check if values lower than median provide better results
	currInput = median - 1
	testRes := calcSingleScenario(positions, currInput, mode)
	if debug {
		fmt.Printf("Median res => %d - Lower res => %d\n", medianRes, testRes)
	}

	if testRes < medianRes {
		return optimize(positions, currInput, testRes, -1)
	}

	// check if values higher than median provide better results
	currInput = median + 1
	testRes = calcSingleScenario(positions, currInput, mode)
	if debug {
		fmt.Printf("Median res => %d - Higher res => %d\n", medianRes, testRes)
	}

	if testRes < medianRes {
		return optimize(positions, currInput, testRes, +1)
	}

	if debug {
		log.Println("Median was already optimum")
	}

	// median was already optimum
	return medianRes
}

func part1(str string, debug bool) (res int) {
	startPos := parseInput(str)

	if debug {
		log.Println("startPos =>", startPos)
	}

	return findOptimum(startPos, debug, CalcPart1)
}

func part2(str string, debug bool) (res int) {
	startPos := parseInput(str)

	if debug {
		log.Println("startPos =>", startPos)
	}

	return findOptimum(startPos, debug, CalcPart2)
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, true)
	log.Println("Part2 result =>", prod2)
}
