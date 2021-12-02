package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

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

func readFile() (str string) {
	dat, err := os.ReadFile("C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\02\\input.txt")
	logErr(err)

	str = string(dat)
	return str
}

func part1(spl []string, debug bool) (horiz int, depth int) {
	horiz = 0
	depth = 0

	for i := 0; i < len(spl); i += 2 {
		dir := spl[i]
		units := strToInt(spl[i+1])

		if debug {
			log.Println("dir =>", dir, ", units =>", units)
		}

		switch dir {
		case "forward":
			horiz += units
			break
		case "down":
			depth += units
		case "up":
			depth -= units
		default:
			log.Panicln("Unknown dir =>", dir)
		}
	}

	return horiz, depth
}

func part2(spl []string, debug bool) (horiz int, depth int) {
	horiz = 0
	depth = 0
	aim := 0

	for i := 0; i < len(spl); i += 2 {
		dir := spl[i]
		units := strToInt(spl[i+1])

		if debug {
			log.Println("dir =>", dir, ", units =>", units)
		}

		switch dir {
		case "forward":
			horiz += units

			if debug {
				log.Println("Fw with current aim =>", aim, " -> depth + ", aim*units)
			}
			depth += aim * units
			break
		case "down":
			aim += units
		case "up":
			aim -= units
		default:
			log.Panicln("Unknown dir =>", dir)
		}
	}

	return horiz, depth
}

func main() {
	str := readFile()
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	horiz, depth := part1(splice, false)
	mult := horiz * depth
	log.Println("Part1: horiz =>", horiz, ", depth => ", depth, ", mult =>", mult)

	horiz2, depth2 := part2(splice, true)
	mult2 := horiz2 * depth2
	log.Println("Part2: horiz =>", horiz2, ", depth => ", depth2, ", mult =>", mult2)
}
