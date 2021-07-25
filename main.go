package main

import (
	"flag"
	"fmt"
	"life/backend"
	"life/frontend"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var size = flag.Int("size", 50, "size of the matrix")
var speed = flag.Int("speed", 100, "number of milliseconds in between each refresh")

func main() {
	flag.Parse()
	iterations := 0

	// handle early termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("toroidal game of life ended in", iterations, "iterations")
		os.Exit(0)
	}()

	rand.Seed(time.Now().Unix())

	matrix := backend.MakeMatrix(*size)
	previousState := backend.CopyMatrix(matrix)

	matrix = backend.RandomState(matrix)

	for !backend.SameMatrix(matrix, previousState) {
		frontend.Refresh(time.Duration(*speed) * time.Millisecond)
		frontend.PrettyPrintMatrix(matrix)
		previousState = backend.CopyMatrix(matrix)
		matrix = backend.UpdateMatrix(matrix)
		iterations++
	}

	fmt.Println("toroidal game of life stalled in", iterations, "iterations")

}
