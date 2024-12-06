package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input.txt
var puzzle string

type Direction struct {
	x int
	y int
}

type Directions []Direction

type Grid [][][]string

func main() {
	puzzleSlice := strings.Split(strings.TrimSpace(puzzle), "\n")

	wordToFind := "XMAS"
	lenOfWordToFind := len(wordToFind)

	directions := Directions{
		{0, 1},   // Right
		{1, 0},   // Down
		{0, -1},  // Left
		{-1, 0},  // Up
		{1, 1},   // Down Right
		{1, -1},  // Down Left
		{-1, 1},  // Up Right
		{-1, -1}, // Up Left
	}

	res := 0
	for i := 0; i < len(puzzleSlice); i++ {
		for j := 0; j < len(puzzleSlice[i]); j++ {
			if puzzleSlice[i][j] == wordToFind[0] {
				for _, direction := range directions {
					if checkWord(puzzleSlice, i, j, direction, wordToFind, lenOfWordToFind) {
						res++
					}
				}
			}
		}
	}

	dPuzzleSlice := make([][]string, len(puzzleSlice))
	for i, row := range puzzleSlice {
		dPuzzleSlice[i] = strings.Split(row, "")
	}

	countXMas := Part2(dPuzzleSlice)

	log.Print("Number of times the word MAS appears test: ", countXMas)

	log.Print("Number of times the word XMAS appears: ", res)
}

func checkWord(puzzleSlice []string, i, j int, direction Direction, wordToFind string, lenOfWordToFind int) bool {
	for k := 1; k < lenOfWordToFind; k++ {
		newI := i + (k * direction.x) // Move in the direction of the word
		newJ := j + (k * direction.y)

		if newI < 0 || newI >= len(puzzleSlice) || newJ < 0 || newJ >= len(puzzleSlice[newI]) {
			return false
		}

		if puzzleSlice[newI][newJ] != wordToFind[k] {
			return false
		}
	}

	return true
}

func Part2(dSlice [][]string) int {
	res := 0
	grids := gridded(dSlice, 3)

	log.Print(grids)
	for _, grid := range grids {
		res += checkDiagMas(grid)
	}
	return res
}

func checkDiagMas(dSlice [][]string) int {
	middleOfX := dSlice[1][1] == "A"
	lefttop := (dSlice[0][0] == "M" && dSlice[2][2] == "S") || (dSlice[0][0] == "S" && dSlice[2][2] == "M")
	righttop := (dSlice[0][2] == "M" && dSlice[2][0] == "S") || (dSlice[0][2] == "S" && dSlice[2][0] == "M")

	if middleOfX && lefttop && righttop {
		return 1
	}
	return 0
}

func gridded(dSlice [][]string, size int) Grid {
	rows := len(dSlice)
	cols := len(dSlice[0])
	grids := make([][][]string, 0, (rows-size+1)*(cols-size+1))

	for i := 0; i < rows-size+1; i++ {
		for j := 0; j < cols-size+1; j++ {
			gridBlock := make([][]string, size)
			for k := 0; k < size; k++ {
				gridBlock[k] = dSlice[i+k][j : j+size]
			}
			grids = append(grids, gridBlock)
		}
	}
	return grids
}
