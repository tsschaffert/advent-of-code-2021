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

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func applyCommand(position Position, command Command) Position {
	switch command.Direction {
	case Forward:
		position.Horizontal += command.Distance
	case Up:
		position.Depth -= command.Distance
	case Down:
		position.Depth += command.Distance
	}

	return position
}

func applyCommandCorrectly(position Position, command Command) Position {
	switch command.Direction {
	case Forward:
		position.Horizontal += command.Distance
		position.Depth += position.Aim * command.Distance
	case Up:
		position.Aim -= command.Distance
	case Down:
		position.Aim += command.Distance
	}

	return position
}

func main() {
	commands, err := readCommands(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	finalPosition := applyCommands(Position{}, commands, applyCommand)
	correctFinalPosition := applyCommands(Position{}, commands, applyCommandCorrectly)

	fmt.Printf("Final position: Horizontal %d, Depth %d (product=%d)\n", finalPosition.Horizontal, finalPosition.Depth, finalPosition.Horizontal*finalPosition.Depth)
	fmt.Printf("Correct final position: Horizontal %d, Depth %d (product=%d)\n", correctFinalPosition.Horizontal, correctFinalPosition.Depth, correctFinalPosition.Horizontal*correctFinalPosition.Depth)
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

func applyCommands(initialPosition Position, commands []Command, applyFunction func(position Position, command Command) Position) Position {
	currentPosition := initialPosition

	for _, command := range commands {
		currentPosition = applyFunction(currentPosition, command)
	}

	return currentPosition
}
