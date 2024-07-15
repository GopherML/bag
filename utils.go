package bag

import (
	"bytes"
	"unicode"
)

func toWords(in string, onWord func(string)) {
	buf := bytes.NewBuffer(nil)
	for _, char := range in {
		switch {
		case unicode.IsLetter(char):
			char = unicode.ToLower(char)
			buf.WriteRune(char)
		case unicode.IsSpace(char):
			onWord(buf.String())
			buf.Reset()
		}
	}

	if buf.Len() > 0 {
		onWord(buf.String())
		buf.Reset()
	}
}

func toCharacters(in string, onChar func(rune)) {
	for _, char := range in {
		switch {
		case unicode.IsLetter(char):
			char = unicode.ToLower(char)
			onChar(char)
		case unicode.IsSpace(char):
			onChar(char)
		}
	}
}
