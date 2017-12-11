package hex

import (
	"math"
)

type Hex struct {
	Position Point
	Furthest float64
	Distance float64
}

type Point struct {
	X, Y float64
}

func (h *Hex) Move(direction string) {
	switch direction {
	case "n":
		h.Position.Y++
	case "ne":
		h.Position.X++
	case "nw":
		h.Position.X--
		h.Position.Y++
	case "s":
		h.Position.Y--
	case "se":
		h.Position.X++
		h.Position.Y--
	case "sw":
		h.Position.X--
	}
	h.Distance = h.distance()
	h.highmark()
}

func (h *Hex) distance() float64 {
	z := computeZ(h.Position)
	max := math.Max(h.Position.X, h.Position.Y)
	return math.Max(max, z)
}

func (h *Hex) highmark() {
	if h.Distance > h.Furthest {
		h.Furthest = h.Distance
	}
}

func computeZ(p Point) float64 {
	return math.Abs(p.X + p.Y)
}
