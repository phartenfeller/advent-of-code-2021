package main

import (
	"strings"
	"testing"
)

const expect1 = 4512
const expect2 = 1924
const expectBoardCount = 3

//const expect2 = 230

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)
	splice := strings.Split(str, "\n")

	prod1, boardCount1 := part1(splice, true)
	if prod1 != expect1 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", prod1, expect1)
	}

	if expectBoardCount != boardCount1 {
		t.Errorf("Part1 board count was incorrect, got: %d, want: %d.", boardCount1, expectBoardCount)
	}
}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)
	splice := strings.Split(str, "\n")

	prod2, boardCount1 := part2(splice, true)
	if prod2 != expect2 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", prod2, expect1)
	}

	if expectBoardCount != boardCount1 {
		t.Errorf("Part1 board count was incorrect, got: %d, want: %d.", boardCount1, expectBoardCount)
	}
}
