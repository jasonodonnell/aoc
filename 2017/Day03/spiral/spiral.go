package spiral

import (
	"math"
)

// Spiral is the grid and holds state about each move
type Spiral struct {
	Board     map[int]Point
	Sums      map[Point]int
	Direction int
	Position  int
	Repeat    int
}

// Move walks the grid
func (s *Spiral) Move(pos Point) *Point {
	switch s.Direction {
	case 0:
		pos.right()
	case 1:
		pos.up()
	case 2:
		pos.left()
	case 3:
		pos.down()
	}
	s.Board[s.Position] = pos
	return &pos
}

// Distance calculates the Manhattan Distance between two points
func (s *Spiral) Distance(p1, p2 int) float64 {
	x1, y1 := s.Board[p1].X, s.Board[p1].Y
	x2, y2 := s.Board[p2].X, s.Board[p2].Y
	return math.Abs(x1-x2) + math.Abs(y1-y2)
}

// Point is a vector coordinate on the grid
type Point struct {
	X, Y float64
}

// SumAdjacent sums all adjacent points around a point
func (p *Point) SumAdjacent(s *Spiral) int {
	sum := 0
	directions := []Point{
		Point{X: p.X, Y: (p.Y + 1)},       // North
		Point{X: p.X, Y: (p.Y - 1)},       // South
		Point{X: (p.X + 1), Y: (p.Y)},     // East
		Point{X: (p.X - 1), Y: (p.Y)},     // West
		Point{X: (p.X - 1), Y: (p.Y + 1)}, // North West
		Point{X: (p.X + 1), Y: (p.Y + 1)}, // North East
		Point{X: (p.X - 1), Y: (p.Y - 1)}, // South West
		Point{X: (p.X + 1), Y: (p.Y - 1)}, // South East
	}

	for _, direction := range directions {
		if val, ok := s.Sums[direction]; ok {
			sum += val
		}
	}

	return sum
}

func (p *Point) right() {
	p.X = p.X + 1
}

func (p *Point) left() {
	p.X = p.X - 1
}

func (p *Point) up() {
	p.Y = p.Y + 1
}

func (p *Point) down() {
	p.Y = p.Y - 1
}
