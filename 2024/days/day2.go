package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Day2(file *os.File, part int) {
	start := time.Now()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	switch part {
	case 1:
		computeDay2Part1(scanner)
	case 2:
		computeDay2Part2(scanner)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)

}

func computeDay2Part1(scanner *bufio.Scanner) {
	var total int

	re := regexp.MustCompile(`Game (?P<ID>\d+):(?P<Games>.*)`)
	gre := regexp.MustCompile(`(?P<Amount>\d+) (?P<Color>[red|green|blue]+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		var red int
		var green int
		var blue int

		idx := re.SubexpIndex("ID")
		gdx := re.SubexpIndex("Games")

		id, _ := strconv.Atoi(matches[idx])
		colors := matches[gdx]
		colorMatches := gre.FindAllStringSubmatch(colors, -1)

		adx := gre.SubexpIndex("Amount")
		cdx := gre.SubexpIndex("Color")

		for _, match := range colorMatches {
			amount, _ := strconv.Atoi(match[adx])
			color := match[cdx]
			switch color {
			case "red":
				if amount > red {
					red = amount
				}
			case "blue":
				if amount > blue {
					blue = amount
				}
			case "green":
				if amount > green {
					green = amount
				}
			}
		}
		if red < 13 && green < 14 && blue < 15 {
			total += id
		}
	}
	fmt.Println(total)
}

func computeDay2Part2(scanner *bufio.Scanner) {
	var total int

	re := regexp.MustCompile(`Game (?P<ID>\d+):(?P<Games>.*)`)
	gre := regexp.MustCompile(`(?P<Amount>\d+) (?P<Color>[red|green|blue]+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		var red int
		var green int
		var blue int

		gdx := re.SubexpIndex("Games")
		colors := matches[gdx]
		colorMatches := gre.FindAllStringSubmatch(colors, -1)

		adx := gre.SubexpIndex("Amount")
		cdx := gre.SubexpIndex("Color")

		for _, match := range colorMatches {
			amount, _ := strconv.Atoi(match[adx])
			color := match[cdx]
			switch color {
			case "red":
				if amount > red {
					red = amount
				}
			case "blue":
				if amount > blue {
					blue = amount
				}
			case "green":
				if amount > green {
					green = amount
				}
			}

		}
		total += red * green * blue
	}
	fmt.Println(total)
}
