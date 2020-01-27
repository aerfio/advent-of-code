package grid

import (
	"sort"
	"strconv"
	"strings"
)

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

func (ms movementSlice) toCornerPoints() path {
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

func getPathBetweenPoints(p1, p2 point) (path, error) {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	if dx == 0 && dy == 0 {
		return path{}, twoSamePointsNextToEachOtherErr
	}

	var out path
	if dx != 0 {
		out = make(path, 0, abs(dx))
	} else {
		out = make(path, 0, abs(dy))
	}

	if dx > 0 {
		for m := 0; m < dx-1; m++ {
			out = append(out, point{y: p1.y, x: p1.x + m + 1})
		}
	} else {
		for m := 0; m > dx+1; m-- {
			out = append(out, point{y: p1.y, x: p1.x + m - 1})
		}
	}

	if dy > 0 {
		for m := 0; m < dy-1; m++ {
			out = append(out, point{y: p1.y + m + 1, x: p1.x})
		}
	} else {
		for m := 0; m > dy+1; m-- {
			out = append(out, point{y: p1.y + m - 1, x: p1.x})
		}
	}

	return out, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (ms movementSlice) GetPath() (path, error) {
	cp := ms.toCornerPoints()
	cornerPointLength := len(cp)

	out := make(path, 0, cornerPointLength*2)

	for i := 0; i < cornerPointLength-1; i++ {
		out = append(out, cp[i])
		p, err := getPathBetweenPoints(cp[i], cp[i+1])
		if err != nil {
			return path{}, err
		}
		out = append(out, p...)

	}
	out = append(out, cp[cornerPointLength-1])

	return out, nil
}

func (p path) FindIntersection(next path) []point {
	out := make([]point, 0, 0)

	for _, pt := range p {
		for _, nextPt := range next {
			if pt.isCentral() && nextPt.isCentral() {
				continue
			}

			if pt.x == nextPt.x && pt.y == nextPt.y {
				out = append(out, pt)
			}
		}
	}

	return out
}

func (p point) isCentral() bool {
	return p.x == 0 && p.y == 0
}

func FindManhattanDistanceOfNearestPoint(pts []point) int {
	distances := make([]int, 0, len(pts))
	for _, pt := range pts {
		distances = append(distances, abs(pt.x)+abs(pt.y))
	}
	sort.Ints(distances)
	return distances[0]
}

// TODO test me bitch
func (p path) findDistanceToIntersection(inter point) int {
	dist := 0

	for i := 0; i < len(p)-1; i++ {
		firstPoint := p[i]
		secondPoint := p[i+1]
		if isBetween := inter.isBetweenTwoPoints(firstPoint, secondPoint); isBetween {
			dist += firstPoint.manhattanDistance(inter)
			break
		}
		dist += firstPoint.manhattanDistance(secondPoint)
	}
	return dist
}

func (p point) manhattanDistance(next point) int {
	return abs(p.x-next.x) + abs(p.y-next.y)
}

func (p point) isBetweenTwoPoints(p1, p2 point) bool {
	// there could be test for p1==p2==p or something like this, but come on...

	if p1.x < p.x && p.x < p2.x && p.y == p1.y && p.y == p2.y {
		// p is between p1 and p2 on X axis
		return true
	} else if p.x < p1.x && p.x > p2.x && p.y == p1.y && p.y == p2.y {
		// p is between p1 and p2 on X axis when p1 is on left side of p2
		return true
	} else if p1.y < p.y && p.y < p2.y && p.x == p1.x && p.x == p2.x {
		// p is between p1 and p2 on Y axis
		return true
	} else if p.y < p1.y && p.y > p2.y && p.x == p1.x && p.x == p2.x {
		// p is between p1 and p2 on X axis when p1 is on top of p2
		return true
	}
	return false
}
