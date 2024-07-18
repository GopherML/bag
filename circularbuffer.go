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

// Shift will add an item to the end of the circular buffer,
// if the buffer is full - it will pop an item from the front
func (c *circularBuffer[T]) Shift(item T) (popped T) {
	// Get oldest buffer item as popped value
	popped = c.s[c.end]
	// Replace oldest position with new item
	c.s[c.end] = item

	// Increment oldest position and check to see if the oldest position exceeds capacity
	if c.end++; c.end >= c.cap {
		// Oldest position exceeds capacity, set to 0
		c.end = 0
	}

	// Check to see if length is less than capaicity
	if c.len < c.cap {
		// Length is not at capacity, increment
		c.len++
		// Increment start value and check to see if new value exceeds capacity
	} else if c.start++; c.start >= c.cap {
		// New start value exceeds start capacity, set to 0
		c.start = 0
	}

	return
}

// ForEach will iterate through the buffer items
func (c *circularBuffer[T]) ForEach(fn func(t T) (end bool)) (ended bool) {
	// First index is at starting position
	index := c.start
	// Iterate for the length of the buffer
	for i := 0; i < c.len; i++ {
		// Get item at current index
		item := c.s[index]
		// Pass item to func
		if fn(item) {
			// Func returned break boolean as true, return true
			return true
		}

		// Increment index and see if index exceeds length
		if index++; index >= c.len {
			// Index exceeds length, set to 0
			index = 0
		}
	}

	return
}

// Len will return the length of a circular buffer
func (c *circularBuffer[T]) Len() int {
	return c.len
}
