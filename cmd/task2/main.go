package main

import (
	"log"
	"os"

	"github.com/aerfio/advent-of-code/pkg/input/task2"
	"github.com/aerfio/advent-of-code/pkg/intcode"
)

func main() {
	prog, err := intcode.New(task2.Data)
	if err != nil {
		log.Fatal(err)

	}

	prog.Run(true)

	if err := prog.PrintAll(os.Stdout); err != nil {
		log.Fatal(err)
	}

}
