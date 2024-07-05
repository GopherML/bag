package bag

const (
	DefaultNGramSize          = 3
	DefaultSmoothingParameter = 1
)

type Config struct {
	NGramSize          int
	SmoothingParameter int
}

func (c *Config) fill() {
	if c.NGramSize == 0 {
		c.NGramSize = DefaultNGramSize
	}

	if c.SmoothingParameter == 0 {
		c.SmoothingParameter = DefaultSmoothingParameter
	}
}
