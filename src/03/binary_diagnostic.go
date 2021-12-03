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

type binaryIndex struct {
	zeros int
	ones  int
}

func coutOccurrences(spl []string) (result []binaryIndex) {
	digits := len(spl[0])

	// init slice of results
	for i := 0; i < digits; i++ {
		result = append(result, binaryIndex{0, 0})
	}

	for _, line := range spl {
		for i, rn := range line {
			switch rn {
			case '0':
				result[i].zeros++
				break
			case '1':
				result[i].ones++
				break
			default:
				log.Panicln("Unhandled rune", rn)
			}
		}
	}

	return result
}

func binaryToInt(binary string) (dec int) {
	res, err := strconv.ParseInt(binary, 2, 64)
	logErr(err)

	return int(res)
}

func part1(spl []string, debug bool) (product int) {
	res := coutOccurrences(spl)
	gammaRate := ""
	epsilonRate := ""

	for i := 0; i < len(res); i++ {
		if res[i].ones == res[i].zeros {
			log.Panicln("Zeroes and Ones are the same =>", res[i].ones, res[i].zeros)
		} else if res[i].ones > res[i].zeros {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	if debug {
		log.Println("gammaRate => ", gammaRate)
		log.Println("epsilonRate => ", epsilonRate)
	}

	gammaInt := binaryToInt(gammaRate)
	epsilonInt := binaryToInt(epsilonRate)

	if debug {
		log.Println("gammaInt => ", gammaInt)
		log.Println("epsilonInt => ", epsilonInt)
	}

	return gammaInt * epsilonInt
}

func main() {
	str := readFile(pathInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	prod1 := part1(splice, true)
	log.Println("Part1 result =>", prod1)

	//prod2 := part2(splice, true)
	//log.Println("Part2 result =>", prod2)
}
