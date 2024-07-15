package bag

import "bytes"

// toNGrams will convert inbound data to an NGram of provided size
func toNGrams(in string, size int) (ns []string) {
	// Initialize NGram with a provided size
	n := make(NGram, size)
	// Iterate inbound data as words
	toWords(in, func(word string) {
		// Append word to NGram
		n = n.Append(word)
		if !n.IsFull() {
			// NGram is not full - we do not want to append yet, return
			return
		}

		// Append current NGram to NGrams slice
		ns = append(ns, n.String())
	})

	if !n.IsFull() && !n.IsZero() {
		// The NGram is not full, so we haven't appended yet
		// The NGram is not empty, so we have something to append
		// Append current NGram to NGrams slice
		ns = append(ns, n.String())
	}

	return
}

// NGram represents an NGram (variable sized)
type NGram []string

// Append will append a given string to an NGram and output the new value
// Note: The original NGram is NOT modified
func (n NGram) Append(str string) (out NGram) {
	// Initialize new NGram with the same size as the original NGram
	out = make(NGram, len(n))
	// Iterate through original NGram, starting at index 1
	for i := 1; i < len(n); i++ {
		// Set the value of the current original NGram index as the value for the previous index for the output NGram
		out[i-1] = n[i]
	}

	// Set the last value of the output NGram as the input string
	out[len(n)-1] = str
	return
}

// String will convert the NGram contents to a string
func (n NGram) String() (out string) {
	// Initialize buffer
	buf := bytes.NewBuffer(nil)
	// Iterate through NGram values
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

// IsZero returns whether or not the NGram is empty
func (n NGram) IsZero() bool {
	// Return result of if the value in the last position is empty
	return len(n[len(n)-1]) == 0
}

// IsFull returns whether or not the NGram is full
func (n NGram) IsFull() bool {
	// Return result of if the value in the first position is populated
	return len(n[0]) > 0
}

// iterate will iterate through the NGram values
func (n NGram) iterate(fn func(word string)) {
	// Iterate through NGram values
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
