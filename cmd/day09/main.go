package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	heightmap, err := readHeightmap(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sumOfRiskLevels := calculateSumOfRiskLevels(heightmap)
	productOfTop3BasinSizes := getProductOfTop3BasinSizes(heightmap)

	fmt.Printf("Sum of risk levels of low points: %d\n", sumOfRiskLevels)
	fmt.Printf("Product of the sizes of the largest 3 basins: %d\n", productOfTop3BasinSizes)
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

func detectBasin(heightmap Heightmap, lowPoint Point) []Point {
	alreadyChecked := make(map[Point]bool)
	pointsToCheck := []Point{lowPoint}

	for len(pointsToCheck) > 0 {
		var pointToCheck Point
		pointToCheck, pointsToCheck = pointsToCheck[0], pointsToCheck[1:]
		alreadyChecked[pointToCheck] = true

		for i := 1; heightmap.getHeightForLowpoints(pointToCheck.x+i, pointToCheck.y) < 9; i++ {
			newPoint := Point{x: pointToCheck.x + i, y: pointToCheck.y}
			if alreadyChecked[newPoint] {
				continue
			}

			pointsToCheck = append(pointsToCheck, newPoint)
		}

		for i := 1; heightmap.getHeightForLowpoints(pointToCheck.x-i, pointToCheck.y) < 9; i++ {
			newPoint := Point{x: pointToCheck.x - i, y: pointToCheck.y}
			if alreadyChecked[newPoint] {
				continue
			}

			pointsToCheck = append(pointsToCheck, newPoint)
		}

		for i := 1; heightmap.getHeightForLowpoints(pointToCheck.x, pointToCheck.y+i) < 9; i++ {
			newPoint := Point{x: pointToCheck.x, y: pointToCheck.y + i}
			if alreadyChecked[newPoint] {
				continue
			}

			pointsToCheck = append(pointsToCheck, newPoint)
		}

		for i := 1; heightmap.getHeightForLowpoints(pointToCheck.x, pointToCheck.y-i) < 9; i++ {
			newPoint := Point{x: pointToCheck.x, y: pointToCheck.y - i}
			if alreadyChecked[newPoint] {
				continue
			}

			pointsToCheck = append(pointsToCheck, newPoint)
		}
	}

	var points []Point
	for point, _ := range alreadyChecked {
		points = append(points, point)
	}

	return points
}

func getProductOfTop3BasinSizes(heightmap Heightmap) int {
	product := 1

	lowPoints := findLowpoints(heightmap)

	var basins [][]Point
	for _, lowPoint := range lowPoints {
		basin := detectBasin(heightmap, lowPoint)
		basins = append(basins, basin)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	for i := 0; i < 3; i++ {
		product *= len(basins[i])
	}

	return product
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
