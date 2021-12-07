package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	positions, err := readCrabPositions(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	minimumCost := findMinimumCost(positions, calculateCost)
	correctMinimumCosts := findMinimumCost(positions, calculateCostsCorrectly)

	fmt.Printf("Minimum costs are %d\n", minimumCost)
	fmt.Printf("Correct minimum costs are %d\n", correctMinimumCosts)
}

func findMinimumCost(positions []int, calculation func([]int, int) int) int {
	minimumCost := math.MaxInt

	min := getMinimum(positions)
	max := getMaximum(positions)

	for target := min; target <= max; target++ {
		cost := calculation(positions, target)
		if cost < minimumCost {
			minimumCost = cost
		}
	}

	return minimumCost
}

func calculateCost(positions []int, targetPosition int) int {
	cost := 0

	for _, position := range positions {
		cost += int(math.Abs(float64(targetPosition - position)))
	}

	return cost
}

func calculateCostsCorrectly(positions []int, targetPosition int) int {
	cost := 0

	for _, position := range positions {
		difference := int(math.Abs(float64(targetPosition - position)))
		// n * (n+1) / 2 is the sum of 1..n
		cost += (difference * (difference + 1)) / 2
	}

	return cost
}

func getMinimum(positions []int) int {
	min := math.MaxInt
	for _, p := range positions {
		if p < min {
			min = p
		}
	}
	return min
}

func getMaximum(positions []int) int {
	max := 0
	for _, p := range positions {
		if p > max {
			max = p
		}
	}
	return max
}

func readCrabPositions(input io.Reader) ([]int, error) {
	var positions []int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		for _, position := range strings.Split(line, ",") {
			p, err := strconv.Atoi(position)
			if err != nil {
				return nil, err
			}
			positions = append(positions, p)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
