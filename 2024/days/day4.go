package days

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

var Day4Matrix [][]string

func Day4(scanner *bufio.Scanner, part int) {
	var total int
	for scanner.Scan() {
		Day4Matrix = append(Day4Matrix, strings.Split(scanner.Text(), ""))
	}

	switch part {
	case 1:
		total = computeDay4Part1()
	case 2:
		total = computeDay4Part2()
	}

	fmt.Println(total)

}

func computeDay4Part1() int {
	var wg sync.WaitGroup
	var ch chan int = make(chan int)
	var total int

	for i := range Day4Matrix {
		for j := range Day4Matrix[i] {
			if Day4Matrix[i][j] == "X" {
				wg.Add(8)
				go E(i, j, ch, &wg)
				go W(i, j, ch, &wg)
				go N(i, j, ch, &wg)
				go S(i, j, ch, &wg)
				go NE(i, j, ch, &wg)
				go NW(i, j, ch, &wg)
				go SE(i, j, ch, &wg)
				go SW(i, j, ch, &wg)
			}
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		total += num
	}
	return total
}

func E(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j+3+1 <= len(Day4Matrix[i]) {
		m := Day4Matrix[i][j+1]
		a := Day4Matrix[i][j+2]
		s := Day4Matrix[i][j+3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func W(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j-3 >= 0 {
		m := Day4Matrix[i][j-1]
		a := Day4Matrix[i][j-2]
		s := Day4Matrix[i][j-3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func N(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if i-3 >= 0 {
		m := Day4Matrix[i-1][j]
		a := Day4Matrix[i-2][j]
		s := Day4Matrix[i-3][j]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func S(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if i+3+1 <= len(Day4Matrix) {
		m := Day4Matrix[i+1][j]
		a := Day4Matrix[i+2][j]
		s := Day4Matrix[i+3][j]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func NE(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j+3+1 <= len(Day4Matrix[i]) && i-3 >= 0 {
		m := Day4Matrix[i-1][j+1]
		a := Day4Matrix[i-2][j+2]
		s := Day4Matrix[i-3][j+3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func NW(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j-3 >= 0 && i-3 >= 0 {
		m := Day4Matrix[i-1][j-1]
		a := Day4Matrix[i-2][j-2]
		s := Day4Matrix[i-3][j-3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func SE(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j+3+1 <= len(Day4Matrix[i]) && i+3+1 <= len(Day4Matrix) {
		m := Day4Matrix[i+1][j+1]
		a := Day4Matrix[i+2][j+2]
		s := Day4Matrix[i+3][j+3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func SW(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j-3 >= 0 && i+3+1 <= len(Day4Matrix) {
		m := Day4Matrix[i+1][j-1]
		a := Day4Matrix[i+2][j-2]
		s := Day4Matrix[i+3][j-3]

		if m == "M" && a == "A" && s == "S" {
			found = 1
		}
	}

	ch <- found
}

func computeDay4Part2() int {
	var wg sync.WaitGroup
	var ch chan int = make(chan int)
	var total int

	for i := range Day4Matrix {
		for j := range Day4Matrix[i] {
			if Day4Matrix[i][j] == "A" {
				wg.Add(1)
				go checkCross(i, j, ch, &wg)
			}
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		total += num
	}
	return total
}

func checkCross(i int, j int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var found int

	if j-1 >= 0 && i-1 >= 0 && i+2 <= len(Day4Matrix) && j+2 <= len(Day4Matrix[i]) {
		tl := Day4Matrix[i-1][j-1]
		br := Day4Matrix[i+1][j+1]
		bl := Day4Matrix[i+1][j-1]
		tr := Day4Matrix[i-1][j+1]

		if ((tl == "M" || tl == "S") && (br == "M" || br == "S") && (tl != br)) && ((bl == "M" || bl == "S") && (tr == "M" || tr == "S") && (bl != tr)) {
			found = 1
		}

	}

	ch <- found
}
