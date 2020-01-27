package main

import (
	"fmt"
	"log"

	"github.com/aerfio/advent-of-code/pkg/grid"
	"github.com/aerfio/advent-of-code/pkg/input/task3"
)

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data := task3.MyData

	ms, err := grid.New(data.First)
	failIfErr(err)
	path, err := ms.GetPath()
	failIfErr(err)

	ms2, err := grid.New(data.Second)
	failIfErr(err)
	path2, err := ms2.GetPath()
	failIfErr(err)

	pts := path.FindIntersection(path2)

	fmt.Println(grid.FindManhattanDistanceOfNearestPoint(pts))
}
