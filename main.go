package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var height, width, steps int

// print matrix to output. live cells are green ones, dead cells are red zeros
func printMatrix(m [][]bool) {
	for y := range m {
		for x := range m[y] {
			if m[y][x] {
				fmt.Print("\033[32m1\033[0m ")
			} else {
				fmt.Print("\033[31m0\033[0m ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// count neighbors that are alive
func countNeighbors(m [][]bool, y, x int) int {
	count := 0
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, neighbor := range neighbors {
		ny, nx := y+neighbor[0], x+neighbor[1]
		if ny >= 0 && ny < height && nx >= 0 && nx < width && m[ny][nx] {
			count++
		}
	}

	return count
}

// return next step in the game of life
func nextStep(m [][]bool) [][]bool {
	next := make([][]bool, height)
	for i := range next {
		next[i] = make([]bool, width)
	}
	for y := range m {
		for x := range m[y] {
			neighbors := countNeighbors(m, y, x)
			if m[y][x] { //alive
				if neighbors < 2 || neighbors > 3 {
					next[y][x] = false // dies if underpopulation or overcrowded
				} else {
					next[y][x] = true // still alive
				}
			} else { // dead
				if neighbors == 3 {
					next[y][x] = true // became live if 3 neighbors are alive
				}
			}
		}
	}
	return next
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <height> <width> <steps>")
		return
	}
	var err error
	height, err = strconv.Atoi(os.Args[1])
	if err != nil || height <= 0 {
		fmt.Println("Invalid height")
		return
	}
	width, err = strconv.Atoi(os.Args[2])
	if err != nil || width <= 0 {
		fmt.Println("Invalid width")
		return
	}
	steps, err = strconv.Atoi(os.Args[3])
	if err != nil || steps <= 0 {
		fmt.Println("Invalid steps")
		return
	}

	matrix := make([][]bool, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}

	rand.Seed(time.Now().UnixNano())
	for y := range matrix {
		for x := range matrix[y] {
			matrix[y][x] = rand.Intn(2) != 0
		}
	}
	for i := 0; i < steps; i++ {
		fmt.Printf("\033[s") // Save the cursor position
		fmt.Printf("\nStep %d:\n", i)
		printMatrix(matrix)
		if i != steps-1 {
			fmt.Printf("\033[u\033[K") // Restore the cursor position and clear the line
		}
		matrix = nextStep(matrix)
		time.Sleep(1 * time.Second) // Sleep for 5 seconds between steps
	}
}
