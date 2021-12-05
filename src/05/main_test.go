package main

import (
	"strings"
	"testing"
)

const expect1 = 5
const expect2 = 12

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)
	splice := strings.Split(str, "\n")

	res1 := part1(splice, true)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}

}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)
	splice := strings.Split(str, "\n")

	res2 := part2(splice, true)
	if res2 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}
