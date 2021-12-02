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
	file, err := os.Open("assets/day01/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	measurements, err := readMeasurements(file)
	if err != nil {
		log.Fatal(err)
	}

	increases := countIncreases(measurements)

	fmt.Printf("Number of increases:\t%d\n", increases)
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
		if measurements[index + 1] > measurements[index] {
			count++
		}
	}

	return count
}

