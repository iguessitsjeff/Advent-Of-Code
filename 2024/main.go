package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/advent-of-code/2024/days"
)

func main() {

	day := flag.Int("day", 1, "What day to run.")
	part := flag.Int("part", 1, "What part to run.")

	flag.Parse()

	file := readInput(*day)

	switch *day {
	case 1:
		days.Day1(file, *part)
	case 2:
		days.Day2(file, *part)
	}
}

func readInput(day int) *os.File {
	filename := fmt.Sprintf("inputs/input-day-%d.txt", day)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return file
}
