package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := readInCSV("input.csv")

	leftOrdered := orderKeys(left)
	rightOrdered := orderKeys(right)

	totalDiff := 0
	for i := 0; i < len(leftOrdered); i++ {
		totalDiff += findDifference(leftOrdered[i], rightOrdered[i])
	}

	log.Println("Total difference is: ", totalDiff)

	part2(left, right)
}

func part2(leftArray []int, rightArray []int) {
	rightFreqMap := make(map[int]int)
	for _, num := range rightArray {
		rightFreqMap[num]++
	}

	total := 0
	for _, num := range leftArray {
		if count, ok := rightFreqMap[num]; ok {
			total += (num * count)
		}
	}

	log.Println("Total: ", total)
}

func findDifference(left int, right int) int {
	return int(math.Abs(float64(left - right)))
}

func orderKeys(a []int) []int {
	sort.Ints(a)
	return a
}

func readInCSV(file string) ([]int, []int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	leftArray := make([]int, 0)
	rightArray := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim whitespace
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		leftAsInt, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatalf("Invalid integer in left part: %v", err)
		}
		rightAsInt, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Fatalf("Invalid integer in right part: %v", err)
		}

		leftArray = append(leftArray, leftAsInt)
		rightArray = append(rightArray, rightAsInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return leftArray, rightArray
}
