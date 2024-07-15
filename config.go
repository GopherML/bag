package bag

import "fmt"

const (
	// DefaultNGramSize is set to 3 (trigram)
	DefaultNGramSize = 3
	// DefaultNGramType is set to word
	DefaultNGramType = "word"
	// DefaultSmoothingParameter is set to 1 (common Laplace smoothing value)
	DefaultSmoothingParameter = 1
)

type Config struct {
	// NGramSize represents the NGram size (unigram, bigram, trigram, etc - default is trigram)
	NGramSize int `yaml:"ngram-size"`
	// NGramType represents the NGram type (word or character - default is word)
	NGramType string `yaml:"ngram-type"`
	// SmoothingParameter represents the smoothing value used for the Laplace Smoothing (default is 1)
	SmoothingParameter float64 `yaml:"smoothing-parameter"`
}

func (c *Config) Validate() (err error) {
	c.fill()

	switch c.NGramType {
	case "word":
	case "character":
	case "":
	default:
		return fmt.Errorf("invalid ngram-type, <%s> is not supported", c.NGramType)
	}

	return
}

func (c *Config) fill() {
	if c.NGramSize == 0 {
		// NGramSize doesn't exist, set to default
		c.NGramSize = DefaultNGramSize
	}

	if c.SmoothingParameter == 0 {
		// SmoothingParameter doesn't exist, set to default
		c.SmoothingParameter = DefaultSmoothingParameter
	}

	if c.NGramType == "" {
		// NGramType doesn't exist, set to default
		c.NGramType = DefaultNGramType
	}
}
