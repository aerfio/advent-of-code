package grid

import (
	"errors"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type movement struct {
	direction direction
	distance  int
}

type movementSlice []movement

type point struct {
	x int
	y int
}

var tooShortPathErr = errors.New("path should be at least 2 points long")
var twoSamePointsNextToEachOtherErr = errors.New("two adjacent points should not be the same")


type path []point
