package grid

import (
	"strconv"
	"strings"
)

type Path struct {
}

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

// func (d Direction) String() string {
// 	return [...]string{"Up", "Right", "Down", "Left"}[d]
// }

func stringToDir(token string) direction {
	switch strings.ToLower(token) {
	case "u":
		return up
	case "r":
		return right
	case "d":
		return down
	case "l":
		return left
	default:
		return up
	}
}

func parse(arg string) (movementSlice, error) {
	split := strings.Split(arg, ",")

	mov := make(movementSlice, 0, len(split))

	for _, token := range split {
		dir := stringToDir(string(token[0]))
		dist, err := strconv.Atoi(token[1:])
		if err != nil {
			return movementSlice{}, err
		}

		mov = append(mov, movement{
			direction: dir,
			distance:  dist,
		})
	}
	return mov, nil
}

func New(data string) (movementSlice, error) {
	return parse(data)
}

type point struct {
	x int
	y int
}

func (ms movementSlice) toCornerPoints() []point {
	grid := make([]point, 1) // one element here is on purpose

	for _, m := range ms {
		lastElem := grid[len(grid)-1]
		switch m.direction {
		case up:
			{
				newGridElem := point{
					x: lastElem.x,
					y: lastElem.y + m.distance,
				}
				grid = append(grid, newGridElem)
			}
		case down:
			{
				newGridElem := point{
					x: lastElem.x,
					y: lastElem.y - m.distance,
				}
				grid = append(grid, newGridElem)
			}
		case right:
			{
				newGridElem := point{
					x: lastElem.x + m.distance,
					y: lastElem.y,
				}
				grid = append(grid, newGridElem)
			}
		case left:
			{
				newGridElem := point{
					x: lastElem.x - m.distance,
					y: lastElem.y,
				}
				grid = append(grid, newGridElem)
			}
		}

	}

	return grid
}
