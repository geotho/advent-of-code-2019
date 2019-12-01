package aoc01

import (
	"bufio"
	"io"
	"strconv"
)

const biggestNumberThatRequiresNoFuel = 2*3 + 2

func fuel(i int) int {
	if i <= biggestNumberThatRequiresNoFuel {
		return 0
	}
	return i/3 - 2
}

func totalFuel(i int) int {
	f := fuel(i)
	total := 0
	for f > 0 {
		total += f
		f = fuel(f)
	}
	return total
}

func totalFuelForFile(r io.Reader) (int, error) {
	scn := bufio.NewScanner(r)
	total := 0
	for scn.Scan() {
		line := scn.Text()
		weight, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		total += totalFuel(weight)
	}
	err := scn.Err()
	if err != nil {
		return 0, err
	}
	return total, nil
}
