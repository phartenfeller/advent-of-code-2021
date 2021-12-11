package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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

type tokenPair struct {
	start      rune
	end        rune
	pts        int
	autoComPts int
}

var tokenPairs []tokenPair = []tokenPair{
	{'(', ')', 3, 1},
	{'[', ']', 57, 2},
	{'{', '}', 1197, 3},
	{'<', '>', 25137, 4},
}

type intStack []int32

func (s *intStack) push(i int32) {
	*s = append(*s, i)
}

func (s *intStack) pop() int32 {
	if len(*s) == 0 {
		return 0
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func checkLine(line string, debug bool) int {
	var lineStack intStack

	if debug {
		log.Println("Checking line =>", line)
	}

	// validate tokens
	for i, char := range line {
		found := false

		// check every char for a start token match
		for _, pair := range tokenPairs {
			if char == int32(pair.start) {
				//if debug {
				//	fmt.Printf("Added %c to stack\n", char)
				//}
				lineStack.push(int32(pair.end))
				found = true
			}
		}

		if found {
			continue
		}

		// check if this is the end of a token
		popped := lineStack.pop()
		if char == popped {
			//if debug {
			//	fmt.Printf("Matched %c at pos %d\n", char, i)
			//}
		} else {
			for _, pair := range tokenPairs {
				if char == int32(pair.end) {
					if debug {
						fmt.Printf("- No match for %c at pos %d. Received %c. --> %d pts\n", popped, i, char, pair.pts)
					}
					return pair.pts
				}
			}
		}
	}

	return 0
}

func scoreInvalidLines(lines []string, debug bool) int {
	points := 0

	for _, line := range lines {
		points += checkLine(line, debug)
	}

	return points
}

func part1(str string, debug bool) (res int) {
	lines := strings.Split(str, "\n")
	res = scoreInvalidLines(lines, debug)

	return res
}

func autocompleteLine(line string, debug bool) int {
	points := 0

	if debug {
		log.Println("Autocompleting line =>", line)
	}

	var lineStack intStack

	// validate tokens
	for _, char := range line {
		found := false

		// check every char for a start token match
		for _, pair := range tokenPairs {
			if char == int32(pair.start) {
				//if debug {
				//	fmt.Printf("Added %c to stack\n", char)
				//}
				lineStack.push(int32(pair.end))
				found = true
			}
		}

		if found {
			continue
		} else {
			// must be an end tag -> empty next stack entry
			lineStack.pop()
		}
	}

	stackLen := len(lineStack)
	for i := 0; i < stackLen; i++ {
		popped := lineStack.pop()

		if debug {
			fmt.Printf("- Popped %c at pos %d\n", popped, i)
		}

		for _, pair := range tokenPairs {
			if popped == int32(pair.end) {
				if debug {
					fmt.Printf("- Add %d pts to curren points %d\n", pair.autoComPts, points)
				}

				points = points * 5
				points += pair.autoComPts
			}
		}
	}
	return points
}

func fixLines(lines []string, debug bool) int {
	var points []int

	for _, line := range lines {
		if debug {
			log.Println("Checking line =>", line)
		}

		lineOk := checkLine(line, debug) == 0

		if lineOk {
			points = append(points, autocompleteLine(line, debug))
		}
	}

	sort.Slice(points, func(a, b int) bool {
		return points[a] < points[b]
	})

	if debug {
		fmt.Println("Points:", points)
	}

	if len(points)%2 != 1 {
		log.Panicln("Even number of points", len(points))
	}

	middle := len(points) / 2

	return points[middle]
}

func part2(str string, debug bool) (res int) {
	lines := strings.Split(str, "\n")
	res = fixLines(lines, debug)

	return res
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, true)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, true)
	log.Println("Part2 result =>", prod2)
}
