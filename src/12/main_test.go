package main

import (
	"log"
	"testing"
)

const expect1Sm = 10
const expect1Md = 19
const expect1 = 226
const expect2Sm = 36
const expect2Md = 103
const expect2 = 3509

func TestPart1Small(t *testing.T) {
	str := readFile(pathTestSmallInput)

	res1 := part1(str, false)
	if res1 != expect1Sm {
		t.Errorf("Part1Small res was incorrect, got: %d, want: %d.", res1, expect1Sm)
	}

}

func TestPart1Medium(t *testing.T) {
	str := readFile(pathTestMediumInput)

	res1 := part1(str, false)
	log.Println(res1, expect1Md)
	if res1 != expect1Md {
		t.Errorf("Part1Medium res was incorrect, got: %d, want: %d.", res1, expect1Md)
	}

}

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part1(str, false)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}

}

func TestPart2Small(t *testing.T) {
	str := readFile(pathTestSmallInput)

	res2 := part2(str, false)
	if res2 != expect2Sm {
		t.Errorf("Part2Sm res was incorrect, got: %d, want: %d.", res2, expect2Sm)
	}
}

func TestPart2Med(t *testing.T) {
	str := readFile(pathTestMediumInput)

	res2 := part2(str, false)
	if res2 != expect2Md {
		t.Errorf("Part2Md res was incorrect, got: %d, want: %d.", res2, expect2Md)
	}
}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part2(str, false)
	if res1 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res1, expect2)
	}

}
