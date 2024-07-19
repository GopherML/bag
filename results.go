package bag

import "math"

// Results (by label) represents the probability of a processed input matching each of the possible labels (classifications)
type Results map[string]float64

func (r Results) GetHighestProbability() (match string) {
	// Since probability values can be negative, initialize to negative infinity
	max := math.Inf(-1)
	// Iterate through probability results
	for label, prob := range r {
		// Check to see if the current probability is higher than the max
		if prob > max {
			// Current probability is higher
			// Set max as the current probability
			max = prob
			// Set match as the current label
			match = label
		}
	}

	return
}
