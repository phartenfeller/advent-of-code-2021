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

const allLetters = "abcdefg"

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

type inputLine struct {
	patterns     []string
	outputValues []string
}

func parseInput(file string) (lines []inputLine) {
	for _, lineStr := range strings.Split(file, "\n") {
		var line inputLine
		segments := strings.Split(lineStr, " | ")

		for _, pattern := range strings.Split(segments[0], " ") {
			line.patterns = append(line.patterns, pattern)
		}

		for _, value := range strings.Split(segments[1], " ") {
			line.outputValues = append(line.outputValues, value)
		}

		lines = append(lines, line)
	}

	return lines
}

/**
	  1
  2   3
    4
  5   6
    7
*/
func initMap() (m map[int]int32) {
	m = make(map[int]int32)

	for i := 0; i <= len(allLetters); i++ {
		m[i] = 0
	}

	return m
}

func figureOutMapping(line inputLine) (mapping map[int]int32) {
	mapping = initMap()

	var one, four, seven, eight string
	var rightSide, leftSide, bottom []int32

	for _, value := range line.patterns {
		switch len(value) {
		// 1
		case 2:
			one = value
			break
			// 7
		case 3:
			seven = value
			break
			// 4
		case 4:
			four = value
			break
			// 8
		case 7:
			eight = value
			break
		default:
			break

		}
	}

	// get top bar from diff of 1 and 7
	for _, char := range seven {
		if strings.Contains(one, string(char)) {
			rightSide = append(rightSide, char)
		} else {
			// top bar
			mapping[1] = char
		}
	}

	// get left bars of 4 from diff of right bars from 1 and 7
	for _, char := range four {
		if !strings.Contains(one, string(char)) {
			leftSide = append(leftSide, char)
		}
	}

	// get bottom bars of 8 from not yet found in 7 and 4
	for _, char := range eight {
		if !strings.Contains(seven, string(char)) && !strings.Contains(four, string(char)) {
			bottom = append(bottom, char)
		}
	}

	// find number nine from input and determine lower left
	// +
	// find number six and determine top right
	for _, pattern := range line.patterns {
		// len 6 and not 0
		if len(pattern) == 6 && strings.Contains(pattern, string(leftSide[0])) && strings.Contains(pattern, string(leftSide[1])) {
			// both right bars -> 9
			if strings.Contains(pattern, string(rightSide[0])) && strings.Contains(pattern, string(rightSide[1])) {
				for _, char := range allLetters {
					if !strings.Contains(pattern, string(char)) {
						mapping[5] = char
					}
				}
				// else 6
			} else {
				for _, char := range allLetters {
					if !strings.Contains(pattern, string(char)) {
						mapping[3] = char
					}
				}
			}
		}
	}

	// because we now know top right we can determine bottom right
	for _, char := range rightSide {
		if char != mapping[3] {
			mapping[6] = char
		}
	}

	// because we now bottom left we can determine bottom
	for _, char := range bottom {
		if char != mapping[5] {
			mapping[7] = char
		}
	}

	/**
		We now have

		  1
	  ?   3
	    ?
	  5   6
	    7
	*/

	// get middle bar from three
	for _, pattern := range line.patterns {
		_a := string(mapping[1])
		_b := string(mapping[3])
		_c := string(mapping[6])
		_d := string(mapping[7])
		log.Println(_a, _b, _c, _d)
		if len(pattern) == 5 && strings.Contains(pattern, string(mapping[1])) && strings.Contains(pattern, string(mapping[3])) && strings.Contains(pattern, string(mapping[6])) && strings.Contains(pattern, string(mapping[7])) {
			for _, char := range pattern {
				if char != mapping[1] && char != mapping[3] && char != mapping[6] && char != mapping[7] {
					mapping[4] = char
				}
			}
		}
	}

	// get last one by checking which we don't already have
	for _, char := range allLetters {
		if mapping[1] != char && mapping[3] != char && mapping[4] != char && mapping[5] != char && mapping[6] != char && mapping[7] != char {
			mapping[2] = char
		}
	}

	return mapping
}

func getBarVal(char int32, mapping map[int]int32) (val int) {
	for i := 1; i <= len(allLetters); i++ {
		if char == mapping[i] {
			// times ten to not make 1 useless in product
			return i * 10
		}
	}

	log.Panicln("Nothing found for char =>", char, string(char))
	return 0
}

/**
	  1
  2   3
    4
  5   6
    7
*/
func getNumberFromBarVal(val int, numCode string) (num int) {
	switch val {
	case 30 * 60:
		return 1
	case 10 * 30 * 40 * 50 * 70:
		return 2
	case 10 * 30 * 40 * 60 * 70:
		return 3
	case 20 * 40 * 30 * 60:
		return 4
	case 10 * 20 * 40 * 60 * 70:
		return 5
	case 10 * 20 * 40 * 50 * 60 * 70:
		return 6
	case 10 * 30 * 60:
		return 7
	case 10 * 20 * 30 * 40 * 50 * 60 * 70:
		return 8
	case 10 * 20 * 30 * 40 * 60 * 70:
		return 9
	case 10 * 20 * 30 * 50 * 60 * 70:
		return 0
	default:
		log.Panicln("Unhandled bar val =>", val, "numCode =>", numCode)
	}

	return 0
}

func decypherNum(serial string, mapping map[int]int32) (prod int) {
	prod = 1

	for _, char := range serial {
		prod = prod * getBarVal(char, mapping)
	}

	return prod
}

func applyMapping(line inputLine, mapping map[int]int32) (number int) {
	numStr := ""

	for _, numCode := range line.outputValues {
		barVal := decypherNum(numCode, mapping)
		digit := getNumberFromBarVal(barVal, numCode)
		numStr += strconv.Itoa(digit)
	}

	res, err := strconv.ParseInt(numStr, 10, 64)
	logErr(err)
	return int(res)
}

func part1(file string, debug bool) (res int) {
	inputLines := parseInput(file)

	if debug {
		log.Println("inputLines =>", inputLines)
	}

	count := 0

	for _, line := range inputLines {
		for _, value := range line.outputValues {
			switch len(value) {
			// 1
			case 2:
				count++
				break
				// 7
			case 3:
				count++
				break
				// 4
			case 4:
				count++
				break
				// 8
			case 7:
				count++
				break
			default:
				break

			}
		}
	}

	return count
}

func part2(file string, debug bool) (sum int) {
	inputLines := parseInput(file)

	if debug {
		log.Println("inputLines =>", inputLines)
	}

	sum = 0

	for i, line := range inputLines {
		mapping := figureOutMapping(line)

		if debug {
			log.Println("Line", i, line)

			for i := 1; i <= len(allLetters); i++ {
				fmt.Printf("Mapping %d => %s\n", i, string(mapping[i]))
			}
		}

		num := applyMapping(line, mapping)
		sum += num

		if debug {
			fmt.Printf("Line %d => %d\n", i, num)
		}
	}

	return sum
}

func main() {
	file := readFile(pathInput)

	//prod1 := part1(file, true)
	//log.Println("Part1 result =>", prod1)

	prod2 := part2(file, true)
	log.Println("Part2 result =>", prod2)
}
