package bag

func newCircularBuffer[T any](capacity int) *circularBuffer[T] {
	var c circularBuffer[T]
	c.cap = capacity
	c.s = make([]T, capacity)
	return &c
}

type circularBuffer[T any] struct {
	start int
	end   int

	len int
	cap int

	s []T
}

func (c *circularBuffer[T]) Shift(item T) (popped T) {
	popped = c.s[c.end]
	c.s[c.end] = item

	c.end++
	if c.len < c.cap {
		c.len++
	} else {
		if c.start++; c.start >= c.cap {
			c.start = 0
		}

	}

	if c.end >= c.cap {
		c.end = 0
	}

	return
}

func (c *circularBuffer[T]) ForEach(fn func(t T) (end bool)) (ended bool) {
	index := c.start
	for i := 0; i < c.len; i++ {
		item := c.s[index]
		if fn(item) {
			return true
		}

		if index++; index >= c.len {
			index = 0
		}
	}

	return
}

func (c *circularBuffer[T]) Len() int {
	return c.len
}
