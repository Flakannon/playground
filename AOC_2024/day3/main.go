package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
)

//go:embed input.txt
var instructions string

func main() {
	regexMatch := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`

	matches := regexFindAll(regexMatch, instructions)

	sum := 0
	enabledInstructions := true
	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabledInstructions = true
		case "don't()":
			enabledInstructions = false
		default:
			if enabledInstructions {
				sum += toInt(match[1]) * toInt(match[2])
			}
		}
	}

	log.Println("Sum of all matches: ", sum)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert %s to int: %v", s, err)
	}

	return i
}

func regexFindAll(regexMatch, instructions string) [][]string {
	matches := regexp.MustCompile(regexMatch).FindAllStringSubmatch(instructions, -1)

	return matches
}
