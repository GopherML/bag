package bag

import (
	"bytes"
	"unicode"
)

// toWords will split a string into words
//   - whitespace is omitted
//   - puncation is omitted
//   - all characters repeating more than 2 times will be truncated to 2
//   - all characters are lowercased
func toWords(in string, onWord func(string)) {
	// Buffer is used to write characters while building words
	buf := bytes.NewBuffer(nil)
	// Circular buffer is used to look back on previous characters
	c := newCircularBuffer[rune](2)
	for _, char := range in {
		switch {
		case unicode.IsLetter(char):
			char = unicode.ToLower(char)
			// Create filter function for letter repetition
			isMatch := func(r rune) (end bool) {
				return r != char
			}

			// If length is less than two, or letter changes encountered
			if c.Len() < 2 || c.ForEach(isMatch) {
				// Write character to buffer
				buf.WriteRune(char)
				// Shift circular buffer with new character
				c.Shift(char)
			}

		case unicode.IsSpace(char):
			// Space encountered, call onWord with word
			onWord(buf.String())
			buf.Reset()
		}
	}

	// Check to see if the buffer is not empty
	if buf.Len() > 0 {
		// Buffer is not empty, call onWord with word
		onWord(buf.String())
		buf.Reset()
	}
}

// toCharacters will split a string into characters
//   - whitespace is included
//   - puncation is omitted
//   - all characters are lowercased
func toCharacters(in string, onChar func(rune)) {
	// Iterate through all input string runes
	for _, char := range in {
		switch {
		case unicode.IsLetter(char):
			// Lowercase character
			char = unicode.ToLower(char)
			// Call onChar with characters
			onChar(char)
		case unicode.IsSpace(char):
			// Call onChar with whitespace
			onChar(char)
		}
	}
}
