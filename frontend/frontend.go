// +build !windows

// The frontend package handles the display utilities for the game of life.
package frontend

import (
	"fmt"
	"time"
)

// Refresh waits a bit then clears the screen and moves the cursor to the top left.
func Refresh(waitTime time.Duration) {
	// wait a bit
	time.Sleep(waitTime)
	// clear screen - this bit only works on *nix systems
	fmt.Print("\033[2J")
	// move cursor top left - this bit only works on *nix systems
	fmt.Print("\033[H")
}

// PrettyPrintMatrix pretty prints `matrix` to the screen using `sym` to indicate a live cell.
func PrettyPrintMatrix(matrix [][]int) {
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				fmt.Print("■", " ")
			} else {
				fmt.Print(" ", " ")
			}

		}
		fmt.Println()
	}
}

// PrintMatrix ugly prints `matrix` to the screen.
func PrintMatrix(matrix [][]int) {
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Print(matrix[row][col], " ")
		}
		fmt.Println()
	}
}
