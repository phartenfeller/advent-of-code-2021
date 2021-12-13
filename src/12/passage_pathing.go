package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const pathInput = "input.txt"
const pathTestInput = "test.txt"
const pathTestSmallInput = "test_small.txt"
const pathTestMediumInput = "test_medium.txt"

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

type mode int

const (
	mPart1 mode = iota
	mPart2
)

type nodeT struct {
	reachable      []*nodeT
	isSmall, isEnd bool
	id             string
}

var startNode *nodeT
var nodeMap map[string]*nodeT
var resultMap map[string]int

func (node *nodeT) addReachable(n *nodeT) {
	if len(node.reachable) > 0 {
		for _, r := range node.reachable {
			if r.id == n.id {
				return
			}
		}
	}

	node.reachable = append(node.reachable, n)
}

func (node *nodeT) log() {
	reachableStrArr := make([]string, len(node.reachable))

	for i, r := range node.reachable {
		reachableStrArr[i] = r.id
	}

	log.Println("id:", node.id, "isSmall:", node.isSmall, "isEnd:", node.isEnd, "reachable:", reachableStrArr)
}

func nodeExists(key string) bool {
	_, ok := nodeMap[key]
	return ok
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isStart(s string) bool {
	return s == "start"
}

func isEnd(s string) bool {
	return s == "end"
}

func getRoadId(path []string) string {
	return strings.Join(path, ":")
}

func alreadyInPath(path []string, id string) bool {
	for _, p := range path {
		if p == id {
			return true
		}
	}
	return false
}

func hasDuplicate(path []string) bool {
	idMap := make(map[string]bool)

	for _, p := range path {
		// ignore non small ones
		if !nodeMap[p].isSmall {
			continue
		}

		_, ok := idMap[p]

		if ok {
			return true
		} else {
			idMap[p] = true
		}
	}
	return false
}

func smallAllowed(m mode, path []string, id string) bool {
	switch m {
	case mPart1:
		return !alreadyInPath(path, id)
	case mPart2:
		return !alreadyInPath(path, id) || !hasDuplicate(path)
	default:
		log.Panicln("Unknown mode", m)
	}

	return false
}

func parseInput(str string, debug bool) {
	nodeMap = make(map[string]*nodeT)
	resultMap = make(map[string]int)

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		var lastNode *nodeT
		var currNode *nodeT
		for i, node := range strings.Split(line, "-") {
			if !nodeExists(node) {
				currNode = &nodeT{
					reachable: make([]*nodeT, 0),
					isSmall:   !isUpper(node),
					isEnd:     isEnd(node),
					id:        node,
				}

				if isStart(node) {
					startNode = currNode
				}

				if debug {
					log.Println("adding node", node)
				}

				nodeMap[node] = currNode
			} else {
				currNode = nodeMap[node]
			}

			if i > 0 && lastNode != nil {
				lastNode.addReachable(currNode)
				currNode.addReachable(lastNode)
			}

			lastNode = currNode
		}
	}
}

func followPath(m mode, currNode *nodeT, path []string, debug bool) {
	if debug {
		currNode.log()
	}

	currPath := append(path, currNode.id)
	roadId := getRoadId(currPath)

	for _, r := range currNode.reachable {
		if r.id == startNode.id {
			continue
		} else if r.isEnd {
			endPath := append(currPath, r.id)
			endRoadId := getRoadId(endPath)
			if debug {
				log.Println("Found end:", endRoadId)
			}
			resultMap[endRoadId] = 1
		} else if r.isSmall && !smallAllowed(m, currPath, r.id) {
			// don't go back to the same small cave
			if debug {
				log.Println("Already in path:", r.id, roadId)
			}
			continue
		} else {
			if debug {
				log.Println("Following path:", roadId, "->", r.id)
			}
			followPath(m, r, currPath, debug)
		}
	}
}

func countRoutes(m mode, debug bool) int {
	followPath(m, startNode, []string{}, debug)

	if debug {
		i := 0

		for k := range resultMap {
			i++
			log.Println("route", i, "=>", k)
		}
	}

	return len(resultMap)
}

func part1(str string, debug bool) (res int) {
	parseInput(str, debug)

	return countRoutes(mPart1, debug)
}

func part2(str string, debug bool) (res int) {
	parseInput(str, debug)

	return countRoutes(mPart2, debug)
}

func main() {
	str := readFile(pathInput)

	prod1 := part1(str, false)
	log.Println("Part1 result =>", prod1)

	prod2 := part2(str, false)
	log.Println("Part2 result =>", prod2)

}
