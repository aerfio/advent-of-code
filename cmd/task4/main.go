package main

import (
	"fmt"
	"log"

	"github.com/aerfio/advent-of-code/pkg/input/task4"
	secure_container "github.com/aerfio/advent-of-code/pkg/secure-container"
)

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lowerBound, upperBound, err := secure_container.ParseRange(task4.GivenRange)
	failIfErr(err)

	numberOfCorrectPasswords := 0

	for i := lowerBound; i <= upperBound; i++ {
		pass, err := secure_container.ConvertToPassword(i)
		failIfErr(err)
		adjRule := pass.CheckAdjacentRule()
		notDecreasingRule:=pass.SatisfiesNotDecreasingRule()
		if adjRule && notDecreasingRule{
			numberOfCorrectPasswords += 1
		}
	}
	fmt.Println(numberOfCorrectPasswords)
}
