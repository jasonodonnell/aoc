package redistribution

import (
	"strconv"
	"strings"
)

type Memory struct {
	Banks               []int
	Blocks              map[string]int
	RedistributionCount int
	index               int
}

func (m *Memory) BankExists() bool {
	for _, v := range m.Blocks {
		if v > 1 {
			return true
		}
	}
	return false
}

func (m *Memory) BlockRedistribution() {
	block, value := m.biggestBlock()
	m.index = (block + 1) % len(m.Banks)
	m.Banks[block] = 0

	for i := 0; i < value; i++ {
		m.Banks[m.index]++
		m.index = (m.index + 1) % len(m.Banks)
	}
	m.Blocks[m.Stringify(m.Banks)]++
	m.RedistributionCount++
}

func (m *Memory) biggestBlock() (index, greatest int) {
	greatest = -1
	for k, v := range m.Banks {
		if v > greatest {
			greatest = v
			index = k
		}
	}
	return index, greatest
}

func (m *Memory) Stringify(nums []int) string {
	value := []string{}
	for _, v := range nums {
		text := strconv.Itoa(v)
		value = append(value, text)
	}
	return strings.Join(value, "")
}
