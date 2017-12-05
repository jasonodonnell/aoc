package maze

type Maze struct {
	Instructions []int
	Position     int
	Steps        int
}

func (m *Maze) Move(offset int, advanced bool) {
	m.Instructions[m.Position]++
	if advanced {
		if offset >= 3 {
			m.Instructions[m.Position] = m.Instructions[m.Position] - 2
		}
	}
	m.Position = m.Position + offset
	m.Steps++
}
