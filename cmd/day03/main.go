package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	report, err := readDiagnosticReport(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	gammaRate := calculateGammaRate(report)
	deltaRate := getDeltaRate(gammaRate)

	gamma := convertToNumber(gammaRate)
	delta := convertToNumber(deltaRate)

	fmt.Printf("Power consumption: %d\n\tGamma: %d\n\tDelta: %d\n", gamma*delta, gamma, delta)
}

type DiagnosticReport []ReportRow
type ReportRow []bool

func calculateGammaRate(report DiagnosticReport) ReportRow {
	reportRows := len(report)
	threshold := reportRows / 2
	numberOfColumns := len(report[0])
	numberOfOnes := make([]int, numberOfColumns)

	for _, row := range report {
		for columnIndex, entry := range row {
			if entry {
				numberOfOnes[columnIndex]++
			}
		}
	}

	condensedRow := make(ReportRow, numberOfColumns)
	for index, count := range numberOfOnes {
		if count == threshold {
			log.Fatalln("no most common bit")
		}

		if count > threshold {
			condensedRow[index] = true
		}
	}

	return condensedRow
}

func getDeltaRate(gammaRate ReportRow) ReportRow {
	deltaRate := make(ReportRow, len(gammaRate))

	for index, isSet := range gammaRate {
		if !isSet {
			deltaRate[index] = true
		}
	}

	return deltaRate
}

func convertToNumber(row ReportRow) int {
	var result int
	for index, isSet := range row {
		if isSet {
			// Set corresponding bit
			bitPosition := len(row) - 1 - index
			result |= (1 << bitPosition)
		}
	}
	return result
}

func readDiagnosticReport(input io.Reader) (DiagnosticReport, error) {
	var report DiagnosticReport

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		row := make(ReportRow, len(line))
		for index, bit := range line {
			if bit == rune('1') {
				row[index] = true
			}
		}

		report = append(report, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return report, nil
}
