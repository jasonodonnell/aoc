package hash

import (
	"fmt"
	"strconv"
)

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

func (h *Hash) XOR() {
	i := 0
	for i < len(h.List) {
		// Not sure how to do this on a slice.. hacky for now.
		denseHash := int64(h.List[i] ^ h.List[i+1] ^
			h.List[i+2] ^ h.List[i+3] ^ h.List[i+4] ^
			h.List[i+5] ^ h.List[i+6] ^ h.List[i+7] ^
			h.List[i+8] ^ h.List[i+9] ^ h.List[i+10] ^
			h.List[i+11] ^ h.List[i+12] ^ h.List[i+13] ^
			h.List[i+14] ^ h.List[i+15])
		fmt.Printf(strconv.FormatInt(denseHash, 16))
		i += 16
	}
	fmt.Println("")
}
