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

func coutOccurrencesForPosition(spl []string, pos int) (result binaryIndex) {
	for _, line := range spl {
		rn := rune(line[pos])

		switch rn {
		case '0':
			result.zeros++
			break
		case '1':
			result.ones++
			break
		default:
			log.Panicln("Unhandled rune", rn)
		}

	}

	return result
}

func filterSinglePosition(spl []string, position int, value rune) (filteredSpl []string) {
	for _, line := range spl {
		if rune(line[position]) == value {
			filteredSpl = append(filteredSpl, line)
		}
	}

	return filteredSpl
}

func part2(spl []string, debug bool) (product int) {
	var OxySeq string
	var Co2Seq string

	filteredSplOxy := spl
	filteredSplCo2 := spl

	for i := 0; i < len(spl[0]); i++ {
		var filterOxyRune rune
		var filterCo2Rune rune

		if OxySeq == "" {
			resOxy := coutOccurrencesForPosition(filteredSplOxy, i)

			if resOxy.ones == resOxy.zeros {
				filterOxyRune = '1'
			} else if resOxy.ones > resOxy.zeros {
				filterOxyRune = '1'
			} else {
				filterOxyRune = '0'
			}

			filteredSplOxy = filterSinglePosition(filteredSplOxy, i, filterOxyRune)

			if len(filteredSplOxy) == 1 {
				OxySeq = filteredSplOxy[0]
			}
		}

		if Co2Seq == "" {
			resCo2 := coutOccurrencesForPosition(filteredSplCo2, i)

			if resCo2.ones == resCo2.zeros {
				filterCo2Rune = '0'
			} else if resCo2.ones > resCo2.zeros {
				filterCo2Rune = '0'
			} else {
				filterCo2Rune = '1'
			}

			filteredSplCo2 = filterSinglePosition(filteredSplCo2, i, filterCo2Rune)

			if len(filteredSplCo2) == 1 {
				Co2Seq = filteredSplCo2[0]
			}
		}
	}

	if debug {
		log.Println("OxySeq => ", OxySeq)
		log.Println("Co2Seq => ", Co2Seq)
	}

	oxyInt := binaryToInt(OxySeq)
	co2Int := binaryToInt(Co2Seq)

	if debug {
		log.Println("oxyInt => ", oxyInt)
		log.Println("co2Int => ", co2Int)
	}

	return oxyInt * co2Int
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
