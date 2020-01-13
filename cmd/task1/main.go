package main

import (
	"log"
	"strconv"

	"github.com/aerfio/advent-of-code/pkg/fuel"
	"github.com/aerfio/advent-of-code/pkg/input/task1"
	"github.com/atotto/clipboard"
)

func main() {
	sum := 0
	for _, line := range task1.Prepared {
		mod := fuel.Module{}
		if err := mod.SetMassFromString(line); err != nil {
			log.Fatal(err)
		}

		sum += mod.ComputeDemand()
	}

	if err := clipboard.WriteAll(strconv.Itoa(sum)); err != nil {
		log.Fatal(err)
	}
}
