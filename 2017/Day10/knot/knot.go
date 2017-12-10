package knot

import "strconv"

type Knot struct {
	List       []int
	currentPos int
	skipSize   int
}

func NewKnot(length int) *Knot {
	var l []int
	for i := 0; i < length; i++ {
		l = append(l, i)
	}
	return &Knot{
		List:       l,
		currentPos: 0,
		skipSize:   0,
	}
}

func (k *Knot) Reverse(length int) {
	var indexes []int
	var values []int
	for i := 0; i < length; i++ {
		index := (i + k.currentPos) % len(k.List)
		indexes = append(indexes, index)
		values = append(values, k.List[index])
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) - 1 - i
		values[i], values[opp] = values[opp], values[i]
	}

	count := 0
	for _, index := range indexes {
		k.List[index] = values[count]
		count++
	}

	k.currentPos = (k.currentPos + length + k.skipSize) % len(k.List)
	k.skipSize++
}

func (k *Knot) Hash() string {
	var hash string
	dense := make([]int, len(k.List)/16)
	for i := 0; i < len(k.List); i += 16 {
		for j := 0; j < 16; j++ {
			dense[i/16] ^= k.List[i+j]
		}
	}
	for _, v := range dense {
		hash += strconv.FormatInt(int64(v), 16)
	}
	return hash
}
