package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	heightmap, err := readHeightmap(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sumOfRiskLevels := calculateSumOfRiskLevels(heightmap)

	fmt.Printf("Sum of risk levels of low points: %d\n", sumOfRiskLevels)
}

type Heightmap [][]int

func (hm Heightmap) getHeightForLowpoints(x int, y int) int {
	if x < 0 || x >= len(hm) {
		return math.MaxInt
	}
	if y < 0 || y >= len(hm[x]) {
		return math.MaxInt
	}
	return hm[x][y]
}

type Point struct {
	x int
	y int
}

func findLowpoints(heightmap Heightmap) []Point {
	var points []Point

	for xIndex, row := range heightmap {
		for yIndex, height := range row {
			if height >= heightmap.getHeightForLowpoints(xIndex-1, yIndex) {
				continue
			}
			if height >= heightmap.getHeightForLowpoints(xIndex+1, yIndex) {
				continue
			}
			if height >= heightmap.getHeightForLowpoints(xIndex, yIndex-1) {
				continue
			}
			if height >= heightmap.getHeightForLowpoints(xIndex, yIndex+1) {
				continue
			}
			points = append(points, Point{x: xIndex, y: yIndex})
		}
	}

	return points
}

func calculateRiskLevel(height int) int {
	return height + 1
}

func calculateSumOfRiskLevels(heightmap Heightmap) int {
	sum := 0

	lowPoints := findLowpoints(heightmap)

	for _, point := range lowPoints {
		sum += calculateRiskLevel(heightmap[point.x][point.y])
	}

	return sum
}

func readHeightmap(input io.Reader) (Heightmap, error) {
	var heightmap Heightmap

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, rune := range line {
			height, err := strconv.Atoi(string(rune))
			if err != nil {
				return nil, err
			}
			row = append(row, height)
		}
		heightmap = append(heightmap, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return heightmap, nil
}
