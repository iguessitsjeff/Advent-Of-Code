package days

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Day1(scanner *bufio.Scanner, part int) {

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		content := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(content[0])
		rightNum, _ := strconv.Atoi(content[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var wg sync.WaitGroup
	var ch chan int = make(chan int)
	var total int
	for i := range left {
		wg.Add(1)
		switch part {
		case 1:
			go computeDay1Part1(left[i], right[i], ch, &wg)
		case 2:
			go computeDay1Part2(left[i], right, ch, &wg)
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

func computeDay1Part1(leftNum int, rightNum int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- int(math.Abs(float64(leftNum) - float64(rightNum)))
}

func computeDay1Part2(leftNum int, right []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var idx int = -1

	for i, num := range right {
		if num == leftNum {
			idx = i
			break
		}
	}

	if idx == -1 {
		ch <- 0
		return
	}

	var counter int = 0

	for _, num := range right[idx:] {
		if num == leftNum {
			counter += 1
		} else {
			break
		}
	}

	fmt.Println(leftNum, counter)

	ch <- leftNum * counter

}
