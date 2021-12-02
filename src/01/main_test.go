package main

import (
	"strings"
	"testing"
)

const expect1 = 7
const expect2 = 5

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	largerMesurements1 := part1(splice, false)
	if largerMesurements1 != expect1 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", largerMesurements1, expect1)
	}
}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)
	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	largerMesurements2 := part2(splice, false)
	if largerMesurements2 != expect2 {
		t.Errorf("Part1 product was incorrect, got: %d, want: %d.", largerMesurements2, expect2)
	}
}
