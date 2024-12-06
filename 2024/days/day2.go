package days

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func Day2(scanner *bufio.Scanner, part int) {
	var wg sync.WaitGroup
	var ch chan int = make(chan int)
	var total int

	for scanner.Scan() {
		line := scanner.Text()
		switch part {
		case 1:
			wg.Add(1)
			go computeDay2Part1(line, ch, &wg)
		case 2:
			wg.Add(1)
			go computeDay2Part2(line, ch, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		total += num
	}

	fmt.Println(total)

}

func computeDay2Part1(line string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var safe = 1
	var increasing = false
	var decreasing = false
	vals := strings.Split(line, " ")
	prev, _ := strconv.Atoi(vals[0])

	for _, val := range vals[1:] {
		curr, _ := strconv.Atoi(val)
		change := prev - curr

		if !decreasing && change < 0 {
			decreasing = true
		}
		if !increasing && change > 0 {
			increasing = true
		}

		change = int(math.Abs(float64(change)))

		if change > 3 || change < 1 {
			safe = 0
			break
		}

		if increasing && decreasing {
			safe = 0
			break
		}
		prev = curr
	}

	ch <- safe

}

func computeDay2Part2(line string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var safe = 1

	var skip = false

	var increasing = false
	var decreasing = false

	vals := strings.Split(line, " ")
	prev, _ := strconv.Atoi(vals[0])

	for _, val := range vals[1:] {
		curr, _ := strconv.Atoi(val)
		distance := prev - curr

		if !decreasing && distance < 0 {
			decreasing = true
		}
		if !increasing && distance > 0 {
			increasing = true
		}

		change := int(math.Abs(float64(distance)))

		if change > 3 || change < 1 {
			if !skip {
				skip = true
				continue
			}

			safe = 0
			break
		}

		if increasing && decreasing {
			if !skip {

				skip = true
			}

			safe = 0
			break
		}
		prev = curr
	}

	ch <- safe

}
