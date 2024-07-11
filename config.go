package bag

const (
	// DefaultNGramSize is set to 3 (trigram)
	DefaultNGramSize = 3
	// DefaultSmoothingParameter is set to 1 (common Laplace smoothing value)
	DefaultSmoothingParameter = 1
)

type Config struct {
	// NGramSize represents the NGram size (unigram, bigram, trigram, etc - default is trigram)
	NGramSize int `yaml:"ngram-size"`
	// SmoothingParameter represents the smoothing value used for the Laplace Smoothing (default is 1)
	SmoothingParameter int `yaml:"smoothing-parameter"`
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
}
