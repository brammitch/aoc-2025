package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gbatagian/deepsort"
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
	if part2 {
		return freshIngredientChecker(part2, input)
	}

	return freshIngredientChecker(part2, input)
}

func freshIngredientChecker(part2 bool, input string) int {
	ids := strings.Split(input, "\n")
	freshIngredientsRanges := make([][]int, 0)
	idsToCheck := make([]int, 0)
	freshIngredientCount := 0

	for _, id := range ids {
		if id == "" {
			continue
		}

		idNum, err := strconv.Atoi(id)

		if err != nil {
			if errors.Is(err, strconv.ErrSyntax) {
				// Id is a range, e.g, 3-5
				idRange := strings.Split(id, "-")
				if len(idRange) == 2 {
					lowerBounds, err := strconv.Atoi(idRange[0])
					if err != nil {
						continue
					}

					upperBounds, err := strconv.Atoi(idRange[1])
					if err != nil {
						continue
					}

					freshIngredientsRanges = append(freshIngredientsRanges, []int{lowerBounds, upperBounds})
				}

			} else {
				continue
			}
		} else {
			idsToCheck = append(idsToCheck, idNum)
		}
	}

	if part2 {
		deepsort.DeepSort(&freshIngredientsRanges, []int{0, 1})

		skip := false
		low := 0
		high := 0

		for i, r := range freshIngredientsRanges {
			if !skip {
				// Set the lower bounds if not combining ranges
				low = r[0]
			}

			// Keep raising the limit as long as ranges overlap
			if r[1] >= high {
				high = r[1]
			}

			if i+1 <= len(freshIngredientsRanges)-1 && high >= freshIngredientsRanges[i+1][0] {
				// If overlap, do not calculate yet; skip to the next iteration
				skip = true
				continue
			}

			// No more overlap, so calculate the range and reset for the next iteration
			freshIngredientCount += (high - low) + 1

			low = 0
			high = 0
			skip = false
		}
	} else {
		// Part 1
		for _, id := range idsToCheck {
			for _, r := range freshIngredientsRanges {
				if id >= r[0] && id <= r[1] {
					freshIngredientCount++
					break
				}
			}
		}
	}

	return freshIngredientCount
}
