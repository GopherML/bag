package bag

import "math"

type Results map[string]float64

func (r Results) GetHighestProbability() (match string) {
	// Since probability values can be negative, initialize to negative infinity
	max := math.Inf(-1)
	for label, prob := range r {
		if prob > max {
			max = prob
			match = label
		}
	}

	return
}
