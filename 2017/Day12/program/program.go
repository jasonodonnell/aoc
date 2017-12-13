package program

type Programs struct {
	Contains int
	Pipes    map[string][]string
	Visited  map[string]bool
}

func (p *Programs) WalkGroup(id string) {
	if _, ok := p.Visited[id]; !ok {
		p.Visited[id] = true
		p.Contains++
		for _, v := range p.Pipes[id] {
			p.WalkGroup(v)
		}
	}
}
