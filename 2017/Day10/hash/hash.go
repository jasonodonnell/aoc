package hash

type Hash struct {
	List       []int
	currentPos int
	skipSize   int
	Length     int
}

func NewHash(length int) *Hash {
	var l []int
	for i := 0; i < 256; i++ {
		l = append(l, i)
	}
	return &Hash{
		List:       l,
		currentPos: 0,
		skipSize:   0,
	}
}

func (h *Hash) Reverse() {
	var indexes []int
	var values []int
	for i := 0; i < h.Length; i++ {
		index := (i + h.currentPos) % len(h.List)
		indexes = append(indexes, index)
		values = append(values, h.List[index])
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) - 1 - i
		values[i], values[opp] = values[opp], values[i]
	}

	count := 0
	for _, index := range indexes {
		h.List[index] = values[count]
		count++
	}
	h.currentPos = (h.currentPos + h.Length + h.skipSize) % len(h.List)
	h.skipSize++
}
