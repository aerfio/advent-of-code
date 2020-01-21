package main

import (
	"fmt"
	"log"

	"github.com/aerfio/advent-of-code/pkg/input/task2"
	"github.com/aerfio/advent-of-code/pkg/intcode"
)

func main() {
	prog, err := intcode.New(task2.Data)
	if err != nil {
		log.Fatal(err)

	}

	runOpts := intcode.RunOpts{InitialNoun: 2, InitialVerb: 12}

	prog.Run(&runOpts)

	firstNum := prog.GetOutput()
	fmt.Println(firstNum)
}
