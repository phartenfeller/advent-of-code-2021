package main

import (
	"testing"
)

const expect1 = 26397
const expect2 = 288957

func TestValidChunks(t *testing.T) {
	res := 0

	str := "([])"
	res = checkLine(str, false)
	if res != 0 {
		t.Errorf("%s should be valid, but is not => %d.", str, res)
	}

	str = "{()()()}"
	res = checkLine(str, false)
	if res != 0 {
		t.Errorf("%s should be valid, but is not => %d.", str, res)
	}

	str = "<([{}])>"
	res = checkLine(str, false)
	if res != 0 {
		t.Errorf("%s should be valid, but is not => %d.", str, res)
	}

	str = "<([{}])>"
	res = checkLine(str, false)
	if res != 0 {
		t.Errorf("%s should be valid, but is not => %d.", str, res)
	}

	str = "(((((((((())))))))))"
	res = checkLine(str, false)
	if res != 0 {
		t.Errorf("%s should be valid, but is not => %d.", str, res)
	}
}

func TestInvalidChunks(t *testing.T) {
	res := 0

	str := "(]"
	res = checkLine(str, false)
	if res == 0 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "{()()()>"
	res = checkLine(str, false)
	if res == 0 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "(((()))}"
	res = checkLine(str, false)
	if res == 0 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "<([]){()}[{}])"
	res = checkLine(str, false)
	if res == 0 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}
}

func TestInvalidScores(t *testing.T) {
	res := 0

	str := "{([(<{}[<>[]}>{[]{[(<()>"
	res = checkLine(str, false)
	if res != 1197 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "[[<[([]))<([[{}[[()]]]"
	res = checkLine(str, false)
	if res != 3 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "[{[{({}]{}}([{[{{{}}([]"
	res = checkLine(str, false)
	if res != 57 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "[<(<(<(<{}))><([]([]()"
	res = checkLine(str, false)
	if res != 3 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}

	str = "<{([([[(<>()){}]>(<<{{"
	res = checkLine(str, false)
	if res != 25137 {
		t.Errorf("%s should be invalid, but is ok => %d.", str, res)
	}
}

func TestPart1(t *testing.T) {
	str := readFile(pathTestInput)

	res1 := part1(str, true)
	if res1 != expect1 {
		t.Errorf("Part1 res was incorrect, got: %d, want: %d.", res1, expect1)
	}
}

func TestPart2(t *testing.T) {
	str := readFile(pathTestInput)

	res2 := part2(str, true)
	if res2 != expect2 {
		t.Errorf("Part2 res was incorrect, got: %d, want: %d.", res2, expect2)
	}
}
