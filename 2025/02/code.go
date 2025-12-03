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
	sum, _ := sumInvalidIds(part2, input)
	return sum
}

func sumInvalidIds(_ bool, input string) (int, error) {
	sumOfInvalidProductIds := 0
	productIdRanges := strings.Split(input, ",")

	for _, productIdRange := range productIdRanges {
		bounds := strings.Split(productIdRange, "-")
		lowerBound, err := strconv.Atoi(bounds[0])
		if err != nil {
			continue
		}
		upperBound, err := strconv.Atoi(bounds[1])
		if err != nil {
			continue
		}

		for i := lowerBound; i <= upperBound; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 == 0 {
				mid := len(str) / 2
				oneHalf, err := strconv.Atoi(str[:mid])
				if err != nil {
					continue
				}
				twoHalf, err := strconv.Atoi(str[mid:])
				if err != nil {
					continue
				}

				if oneHalf == twoHalf {
					sumOfInvalidProductIds += i
				}
			}
		}
	}

	return sumOfInvalidProductIds, nil
}
