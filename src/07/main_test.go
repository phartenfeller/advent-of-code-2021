package main

import (
	"testing"
)

const expect1 = 37
const expect2 = 26984457539

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part1(str, true)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}

}

/*
func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)

	res2 := part2(str, true)
	if res2 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}

*/
