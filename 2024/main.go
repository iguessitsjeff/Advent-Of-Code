package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/advent-of-code/2024/days"
)

func main() {
	start := time.Now()
	day := flag.Int("day", 1, "What day to run.")
	part := flag.Int("part", 1, "What part to run.")

	flag.Parse()

	file := readInput(*day)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	switch *day {
	case 1:
		days.Day1(scanner, *part)
	case 2:
		days.Day2(scanner, *part)
	case 3:
		days.Day3(scanner, *part)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
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
