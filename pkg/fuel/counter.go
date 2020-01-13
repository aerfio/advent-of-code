package fuel

import (
	"math"
	"strconv"
)

const (
	factor     = 3
	toSubtract = 2
)

type Module struct {
	mass int
}

func (m *Module) SetMass(mass int) {
	m.mass = mass
}

func (m *Module) SetMassFromString(arg string) error {
	num, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}
	m.mass = num
	return nil
}

func compute(arg int) int {
	return int(math.Floor(float64(arg/factor))) - toSubtract
}

func computeRec(arg int, acc int) int {
	data := compute(arg)

	if data <= 0 {
		return acc
	}

	return computeRec(data, acc+data)
}

func (m Module) ComputeDemand() int {
	return computeRec(m.mass, 0)
}
