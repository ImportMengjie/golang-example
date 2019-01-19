package main

import (
	"bytes"
	"fmt"
)

type bitMap struct {
	words []uint64
}

func (b *bitMap) Has(x int) bool {
	word, offset := x>>6, x&(64-1)
	return word < len(b.words) && b.words[word]&(uint64(1)<<uint(offset)) != 0

}

func (b *bitMap) Add(x int) {
	word, offset := x>>6, uint(x&(64-1))
	if word >= len(b.words) {
		b.words = append(b.words, make([]uint64, word-len(b.words)+1)...)
	}
	b.words[word] |= (1 << offset)
}

func (b *bitMap) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		// base := i * (2 ^ 6)
		for j := 0; j < 64; j++ {
			if (word & (1 << uint64(j))) != 0 {
				fmt.Fprintf(&buf, "%d ", i*64+j)
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func main() {
	var m bitMap
	m.Add(64)
	m.Add(128)
	m.Add(10)
	m.Add(11)
	m.Add(66)

	fmt.Println(m.Has(10))
	fmt.Println(m.Has(11))
	fmt.Println(m.Has(64))
	fmt.Println(m.Has(63))
	fmt.Println(m.Has(128))

	fmt.Println(m.String())
}
