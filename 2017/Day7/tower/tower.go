package tower

type Program struct {
	Name        string
	Weight      int
	Supports    []string
	TotalWeight int
	Unbalanced  bool
}

type Tower struct {
	Programs []*Program
	Bases    []string
}

func (t *Tower) FindBase() string {
Loop:
	for _, v := range t.Bases {
		for _, x := range t.Programs {
			if x.Name == v || x.Supports == nil {
				continue
			}
			for _, each := range x.Supports {
				if v == each {
					continue Loop
				}
			}
		}
		return v
	}
	return ""
}
