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

func compressPop(startPop []int, debug bool) (compressed map[int]int) {
	m := make(map[int]int)

	for _, fishTime := range startPop {
		if _, ok := m[fishTime]; ok {
			m[fishTime]++
		} else {
			m[fishTime] = 1
		}
	}

	if debug {
		fmt.Println("==== Start Pop ======")
		for key, element := range m {
			fmt.Println("Key:", key, "=>", "Element:", element)
		}
		fmt.Println("===================")
	}

	return m
}

func getPopSum(popMap map[int]int, debug bool) (count int) {
	count = 0

	for key, element := range popMap {
		count += element

		if debug {
			fmt.Println("Key:", key, "=>", "Element:", element)
		}
	}

	return count
}

func initNextDayMap() (nextDayMap map[int]int) {
	nextDayMap = make(map[int]int)

	for i := 0; i <= NewbornTimer; i++ {
		nextDayMap[i] = 0
	}

	return nextDayMap
}

func simulateDays(startPop []int, days int, debug bool) (popCount int) {
	popMap := compressPop(startPop, debug)

	for i := 0; i < days; i++ {
		nextDayMap := initNextDayMap()

		for j := 0; j <= NewbornTimer; j++ {
			if num, ok := popMap[j]; ok {
				if j == 0 {
					nextDayMap[6] += num
					nextDayMap[NewbornTimer] += num
				} else {
					nextDayMap[j-1] += num
				}
			}
		}

		popMap = nextDayMap
	}

	return getPopSum(popMap, debug)
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

	prod1 := part1(str, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, true)
	log.Println("Part2 result =>", prod2)
}
