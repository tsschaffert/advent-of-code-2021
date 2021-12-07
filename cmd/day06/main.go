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
	convertedPopulation := convertToDensePopulation(population)

	endPopulation := simulateDensePopulation(convertedPopulation, 80)
	largeEndPopulation := simulateDensePopulation(convertedPopulation, 256)

	fmt.Printf("Population after 80 days: %d\n", endPopulation.size())
	fmt.Printf("Population after 256 days: %d\n", largeEndPopulation.size())
}

type densePopulation struct {
	buckets [9]int64
}

func (p *densePopulation) simulate() {
	oldBuckets := p.buckets

	for i := 0; i < 8; i++ {
		p.buckets[i] = oldBuckets[i+1]
	}
	p.buckets[8] = oldBuckets[0]
	p.buckets[6] += oldBuckets[0]
}

func (p densePopulation) size() int64 {
	var size int64
	for _, count := range p.buckets {
		size += count
	}
	return size
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
	// TODO Remove side-effects introduced by this and access by index in the loop
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

func simulateDensePopulation(population densePopulation, steps int) densePopulation {
	for i := 0; i < steps; i++ {
		population.simulate()
	}
	return population
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

func convertToDensePopulation(population []lanternfish) densePopulation {
	var result densePopulation

	for _, fish := range population {
		result.buckets[fish.timer]++
	}

	return result
}
