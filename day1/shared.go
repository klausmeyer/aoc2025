package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type direction string

const (
	DirectionLeft  direction = "L"
	DirectionRight direction = "R"
)

type rotation struct {
	direction direction
	steps     int
}

func parseDirection(line string) direction {
	return direction(line[0])
}

func parseSteps(line string) int {
	num, err := strconv.Atoi(line[1:])

	if err != nil {
		log.Fatalf("can not convert to number: %s", err)
	}

	return num
}

func parseLine(line string) rotation {
	return rotation{direction: parseDirection(line), steps: parseSteps((line))}
}

func ReadSequence(filename string) []rotation {
	var sequence []rotation

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sequence = append(sequence, parseLine(scanner.Text()))
	}

	return sequence
}
