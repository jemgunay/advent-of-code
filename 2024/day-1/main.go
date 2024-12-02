package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/day-1/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const fileLines = 1000
	var (
		leftList        = make([]int64, 0, fileLines)
		rightList       = make([]int64, 0, fileLines)
		rightListCounts = make(map[int64]int64)
	)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		if len(nums) != 2 {
			log.Fatalf("invalid line input: %v", nums)
		}

		left, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			log.Fatalf("failed to parse left number: %v", err)
		}
		leftList = append(leftList, left)

		right, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			log.Fatalf("failed to parse right number: %v", err)
		}
		rightList = append(rightList, right)
		rightListCounts[right]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %v", err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var sum, similarityScore int64
	for i := 0; i < len(leftList); i++ {
		left, right := leftList[i], rightList[i]
		sum += int64(math.Abs(float64(left - right)))

		similarityScore += left * rightListCounts[left]
	}

	fmt.Printf("sum: %v\n", sum)                    // sum: 1580061
	fmt.Printf("similarity: %d\n", similarityScore) // similarity: 23046913
}
