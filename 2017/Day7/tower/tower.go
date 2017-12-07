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
	supports := make(map[string]int)
	for _, v := range t.Bases {
		if _, ok := supports[v]; !ok {
			supports[v] = 0
		}
		for _, x := range t.Programs {
			if x.Name == v || x.Supports == nil {
				continue
			}
			for _, each := range x.Supports {
				if v == each {
					supports[v]++
				}
			}
		}
	}
	for k, v := range supports {
		if v == 0 {
			return k
		}
	}
	return ""
}
