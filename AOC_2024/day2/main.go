package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type level []int

type levels []level

func main() {
	lvls := readIn("input.csv")

	safeCount := 0

	unsafeLvls := make([]level, 0)
	for _, lvl := range lvls {
		if isSafe(lvl) {
			safeCount++
		} else {
			unsafeLvls = append(unsafeLvls, lvl)
		}
	}

	log.Printf("Safe levels part 1: %d", safeCount)

	count := part2(unsafeLvls)

	log.Printf("Safe levels part 2 : %d", count+safeCount)
}

func part2(lvls levels) int {
	safeCount := 0
	for _, lvl := range lvls {
		for i := 0; i < len(lvl); i++ {
			isSafe := isSafeWithDelete(lvl, i)
			if isSafe {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isSafeWithDelete(lvl level, toDelete int) bool {
	temp := make([]int, len(lvl))
	copy(temp, lvl)

	if toDelete == len(temp)-1 {
		temp = temp[:toDelete]
	} else {
		temp = append(temp[:toDelete], temp[toDelete+1:]...)
	}
	return isSafe(temp)
}

func isSafe(lvl level) bool {
	if len(lvl) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(lvl); i++ {
		diff := math.Abs(float64(lvl[i] - lvl[i-1]))
		if diff < float64(1) || diff > float64(3) {
			return false
		}
		if lvl[i] > lvl[i-1] {
			decreasing = false
		} else if lvl[i] < lvl[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func readIn(file string) levels {
	var res levels
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			log.Fatalf("Invalid line format: %s", scanner.Text())
		}

		var lvl level
		for _, part := range parts {
			asNumb, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalf("Invalid integer in line: %v", err)
			}
			lvl = append(lvl, asNumb)
		}

		res = append(res, lvl)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return res
}
