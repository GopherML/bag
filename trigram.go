package bag

import (
	"bytes"
)

func toTrigrams(in string) (ts []Trigram) {
	var t Trigram
	toWords(in, func(word string) {
		t = t.Append(word)
		if !t.IsFull() {
			return
		}

		ts = append(ts, t)
	})

	if !t.IsFull() && !t.IsZero() {
		ts = append(ts, t)
	}

	return
}

type Trigram [3]string

func (t *Trigram) Append(str string) (out Trigram) {
	out[0] = t[1]
	out[1] = t[2]
	out[2] = str
	return
}

func (t Trigram) String() (out string) {
	buf := bytes.NewBuffer(nil)
	t.iterate(func(word string) {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}

		buf.WriteString(word)
	})

	return buf.String()
}

func (t Trigram) IsZero() bool {
	return len(t[2]) == 0
}

func (t Trigram) IsFull() bool {
	return len(t[0]) > 0
}

func (t Trigram) iterate(fn func(word string)) {
	for i := 0; i < 3; i++ {
		str := t[i]
		if len(str) == 0 {
			continue
		}

		fn(str)
	}
}
