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
		zeroDialCount, _ := spinTheDial(part2, input)
		return zeroDialCount
	}
	// solve part 1 here
	zeroDialCount, _ := spinTheDial(part2, input)
	return zeroDialCount

}

func spinTheDial(part2 bool, input string) (int, error) {
	dial := 50
	zeroDialCount := 0

	rotations := strings.Split(input, "\n") // Expect slice of strings formatted like "L48", "R5", etc.

	for _, rotation := range rotations {
		// For part 2
		startingDial := dial // Keep track of the dial's starting point at the beginning of each rotation
		spunAround := false  // Determine if the dial passed zero in either direction 99 -> 0, or 99 <- 0

		direction := string(rune(rotation[0]))      // "L" or "R"
		distance, err := strconv.Atoi(rotation[1:]) // e.g, 48, 5
		if err != nil {
			return 0, err
		}

		if part2 {
			// Add the number of times the dial would've touched the zero point
			zeroDialCount += distance / 100
		}

		if direction == "L" {
			// Subtract when moving to the left
			dial -= distance % 100
		} else {
			// Add when moving to the right
			dial += distance % 100
		}

		if dial > 99 {
			// If the dial went over 99, it needs to start from zero again
			dial -= 100
			spunAround = true
		} else if dial < 0 {
			// If the dial went below zero, it needs to start from 99 again
			dial += 100
			spunAround = true
		}

		if dial == 0 {
			zeroDialCount++
		} else if part2 && spunAround && startingDial != 0 {
			// If the dial spun around and wasn't already at zero when it started (and therefore was counted last rotation),
			// we need to add +1 to count this rotation for part 2.
			zeroDialCount++
		}
		// fmt.Printf("The dial is rotated %s%d to point at %d.\n", direction, distance, dial)
	}

	return zeroDialCount, nil
}
