package bag

import "bytes"

// toNGrams will convert inbound data to an nGram of provided size
func toNGrams(in string, size int) (ns []string) {
	// Initialize nGram with a provided size
	n := make(nGram, size)
	// Iterate inbound data as words
	toWords(in, func(word string) {
		// Append word to nGram
		n = n.Append(word)
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
type nGram []string

// Append will append a given string to an nGram and output the new value
// Note: The original nGram is NOT modified
func (n nGram) Append(str string) (out nGram) {
	// Initialize new nGram with the same size as the original nGram
	out = make(nGram, len(n))
	// Iterate through original nGram, starting at index 1
	for i := 1; i < len(n); i++ {
		// Set the value of the current original nGram index as the value for the previous index for the output nGram
		out[i-1] = n[i]
	}

	// Set the last value of the output nGram as the input string
	out[len(n)-1] = str
	return
}

// String will convert the nGram contents to a string
func (n nGram) String() (out string) {
	// Initialize buffer
	buf := bytes.NewBuffer(nil)
	// Iterate through nGram values
	n.iterate(func(value string) {
		if buf.Len() > 0 {
			// Buffer is not empty, prefix the iterating value with a space
			buf.WriteByte(' ')
		}

		// Write value to buffer
		buf.WriteString(value)
	})

	// Return buffer as string
	return buf.String()
}

// IsZero returns whether or not the nGram is empty
func (n nGram) IsZero() bool {
	// Return result of if the value in the last position is empty
	return len(n[len(n)-1]) == 0
}

// IsFull returns whether or not the nGram is full
func (n nGram) IsFull() bool {
	// Return result of if the value in the first position is populated
	return len(n[0]) > 0
}

// iterate will iterate through the nGram values
func (n nGram) iterate(fn func(word string)) {
	// Iterate through nGram values
	for _, str := range n {
		// Check if value is empty
		if len(str) == 0 {
			// Value is empty, continue
			continue
		}

		// Value is populated, pass to provided func
		fn(str)
	}
}
