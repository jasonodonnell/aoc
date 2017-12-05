package maze

type Maze struct {
	Instructions []int
	Position     int
	Steps        int
}

func (m *Maze) Move(offset int, advanced bool) {
	if advanced && offset >= 3 {
		m.Instructions[m.Position]--
	} else {
		m.Instructions[m.Position]++
	}
	m.Position = m.Position + offset
	m.Steps++
}
