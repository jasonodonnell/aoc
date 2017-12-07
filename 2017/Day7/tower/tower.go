package tower

import (
	"math"
)

type Program struct {
	Name        string
	Weight      float64
	Children    []string
	TotalWeight float64
}

type Tower struct {
	Programs map[string]*Program
	Bases    []string
}

func (t *Tower) FindBase() string {
Loop:
	for _, parent := range t.Bases {
		for _, program := range t.Programs {
			// Ignore self or leaf nodes
			if program.Name == parent || program.Children == nil {
				continue
			}
			for _, each := range program.Children {
				if parent == each {
					continue Loop
				}
			}
		}
		return parent
	}
	return ""
}

func (t *Tower) totalWeight(root string) float64 {
	if t.Programs[root].Children != nil {
		for _, v := range t.Programs[root].Children {
			t.Programs[root].TotalWeight += +t.totalWeight(v)
		}
		t.Programs[root].TotalWeight += t.Programs[root].Weight
	} else {
		t.Programs[root].TotalWeight += t.Programs[root].Weight
	}
	return t.Programs[root].TotalWeight
}

// FindBalanced is a desperate attempt at solving this problem and
// should not reflect too harshly on the author.
func (t *Tower) FindUnbalanced(root string) float64 {
	_ = t.totalWeight(root)
	var unbalanced string
Loop:
	for _, program := range t.Programs {
		programs := make(map[float64]int)
		// Leaf node
		if program.Children == nil || program.Name == root {
			continue
		}
		// Not leaf, for each child entry, hash weights
		// If hash has more then one entry, its unbalanced
		for _, child := range program.Children {
			programs[t.Programs[child].TotalWeight]++
			// The culprit will have balanced child entries,
			// so we need to find that next and return.
			if len(programs) > 1 {
				balanced := make(map[float64]int)
				for _, v := range t.Programs[child].Children {
					balanced[t.Programs[v].TotalWeight]++
				}
				if len(balanced) == 1 {
					unbalanced = child
					break Loop
				}
				continue Loop
			}
		}
	}
	parent := t.getParent(unbalanced)
	return t.weightAdjustment(unbalanced, parent)
}

func (t *Tower) getParent(node string) string {
	for _, program := range t.Programs {
		if program.Children == nil {
			continue
		}
		for _, v := range program.Children {
			if v == node {
				return program.Name
			}
		}
	}
	return ""
}

func (t *Tower) weightAdjustment(unbalanced, parent string) float64 {
	var unbalancedWeight float64
	var parentWeight float64
	for _, v := range t.Programs[parent].Children {
		if v == unbalanced {
			continue
		}
		unbalancedWeight = math.Abs(t.Programs[v].TotalWeight - t.Programs[unbalanced].TotalWeight)
		parentWeight = t.Programs[unbalanced].TotalWeight
		break
	}
	return parentWeight - unbalancedWeight
}
