package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type DirectionType int

const (
	Forward DirectionType = iota
	Up
	Down
	Undefined
)

func NewDirection(value string) DirectionType {
	switch value {
	case "forward":
		return Forward
	case "up":
		return Up
	case "down":
		return Down
	default:
		return Undefined
	}
}

type Command struct {
	Direction DirectionType
	Distance  int
}

func main() {

}

func readCommands(input io.Reader) ([]Command, error) {
	var commands []Command

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.SplitN(line, " ", 2)

		direction := NewDirection(split[0])
		if direction == Undefined {
			return nil, fmt.Errorf("unknown direction '%s'", split[0])
		}

		distance, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		commands = append(commands, Command{
			Direction: direction,
			Distance:  distance,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
