package main

import (
	"fmt"
	"log"

	"github.com/aerfio/advent-of-code/pkg/input/task2"
	"github.com/aerfio/advent-of-code/pkg/intcode"
)

func main() {
	for noun := 0; noun < 100; noun = noun + 1 {
		for verb := 0; verb < 100; verb = verb + 1 {
			prog, err := intcode.New(task2.Data)
			if err != nil {
				log.Fatal(err)
			}
			runOpts := intcode.RunOpts{InitialNoun: noun, InitialVerb: verb}
			prog.Run(&runOpts)

			firstNum := prog.GetOutput()

			if firstNum == task2.DesiredOutput {
				fmt.Printf("%+v\n", runOpts)
				fmt.Printf("100 * noun + verb = %d\n", 100*noun+verb)
				// do not break -> maybe there's more answers?
			}
		}
	}
}
