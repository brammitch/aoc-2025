package main

import (
	"math"
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
		return getJoltage(part2, input)
	}

	return getJoltage(part2, input)
}

func getJoltage(part2 bool, input string) int {
	batteriesToTurnOn := 2
	if part2 {
		batteriesToTurnOn = 12
	}

	total := 0

	banks := strings.Split(input, "\n")
	for _, bank := range banks {
		// Create an array to store our selected batteries
		nums := make([]int, batteriesToTurnOn)

		for i, char := range bank {
			// Iterate over each battery in the bank, getting the "joltage" as an int
			num, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			for j := 0; j < batteriesToTurnOn; j++ {
				// Iterate over each existing battery and determine if a larger battery can be turned on instead.
				// To meet this criteria, there must be sufficient space let in the remainder of the bank to satisfy
				// the required number of batteries that must be turned on.
				if num > nums[j] && i-j <= len(bank)-batteriesToTurnOn {
					nums[j] = num

					// Since a new highest number has been found, reset all remaining batteries
					for k := j + 1; k < len(nums); k++ {
						nums[k] = 0
					}

					break
				}
			}
		}

		for i, num := range nums {
			// Turn nums into a regular number, e.g, []int{ 9, 8 } becomes 98, and add them to our total
			total += num * int(math.Pow10(len(nums)-1-i))
		}
	}

	return total
}
