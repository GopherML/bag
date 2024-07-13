package bag

import "math"

func New(c Config) (out *Bag, err error) {
	// Validate Config
	if err = c.Validate(); err != nil {
		return
	}

	var b Bag
	b.c = c
	b.vocabByLabel = map[string]Vocabulary{}
	b.countByLabel = map[string]int{}
	out = &b
	return
}

func NewFromTrainingSet(t TrainingSet) (b *Bag, err error) {
	if b, err = New(t.Config); err != nil {
		return
	}

	for label, samples := range t.Samples {
		for _, sample := range samples {
			b.Train(sample, label)
		}
	}

	return
}

type Bag struct {
	// Configuration values
	c Config
	// Vocabulary sets by label
	vocabByLabel map[string]Vocabulary
	// Count of trained documents by label
	countByLabel map[string]int
	// Total count of trained documents
	totalCount int
}

func (b *Bag) GetResults(in string) (r Results) {
	// Convert inbound data to NGrams
	ns := toNGrams(in, b.c.NGramSize)
	// Initialize results with the same size as the current number of vocabulary labels
	r = make(Results, len(b.vocabByLabel))
	// Iterate through vocabulary sets by label
	for label, vocab := range b.vocabByLabel {
		// Set probability value for iterating label
		r[label] = b.getProbability(ns, label, vocab)
	}

	return
}
func (b *Bag) toNGrams(in string) (ns []string) {
	if b.c.NGramType == "word" {
		return toNGrams(in, b.c.NGramSize)
	}

	return tocharacterNGrams(in, b.c.NGramSize)
}

func (b *Bag) Train(in, label string) {
	// Convert inbound data to a slice of NGrams
	ns := b.toNGrams(in)
	// Get vocabulary for a provided label, if the vocabulary doesn't exist, it will be created)
	v := b.getOrCreateVocabulary(label)
	// Iterate through NGrams
	for _, n := range ns {
		// Increment the vocabulary value for the current NGram
		v[n]++
	}

	// Increment count of trained documents for the provided label
	b.countByLabel[label]++
	// Increment total count of trained documents
	b.totalCount++
}

// getProbability uses a Naive Bayes classifier to determine probability for a given label
func (b *Bag) getProbability(ns []string, label string, vocab Vocabulary) (probability float64) {
	// Set initial probability value as the prior probability value
	probability = b.getPriorProbability(label)
	// Get the current counts by label (to be used by Laplace smoothing during for-loop)
	countsByLabel := float64(b.countByLabel[label] + len(vocab))
	// Iterate through NGrams
	for _, n := range ns {
		// Utilize Laplace smoothing to improve our results when an ngram isn't found within the trained dataset
		// Likelihood with Laplace smoothing
		count := float64(vocab[n] + b.c.SmoothingParameter)
		// Add logarithmic result of count (plus )
		probability += math.Log(count / countsByLabel)
	}

	return
}

func (b *Bag) getPriorProbability(label string) (probability float64) {
	count := float64(b.countByLabel[label])
	total := float64(b.totalCount)
	// Get the logarithmic value of count divided by total count
	return math.Log(count / total)
}

func (b *Bag) getOrCreateVocabulary(label string) (v Vocabulary) {
	var ok bool
	v, ok = b.vocabByLabel[label]
	// Check if vocabulary set does not exist for the provided label
	if !ok {
		// Create new vocabulary set
		v = make(Vocabulary)
		// Set vocabulary set by label as newly created value
		b.vocabByLabel[label] = v
	}

	return
}
