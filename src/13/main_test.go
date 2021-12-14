package main

import (
	"testing"
)

const expect1 = 17
const expect2 = 16

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part1(str, true)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}

}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)

	part2(str, true)
	res2 := gBoard.getCount()
	if res2 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}

func TestPart3(t *testing.T) {
	str := readFile("test2.txt")

	part2(str, true)
	res2 := gBoard.getCount()
	if res2 != 7 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}
