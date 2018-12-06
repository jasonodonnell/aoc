package grid

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Grid struct {
	pointCount       int
	xMax, yMax       int
	xPoints, yPoints []int
	grid             [][]int
}

func (g *Grid) PrintAll() {
	fmt.Println(g.xMax, g.yMax, g.grid)
}

func New(points []Point) *Grid {
	var maxX, maxY int
	var xPoints, yPoints []int

	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
		xPoints = append(xPoints, point.X)
		yPoints = append(yPoints, point.Y)
	}

	grid := make([][]int, maxX+1)
	for i := range grid {
		grid[i] = make([]int, maxY+1)
	}

	return &Grid{
		grid:       grid,
		pointCount: len(points),
		xPoints:    xPoints,
		yPoints:    yPoints,
		xMax:       maxX,
		yMax:       maxY,
	}
}

func (g *Grid) LargestArea() int {
	points := make([]int, len(g.xPoints))
	for i := 0; i < g.xMax+1; i++ {
		for j := 0; j < g.yMax+1; j++ {
			closest := g.closestDistance(i, j)
			g.grid[i][j] = closest
			if closest != -1 && points[closest] != -1 {
				if i == g.xMax || j == g.yMax || i == 0 || j == 0 {
					points[closest]--
				} else {
					points[closest]++
				}
			}
		}
	}

	var largestArea int
	for i := 0; i < len(points); i++ {
		if points[i] > largestArea {
			largestArea = points[i]
		}
	}
	return largestArea
}

func (g *Grid) RegionSum(limit int) int {
	var region int
	for i := 0; i < g.xMax+1; i++ {
		for j := 0; j < g.yMax+1; j++ {
			sum := g.sumDistances(i, j)
			g.grid[i][j] = sum
			if sum < limit {
				region++
			}
		}
	}
	return region
}

func (g *Grid) sumDistances(x, y int) int {
	var distance int
	for i := 0; i < len(g.xPoints); i++ {
		md := manhattanDistance(x, y, g.xPoints[i], g.yPoints[i])
		distance = distance + md
	}
	return distance
}

func (g *Grid) closestDistance(x, y int) int {
	var closest int
	var equal bool
	distance := math.MaxInt32

	for i := 0; i < len(g.xPoints); i++ {
		md := manhattanDistance(x, y, g.xPoints[i], g.yPoints[i])
		if md == distance {
			equal = true
		} else if md < distance {
			distance = md
			closest = i
			equal = false
		}
	}
	if equal {
		return -1
	}
	return closest
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	x := math.Abs(float64(x1) - float64(x2))
	y := math.Abs(float64(y1) - float64(y2))
	return int(x + y)
}
