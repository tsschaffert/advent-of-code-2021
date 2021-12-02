package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	measurements, err := readMeasurements(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	increasesSimple := countIncreases(measurements)

	slidingWindows := generateSlidingWindows(measurements)
	increasesSlidingWindow := countIncreases(slidingWindows)

	fmt.Printf("Number of increases (noisy):\t%d\n", increasesSimple)
	fmt.Printf("Number of increases (noise reduced):\t%d\n", increasesSlidingWindow)
}

func readMeasurements(input io.Reader) ([]int, error) {
	var measurements []int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, int(number))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return measurements, nil
}

func countIncreases(measurements []int) int {
	count := 0
	for index, _ := range measurements[1:] {
		if measurements[index+1] > measurements[index] {
			count++
		}
	}

	return count
}

func generateSlidingWindows(measurements []int) []int {
	var slidingWindows []int

	for index, _ := range measurements[2:] {
		slidingWindow := measurements[index] + measurements[index+1] + measurements[index+2]
		slidingWindows = append(slidingWindows, slidingWindow)
	}

	return slidingWindows
}
