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
	if part2 {
		sum, _ := sumInvalidIds(part2, input)
		return sum
	}
	// solve part 1 here
	sum, _ := sumInvalidIds(part2, input)
	return sum
}

func sumInvalidIds(part2 bool, input string) (int, error) {
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

			if part2 {
				intSet := make(map[int]bool)

				for j := 1; j <= len(str); j++ {
					// Check for each way the string can be evenly split
					if len(str)%j == 0 {
						chunks := getStringChunks(str, j)
						if len(chunks) > 1 {
							chunksMatch := true

							// Check if all chunks of the string match
							for k := 0; k < len(chunks); k++ {
								if k+1 < len(chunks) && chunks[k] != chunks[k+1] {
									chunksMatch = false
									break
								}
							}

							if chunksMatch {
								value, err := strconv.Atoi(str)

								if err != nil || intSet[value] {
									continue
								}

								// Track the value to ensure we don't add duplicates (e.g., "222222" could be split different ways and would match)
								intSet[value] = true
								sumOfInvalidProductIds += value
							}
						}
					}
				}
			} else {
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
	}

	return sumOfInvalidProductIds, nil
}

func getStringChunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}
