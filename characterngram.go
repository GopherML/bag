package bag

import "bytes"

// toCharacterNGrams will convert inbound data to an characterNGram of provided size
func toCharacterNGrams(in string, size int) (ns []string) {
	var c characterNGram
	// Initialize characterNGram with a provided size
	c.circularBuffer = newCircularBuffer[rune](size)
	// Iterate inbound data as words
	toCharacters(in, func(char rune) {
		// Append word to characterNGram
		c.Shift(char)
		if !c.IsFull() {
			// characterNGram is not full - we do not want to append yet, return
			return
		}

		// Append current characterNGram to characterNGrams slice
		ns = append(ns, c.String())
	})

	if !c.IsFull() && !c.IsZero() {
		// The characterNGram is not full, so we haven't appended yet
		// The characterNGram is not empty, so we have something to append
		// Append current characterNGram to characterNGrams slice
		ns = append(ns, c.String())
	}

	return
}

// characterNGram represents an characterNGram (variable sized)
type characterNGram struct {
	*circularBuffer[rune]
}

// String will convert the characterNGram contents to a string
func (n characterNGram) String() (out string) {
	// Initialize buffer
	buf := bytes.NewBuffer(nil)
	// Iterate through characterNGram values
	n.ForEach(func(char rune) (end bool) {
		// Write value to buffer
		buf.WriteRune(char)
		return
	})

	// Return buffer as string
	return buf.String()
}

// IsZero returns whether or not the characterNGram is empty
func (n characterNGram) IsZero() bool {
	// Return result of if the value in the first position is populated
	return n.s[0] == 0
}

// IsFull returns whether or not the characterNGram is full
func (n characterNGram) IsFull() bool {
	// Return result of if the value in the last position is empty
	return n.s[len(n.s)-1] > 0
}
