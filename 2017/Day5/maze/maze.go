package maze

type Maze struct {
	Instructions []int
	position     int
	offset       int
	Steps        int
}

func (m *Maze) Move(advanced bool) (escaped bool) {
	if m.position >= len(m.Instructions) {
		return true
	}

	m.offset = m.Instructions[m.position]
	if advanced && m.offset >= 3 {
		m.Instructions[m.position]--
	} else {
		m.Instructions[m.position]++
	}

	m.position = m.position + m.offset
	m.Steps++
	return false
}
