package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const pathInput = "C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\02\\input.txt"
const pathTestInput = "C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\02\\test.txt"

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

func part1(spl []string, debug bool) (product int) {
	horiz := 0
	depth := 0

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

	log.Println("Part1: horiz =>", horiz, ", depth => ", depth)
	return horiz * depth
}

func part2(spl []string, debug bool) (product int) {
	horiz := 0
	depth := 0
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

	log.Println("Part2: horiz =>", horiz, ", depth => ", depth)
	return horiz * depth
}

func main() {
	str := readFile(pathInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	prod1 := part1(splice, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(splice, true)
	log.Println("Part2 result =>", prod2)
}
