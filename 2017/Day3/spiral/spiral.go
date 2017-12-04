package spiral

import (
	"math"
)

type Spiral struct {
	Board     map[int]Point
	Sums      map[Point]int
	Direction int
	Position  int
	Repeat    int
}

func (s *Spiral) Move(pos Point) *Point {
	switch s.Direction {
	case 0:
		pos.Right()
	case 1:
		pos.Up()
	case 2:
		pos.Left()
	case 3:
		pos.Down()
	}
	s.Board[s.Position] = pos
	return &pos
}

func (s *Spiral) Distance(p1, p2 int) float64 {
	x1, y1 := s.Board[p1].X, s.Board[p1].Y
	x2, y2 := s.Board[p2].X, s.Board[p2].Y
	return math.Abs(x1-x2) + math.Abs(y1-y2)
}

type Point struct {
	X, Y float64
}

func (p *Point) Right() {
	p.X = p.X + 1
}

func (p *Point) Left() {
	p.X = p.X - 1
}

func (p *Point) Up() {
	p.Y = p.Y + 1
}

func (p *Point) Down() {
	p.Y = p.Y - 1
}

func (p *Point) SumAdjacent(s *Spiral) int {
	sum := 0
	directions := []Point{
		Point{X: p.X, Y: (p.Y - 1)},
		Point{X: p.X, Y: (p.Y + 1)},
		Point{X: p.X + 1, Y: (p.Y)},
		Point{X: p.X - 1, Y: (p.Y)},
		Point{X: (p.X - 1), Y: (p.Y + 1)},
		Point{X: (p.X + 1), Y: (p.Y + 1)},
		Point{X: (p.X + 1), Y: (p.Y - 1)},
		Point{X: (p.X - 1), Y: (p.Y - 1)},
	}

	for _, direction := range directions {
		if val, ok := s.Sums[direction]; ok {
			sum += val
		}
	}

	return sum
}
