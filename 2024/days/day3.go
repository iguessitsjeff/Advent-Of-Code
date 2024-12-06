package days

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func Day3(scanner *bufio.Scanner, part int) {
	var total int
	var allLines string
	for scanner.Scan() {
		allLines += scanner.Text()
	}
	switch part {
	case 1:
		total += computeDay3Part1(allLines)
	case 2:
		total += computeDay3Part2(allLines)
	}
	fmt.Println(total)

}

func computeDay3Part1(line string) int {
	var total int
	re := regexp.MustCompile(`mul\((?P<A>[0-9]{1,3}),(?P<B>[0-9]{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)
	adx := re.SubexpIndex("A")
	bdx := re.SubexpIndex("B")
	for _, match := range matches {
		a, _ := strconv.Atoi(match[adx])
		b, _ := strconv.Atoi(match[bdx])
		total += a * b
	}
	return total
}

func computeDay3Part2(line string) int {
	remove := regexp.MustCompile(`don't\(\).*?do\(\)`)

	valid := remove.ReplaceAllString(line, "do()")
	return computeDay3Part1(valid)
}
