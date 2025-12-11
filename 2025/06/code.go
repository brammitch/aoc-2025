package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return solution(part2, input)
}

type problem struct {
	numbers  []int
	operator string
}

func solution(_ bool, input string) int {
	problems := map[int]problem{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		items := strings.Fields(line)
		for i, item := range items {
			p, ok := problems[i]
			if !ok {
				problems[i] = problem{
					numbers:  make([]int, 0),
					operator: "",
				}
				p = problems[i]
			}
			if number, err := strconv.Atoi(item); err != nil {
				p.operator = item
			} else {
				p.numbers = append(p.numbers, number)
			}

			problems[i] = p
		}
	}

	grandTotal := 0
	for _, problem := range problems {
		switch problem.operator {
		case "*":
			subTotal := 1
			for _, number := range problem.numbers {
				subTotal *= number
			}
			grandTotal += subTotal

		case "+":
			subTotal := 0
			for _, number := range problem.numbers {
				subTotal += number
			}
			grandTotal += subTotal
		}

	}

	return grandTotal
}
