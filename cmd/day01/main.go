package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	measurements, _ := readMeasurements("assets/day01/input")
	increases := countIncreases(measurements)
	fmt.Println(increases)
}

func readMeasurements(path string) ([]int, error) {
	var measurements []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		measurements = append(measurements, int(number))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

