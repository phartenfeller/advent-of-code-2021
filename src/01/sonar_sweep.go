package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const pathInput = "C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\01\\input.txt"
const pathTestInput = "C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\01\\test.txt"

func logErr(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func strToInt(str string) (num int) {
	num, err := strconv.Atoi(str)
	logErr(err)

	return num
}

func readFile(path string) (str string) {
	dat, err := os.ReadFile(path)
	logErr(err)

	str = string(dat)
	return str
}

func part1(spl []string, debug bool) (increases int) {
	increases = 0
	prev := -1

	for _, line := range spl {
		num := strToInt(line)
		if prev < num && prev != -1 {
			increases++
		} else {
			if debug {
				log.Println("No increase =>", prev, num)
			}
		}

		prev = num
	}

	return increases
}

func getSumOfLast3(spl []string, start int, debug bool) (sum int) {
	if debug {
		log.Println("get sum of", start, start+1, start+2)
	}

	sum = 0
	sum += strToInt(spl[start])
	sum += strToInt(spl[start+1])
	sum += strToInt(spl[start+2])

	return sum
}

func part2(spl []string, debug bool) (increases int) {
	increases = 0

	// start at 3 because we want to compare spl[0, 1, 2] and spl[1, 2, 3]
	for i := 3; i < len(spl); i++ {
		if debug {
			log.Println(i, "======")
		}

		prev := getSumOfLast3(spl, i-3, debug)
		curr := getSumOfLast3(spl, i-2, debug)

		if prev < curr {
			increases++
		} else {
			if debug {
				log.Println("No increase =>", prev, curr, i)
			}
		}
	}

	return increases
}

func main() {
	str := readFile(pathInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	part1Res := part1(splice, false)
	log.Println("Part1: increases =>", part1Res)

	part2Res := part2(splice, false)
	log.Println("Part2: increases =>", part2Res)
}
