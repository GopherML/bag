package bag

import (
	"bytes"
)

// toNGrams will convert inbound data to an nGram of provided size
func toNGrams(in string, size int) (ns []string) {
	var n nGram
	n.circularBuffer = newCircularBuffer[string](size)
	// Iterate inbound data as words
	toWords(in, func(word string) {
		// Append word to nGram
		n.Shift(word)
		if !n.IsFull() {
			// NGram is not full - we do not want to append yet, return
			return
		}

		// Append current nGram to nGrams slice
		ns = append(ns, n.String())
	})

	if !n.IsFull() && !n.IsZero() {
		// The nGram is not full, so we haven't appended yet
		// The nGram is not empty, so we have something to append
		// Append current nGram to nGrams slice
		ns = append(ns, n.String())
	}

	return
}

// nGram represents an N-Gram (variable sized)
type nGram struct {
	*circularBuffer[string]
}

// String will convert the nGram contents to a string
func (n *nGram) String() (out string) {
	// Initialize buffer
	buf := bytes.NewBuffer(nil)
	// Iterate through nGram values
	n.ForEach(func(value string) (end bool) {
		if buf.Len() > 0 {
			// Buffer is not empty, prefix the iterating value with a space
			buf.WriteByte(' ')
		}

		// Write value to buffer
		buf.WriteString(value)
		return
	})

	// Return buffer as string
	return buf.String()
}

// IsZero returns whether or not the nGram is empty
func (n nGram) IsZero() bool {
	// Return result of if the value in the first position is populated
	return len(n.s[0]) == 0
}

// IsFull returns whether or not the nGram is full
func (n nGram) IsFull() bool {
	// Return result of if the value in the last position is empty
	return len(n.s[len(n.s)-1]) > 0
}
