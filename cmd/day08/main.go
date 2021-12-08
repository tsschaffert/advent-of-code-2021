package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	readings, err := readSignals(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	uniqueSignals := countUniqueSignals(readings)

	fmt.Printf("Number of signals with unique of segments: %d\n", uniqueSignals)
}

type Signal string

type Reading struct {
	signalPatterns []Signal
	output         []Signal
}

func countUniqueSignals(readings []Reading) int {
	count := 0
	for _, reading := range readings {
		for _, output := range reading.output {
			length := len(output)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}
	return count
}

func readSignals(input io.Reader) ([]Reading, error) {
	var readings []Reading

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " | ", 2)
		signals := strings.Split(parts[0], " ")
		outputs := strings.Split(parts[1], " ")

		reading := Reading{}
		for _, signal := range signals {
			reading.signalPatterns = append(reading.signalPatterns, Signal(signal))
		}
		for _, output := range outputs {
			reading.output = append(reading.output, Signal(output))
		}

		readings = append(readings, reading)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return readings, nil
}
