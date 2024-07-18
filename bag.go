package bag

import "math"

// New will initialize and return a new Bag with a provided configuration
func New(c Config) (out *Bag, err error) {
	// Validate Config
	if err = c.Validate(); err != nil {
		return
	}

	var b Bag
	b.c = c
	b.vocabByLabel = map[string]Vocabulary{}
	b.documentCountByLabel = map[string]int{}
	out = &b
	return
}

// NewFromTrainingSet will initialize and return a new pre-trained Bag from a provided training set
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

// NewFromTrainingSetFile will initialize and return a new pre-trained Bag from a provided training set filepath
func NewFromTrainingSetFile(filepath string) (b *Bag, err error) {
	var t TrainingSet
	if t, err = makeTrainingSetFromFile(filepath); err != nil {
		return
	}

	return NewFromTrainingSet(t)
}

// Bag represents a bag of words (BoW) model
type Bag struct {
	// Configuration values
	c Config
	// Vocabulary sets by label
	vocabByLabel map[string]Vocabulary
	// Count of trained documents by label
	documentCountByLabel map[string]int
	// Total count of trained documents
	totalDocumentCount int
}

// GetResults will return the classification results for a given input string
func (b *Bag) GetResults(in string) (r Results) {
	// Convert inbound data to NGrams
	ns := b.toNGrams(in)
	// Initialize results with the same size as the current number of vocabulary labels
	r = make(Results, len(b.vocabByLabel))
	// Iterate through vocabulary sets by label
	for label, vocab := range b.vocabByLabel {
		// Set probability value for iterating label
		r[label] = b.getProbability(ns, label, vocab)
	}

	return
}

// Train will process a given input string and assign it the provided label for training
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

	// Increment model counters
	b.incrementCounts(label)
}

// toNGrams converts the inbound string into n-grams based on the configuration settings
func (b *Bag) toNGrams(in string) (ns []string) {
	if b.c.NGramType == "word" {
		// NGram type is word, use n-grams
		return toNGrams(in, b.c.NGramSize)
	}

	// NGram type is character, use character n-grams
	return toCharacterNGrams(in, b.c.NGramSize)
}

// getProbability uses a Naive Bayes classifier to determine probability for a given label
func (b *Bag) getProbability(ns []string, label string, vocab Vocabulary) (probability float64) {
	// Set initial probability value as the prior probability value
	probability = b.getLogPriorProbability(label)
	// Get the current counts by label (to be used by Laplace smoothing during for-loop)
	countsByLabel := float64(b.documentCountByLabel[label]) + b.c.SmoothingParameter*float64(len(vocab))

	// Iterate through NGrams
	for _, n := range ns {
		// Utilize Laplace smoothing to improve our results when an ngram isn't found within the trained dataset
		// Likelihood with Laplace smoothing
		count := float64(vocab[n]) + b.c.SmoothingParameter
		// Add logarithmic result of count (plus )
		probability += math.Log(count / countsByLabel)
	}

	return
}

// getLogPriorProbability will get the starting probability value for a given label
func (b *Bag) getLogPriorProbability(label string) (probability float64) {
	// Document count for the given label
	countByLabel := float64(b.documentCountByLabel[label])
	// Total document count
	total := float64(b.totalDocumentCount)
	// Get the logarithmic value of count divided by total count
	return math.Log(countByLabel / total)
}

// getOrCreate vocabulary will get a vocabulary set for a given label,
// if the vocabulary doesn't exist - it is created
func (b *Bag) getOrCreateVocabulary(label string) (v Vocabulary) {
	var ok bool
	// Attempt to get vocabulary for the given label
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

// incrementCounts will increment trained documents count globally and by label
func (b *Bag) incrementCounts(label string) {
	// Increment count of trained documents for the provided label
	b.documentCountByLabel[label]++
	// Increment total count of trained documents
	b.totalDocumentCount++
}
