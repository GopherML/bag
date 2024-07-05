package bag

import "math"

func New(c Config) *Bag {
	var b Bag
	c.fill()
	b.c = c
	b.labels = map[string]Vocabulary{}
	b.countByLabel = map[string]int{}
	return &b
}

type Bag struct {
	c Config

	labels map[string]Vocabulary

	// Count of trained documents by label
	countByLabel map[string]int
	// Total count of trained documents
	totalCount int
}

func (b *Bag) GetSentiment(in string) (r Results) {
	ns := toNGrams(in, b.c.NGramSize)
	r = make(Results, len(b.labels))

	for label, vocab := range b.labels {
		r[label] = b.getProbability(ns, label, vocab)
	}

	return
}

func (b *Bag) Train(in, label string) {
	// Convert inbound data to a slice of NGrams
	ns := toNGrams(in, b.c.NGramSize)
	// Get vocabulary for a provided label, if the vocabulary doesn't exist, it will be created)
	v := b.getOrCreateVocabulary(label)
	// Iterate through NGrams
	for _, n := range ns {
		// Increment the vocabulary value for the current NGram
		v[n.String()]++
	}

	// Increment count of trained documents for the provided label
	b.countByLabel[label]++
	// Increment total count of trained documents
	b.totalCount++
}

func (b *Bag) getProbability(ns []NGram, label string, vocab Vocabulary) (probability float64) {
	// Set initial probability value as the prior probability value
	probability = b.getPriorProbability(label)
	// Get the current counts by label (to be used by Laplace smoothing during for-loop)
	countsByLabel := float64(b.countByLabel[label] + len(vocab))
	// Iterate through NGrams
	for _, n := range ns {
		// Utilize Laplace smoothing to improve our results when an ngram isn't found within the trained dataset
		// Likelihood with Laplace smoothing
		count := float64(vocab[n.String()] + b.c.SmoothingParameter)
		// Add logarithmic result of count (plus )
		probability += math.Log(count / countsByLabel)
	}

	return
}

func (b *Bag) getPriorProbability(label string) (probability float64) {
	count := float64(b.countByLabel[label])
	total := float64(b.totalCount)
	return math.Log(count / total)
}

func (b *Bag) getOrCreateVocabulary(label string) (v Vocabulary) {
	var ok bool
	v, ok = b.labels[label]
	if !ok {
		v = make(Vocabulary)
		b.labels[label] = v
	}

	return
}
