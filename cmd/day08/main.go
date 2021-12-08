package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	readings, err := readSignals(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	uniqueSignals := countUniqueSignals(readings)

	sum := 0
	analysedReadings := analyseMappings(readings)
	for _, reading := range analysedReadings {
		number := convertToNumber(reading)
		sum += number
		fmt.Println(number)
	}

	fmt.Printf("Number of signals with unique of segments: %d\n", uniqueSignals)
	fmt.Printf("Sum of all numbers in output: %d\n", sum)
}

type Signal string

func (s Signal) equals(other Signal) bool {
	if len(s) != len(other) {
		return false
	}
	return s.countOverlap(other) == len(s)
}

func (s Signal) countOverlap(other Signal) int {
	overlap := 0
	for _, char := range s {
		if strings.ContainsRune(string(other), char) {
			overlap++
		}
	}
	return overlap
}

type Reading struct {
	signalPatterns []Signal
	output         []Signal
	mapping        map[Signal]int
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

func analyseMappings(readings []Reading) []Reading {
	var mappedReadings []Reading

	for _, reading := range readings {
		reverseMapping := make(map[int]Signal)
		mapping := make(map[Signal]int)

		mapTo := 1
		signal := findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 2
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 7
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 3
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 4
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 4
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 8
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 7
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 3
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 5 && signal.countOverlap(m[7]) == len(m[7])
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 5
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 5 && !signal.equals(m[3]) && signal.countOverlap(m[4]) == 3
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 2
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 5 && !signal.equals(m[3]) && !signal.equals(m[5])
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 6
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 6 && signal.countOverlap(m[7]) == 2
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 9
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 6 && !signal.equals(m[6]) && signal.countOverlap(m[4]) == len(m[4])
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mapTo = 0
		signal = findMatch(reading.signalPatterns, reverseMapping, func(signal Signal, m map[int]Signal) bool {
			return len(signal) == 6 && !signal.equals(m[6]) && !signal.equals(m[9])
		})
		mapping[signal] = mapTo
		reverseMapping[mapTo] = signal

		mappedReadings = append(mappedReadings, Reading{
			signalPatterns: reading.signalPatterns,
			output:         reading.output,
			mapping:        mapping,
		})
	}

	return mappedReadings
}

func convertToNumber(reading Reading) int {
	result := 0
	for index, output := range reading.output {
		result += lookupInMap(output, reading.mapping) * int(math.Pow10(3-index))
	}
	return result
}

func lookupInMap(output Signal, mapping map[Signal]int) int {
	for signal, value := range mapping {
		if signal.equals(output) {
			return value
		}
	}
	log.Fatalf("no match found for %s in %v", output, mapping)
	return -1
}

func findMatch(signalPatterns []Signal, reverseMapping map[int]Signal, matches func(Signal, map[int]Signal) bool) Signal {
	for _, signal := range signalPatterns {
		if matches(signal, reverseMapping) {
			return signal
		}
	}
	// TODO return error
	log.Fatalf("no match found")
	return ""
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
