package day1

import (
	"os"
	"strconv"
	"strings"
)

type Rotation struct {
	Direction string
	Distance  int
}

func getPass1(file string) int {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	var position = 50
	password := 0

	for _, l := range lines {
		direction := l[:1]
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		if direction == "R" {
			position += distance
		} else {
			position -= distance
		}
		position = position % 100

		if position == 0 {
			password++
		}
	}

	return password
}

func getPass2(file string) int {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	position := 50
	password := 0
	
	for _, l := range lines {
		initialPosition := position
		direction := l[:1]
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		password += distance / 100
		distance = distance % 100

		if direction == "L" {
			distance *= -1
		}

		if (position + distance) > 100 {
			increment(&password, initialPosition)
		}
		position = (position + distance) % 100

		if position < 0 {
			increment(&password, initialPosition)
			position += 100
		}

		if position == 0 {
			increment(&password, initialPosition)
		}

	
	}

	return password
}

func increment(password *int, initialPos int) {
	if initialPos != 0 {
		*password++
	}
}
