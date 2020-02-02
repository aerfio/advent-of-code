package secure_container

import (
	"errors"
	"strconv"
	"strings"
)

func ParseRange(arg string) (int, int, error) {
	ranges := strings.Split(arg, "-")
	lowerBound, err := strconv.Atoi(ranges[0])
	if err != nil {
		return 0, 0, err
	}
	upperBound, err := strconv.Atoi(ranges[1])
	if err != nil {
		return 0, 0, err
	}

	return lowerBound, upperBound, nil
}

type Password [6]int

func ConvertToPassword(arg int) (Password, error) {
	seq := make([]int, 0, 6)
	data := intToSlice(arg, seq)

	if len(data) != 6 {
		return Password{}, errors.New("is not a correct password")
	}

	return Password{
		data[0],
		data[1],
		data[2],
		data[3],
		data[4],
		data[5],
	}, nil
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

func (p Password) CheckAdjacentRule() bool {
	isSameAdjacent := false

	for i := 0; i < 5; i++ {
		if p[i] == p[i+1] {
			isSameAdjacent = true
		}
	}
	return isSameAdjacent
}

func (p Password) SatisfiesNotDecreasingRule() bool {
	isDecreasing := false

	for i := 0; i < 5; i++ {
		if p[i+1] < p[i] {
			isDecreasing = true
		}
	}

	return !isDecreasing
}
