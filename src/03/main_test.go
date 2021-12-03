package main

import (
	"strings"
	"testing"
)

const expect1 = 198
const expect2 = 230

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	prod1 := part1(splice, true)
	if prod1 != expect1 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", prod1, expect1)
	}
}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	prod2 := part2(splice, true)
	if prod2 != expect2 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", prod2, expect2)
	}
}
