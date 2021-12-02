package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func logErr(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

func strToInt(str string) (num int) {
	num, err := strconv.Atoi(str)
	logErr(err)

	return num
}

func main() {
	increases := 0

	dat, err := os.ReadFile("C:\\Users\\phart\\Documents\\Code\\_random\\advent-of-code-2021\\src\\01\\input.txt")
	logErr(err)

	str := string(dat)

	// fields --> split by whitespace and newline
	splice := strings.Fields(str)

	prev := -1

	for _, line := range splice {
		num := strToInt(line)
		if prev < num && prev != -1 {
			increases++
		} else {
			log.Println("No increase =>", prev, num)
		}

		prev = num
	}

	log.Println("increases =>", increases)

}
