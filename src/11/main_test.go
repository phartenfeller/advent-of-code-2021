package main

import "testing"

const pathTestInputSmall = "small_test.txt"

const expectSmall = 9
const expect1 = 1656
const expect2 = 195

func TestPart1Small(t *testing.T) {
	str := readFile(pathTestInputSmall)

	res1 := part1(str, true, 1)
	if res1 != expectSmall {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expectSmall)
	}

	gBoard.print()

}

func TestPart1Small2(t *testing.T) {
	str := readFile(pathTestInputSmall)

	res1 := part1(str, true, 2)
	if res1 != expectSmall {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expectSmall)
	}

	gBoard.print()

}

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part1(str, true, part1Flashes)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}

	gBoard.print()

}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)

	res2 := part2(str, true)
	if res2 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}
