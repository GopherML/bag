package bag

import (
	"bytes"
)

func toNGrams(in string, size int) (ns []NGram) {
	n := make(NGram, size)
	toWords(in, func(word string) {
		n = n.Append(word)
		if !n.IsFull() {
			return
		}

		ns = append(ns, n)
	})

	if !n.IsFull() && !n.IsZero() {
		ns = append(ns, n)
	}

	return
}

type NGram []string

func (n NGram) Append(str string) (out NGram) {
	out = make(NGram, len(n))
	for i := 1; i < len(n); i++ {
		out[i-1] = n[i]
	}

	out[len(n)-1] = str
	return
}

func (n NGram) String() (out string) {
	buf := bytes.NewBuffer(nil)
	n.iterate(func(word string) {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}

		buf.WriteString(word)
	})

	return buf.String()
}

func (n NGram) IsZero() bool {
	return len(n[len(n)-1]) == 0
}

func (n NGram) IsFull() bool {
	return len(n[0]) > 0
}

func (n NGram) iterate(fn func(word string)) {
	for _, str := range n {
		if len(str) == 0 {
			continue
		}

		fn(str)
	}
}
