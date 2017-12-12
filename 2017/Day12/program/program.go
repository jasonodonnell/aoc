package program

type Programs struct {
	Pipes    map[string][]string
	Contains int
	Visited  map[string]bool
	Groups   int
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
