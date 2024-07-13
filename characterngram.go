package bag

import "bytes"

// tocharacterNGrams will convert inbound data to an characterNGram of provided size
func tocharacterNGrams(in string, size int) (ns []string) {
	// Initialize characterNGram with a provided size
	n := make(characterNGram, size)
	// Iterate inbound data as words
	toCharacters(in, func(char rune) {
		// Append word to characterNGram
		n = n.Append(char)
		if !n.IsFull() {
			// characterNGram is not full - we do not want to append yet, return
			return
		}

		// Append current characterNGram to characterNGrams slice
		ns = append(ns, n.String())
	})

	if !n.IsFull() && !n.IsZero() {
		// The characterNGram is not full, so we haven't appended yet
		// The characterNGram is not empty, so we have something to append
		// Append current characterNGram to characterNGrams slice
		ns = append(ns, n.String())
	}

	return
}

// characterNGram represents an characterNGram (variable sized)
type characterNGram []rune

// Append will append a given string to an characterNGram and output the new value
// Note: The original characterNGram is NOT modified
func (n characterNGram) Append(char rune) (out characterNGram) {
	// Initialize new characterNGram with the same size as the original characterNGram
	out = make(characterNGram, len(n))
	// Iterate through original characterNGram, starting at index 1
	for i := 1; i < len(n); i++ {
		// Set the value of the current original characterNGram index as the value for the previous index for the output characterNGram
		out[i-1] = n[i]
	}

	// Set the last value of the output characterNGram as the input string
	out[len(n)-1] = char
	return
}

// String will convert the characterNGram contents to a string
func (n characterNGram) String() (out string) {
	// Initialize buffer
	buf := bytes.NewBuffer(nil)
	// Iterate through characterNGram values
	n.iterate(func(char rune) {
		// Write value to buffer
		buf.WriteRune(char)
	})

	// Return buffer as string
	return buf.String()
}

// IsZero returns whether or not the characterNGram is empty
func (n characterNGram) IsZero() bool {
	// Return result of if the value in the last position is empty
	return n[len(n)-1] == 0
}

// IsFull returns whether or not the characterNGram is full
func (n characterNGram) IsFull() bool {
	// Return result of if the value in the first position is populated
	return n[0] != 0
}

// iterate will iterate through the characterNGram values
func (n characterNGram) iterate(fn func(char rune)) {
	// Iterate through characterNGram values
	for _, char := range n {
		// Check if value is empty
		if char == 0 {
			// Value is empty, continue
			continue
		}

		// Value is populated, pass to provided func
		fn(char)
	}
}
