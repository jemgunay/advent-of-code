package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/day-2/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReports int64

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())

		prevDirection := ""
		prevNum, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			log.Fatalf("failed to parse number: %v", err)
		}

		for i := 1; i < len(nums); i++ {
			nextNum, err := strconv.ParseInt(nums[i], 10, 64)
			if err != nil {
				log.Fatalf("failed to parse number: %v", err)
			}

			diff := nextNum - prevNum
			absDiff := math.Abs(float64(diff))
			if absDiff > 3 || absDiff < 1 {
				// out of range, not safe
				break
			}

			direction := prevDirection
			if diff < 0 {
				direction = "decr"
			} else if diff > 0 {
				direction = "incr"
			}

			if prevDirection != "" && prevDirection != direction {
				// direction has changed, not safe
				break
			}

			if i == len(nums)-1 {
				safeReports++
				break
			}

			prevDirection = direction
			prevNum = nextNum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %v", err)
	}

	fmt.Printf("safe reports: %d\n", safeReports)
}
