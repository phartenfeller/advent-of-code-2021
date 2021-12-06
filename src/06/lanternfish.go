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

const NewbornTimer = 8
const DaysPart1 = 80
const DaysPart2 = 256

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

func parseInput(str string) (startPop []int) {
	spl := strings.Split(str, ",")
	for _, valStr := range spl {
		num, err := strconv.ParseInt(valStr, 10, 64)
		logErr(err)

		startPop = append(startPop, int(num))
	}

	return startPop
}

func simulateDays(startPop []int, days int, debug bool) (popCount int) {
	currentPop := startPop
	addCount := 0

	for i := 0; i < days; i++ {
		addCount = 0

		for j := 0; j < len(currentPop); j++ {
			if currentPop[j] == 0 {
				addCount++
				currentPop[j] = 6
			} else {
				currentPop[j]--
			}
		}

		for i := 0; i < addCount; i++ {
			currentPop = append(currentPop, NewbornTimer)
		}

		if debug {
			log.Println("Day", i, " => ", len(currentPop))
		}
	}

	return len(currentPop)
}

func part1(str string, debug bool) (res int) {
	startPop := parseInput(str)

	if debug {
		log.Println("startPop =>", startPop)
	}

	res = simulateDays(startPop, DaysPart1, debug)
	return res
}

func part2(str string, debug bool) (res int) {
	startPop := parseInput(str)

	if debug {
		log.Println("startPop =>", startPop)
	}

	res = simulateDays(startPop, DaysPart2, debug)
	return res
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, true)
	log.Println("Part1 result =>", prod1)

	//prod2 := part2(splice, true)
	//log.Println("Part2 result =>", prod2)
}
