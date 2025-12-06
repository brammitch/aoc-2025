package main

import (
	"slices"
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
		getRollsOfPaper(part2, input)
	}
	// solve part 1 here
	return getRollsOfPaper(part2, input)
}

func getRollsOfPaper(part2 bool, input string) int {

	rows := strings.Split(input, "\n")
	grid := make([][]string, len(rows))
	newGrid := make([][]string, len(rows))

	for y, row := range rows {
		grid[y] = make([]string, len(row))
		newGrid[y] = make([]string, len(row))

		for x, char := range row {
			grid[y][x] = string(char)
		}
	}

	return removeRollsOfPaper(part2, grid, newGrid, 0)
}

func removeRollsOfPaper(part2 bool, grid, newGrid [][]string, rolls int) int {
	startingRollCount := rolls
	endingRollCount := startingRollCount

	for y, row := range grid {
		yMax := len(grid) - 1

		for x, char := range row {
			newGrid[y][x] = grid[y][x]
			if string(char) != "@" {
				continue
			}

			xMax := len(row) - 1
			adjacentRolls := 0

			if y-1 >= 0 {
				if grid[y-1][x] == "@" {
					adjacentRolls++
				}
				if x-1 >= 0 {
					if grid[y-1][x-1] == "@" {
						adjacentRolls++
					}
				}
				if x+1 <= xMax {
					if grid[y-1][x+1] == "@" {
						adjacentRolls++
					}
				}
			}

			if x+1 <= xMax {
				if grid[y][x+1] == "@" {
					adjacentRolls++
				}
			}

			if x-1 >= 0 {
				if grid[y][x-1] == "@" {
					adjacentRolls++
				}
			}

			if y+1 <= yMax {
				if x+1 <= xMax {
					if grid[y+1][x+1] == "@" {
						adjacentRolls++
					}
				}
				if grid[y+1][x] == "@" {
					adjacentRolls++
				}
				if x-1 >= 0 {
					if grid[y+1][x-1] == "@" {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				newGrid[y][x] = "x"
				endingRollCount++
			}
		}
	}

	if part2 && endingRollCount > startingRollCount {
		return removeRollsOfPaper(part2, newGrid, slices.Clone(newGrid), endingRollCount)
	}

	return endingRollCount
}
