package knot

import (
	"strconv"
)

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
	i := 0
	var hash string
	for i < len(k.List) {
		// Not sure how to do this on a slice.. hacky for now.
		// Doing them two at a time didn't seem to work
		denseHash := int64(k.List[i] ^ k.List[i+1] ^ k.List[i+2] ^ k.List[i+3] ^
			k.List[i+4] ^ k.List[i+5] ^ k.List[i+6] ^ k.List[i+7] ^
			k.List[i+8] ^ k.List[i+9] ^ k.List[i+10] ^ k.List[i+11] ^
			k.List[i+12] ^ k.List[i+13] ^ k.List[i+14] ^ k.List[i+15])
		hash += strconv.FormatInt(denseHash, 16)
		i += 16
	}
	return hash
}
