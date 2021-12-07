package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	population, err := readPopulation(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	endPopulation := simulatePopulation(population, 80)

	fmt.Printf("Population after 80 days: %d\n", len(endPopulation))
}

type lanternfish struct {
	timer int
}

func (fish *lanternfish) simulate() bool {
	fish.timer--

	if fish.timer >= 0 {
		return false
	}

	fish.timer = 6
	return true
}

func newLanternfish() lanternfish {
	return lanternfish{8}
}

func simulatePopulationStep(population []lanternfish) []lanternfish {
	newPopulation := population

	for index, fish := range population {
		spawnNewFish := fish.simulate()
		newPopulation[index] = fish
		if spawnNewFish {
			newPopulation = append(newPopulation, newLanternfish())
		}
	}

	return newPopulation
}

func simulatePopulation(population []lanternfish, steps int) []lanternfish {
	newPopulation := population
	for i := 0; i < steps; i++ {
		newPopulation = simulatePopulationStep(newPopulation)
	}
	return newPopulation
}

func readPopulation(input io.Reader) ([]lanternfish, error) {
	var population []lanternfish

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		for _, entry := range strings.Split(line, ",") {
			timer, err := strconv.Atoi(entry)
			if err != nil {
				return nil, err
			}
			population = append(population, lanternfish{timer})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return population, nil
}
