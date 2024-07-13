# Bag [![GoDoc](https://godoc.org/github.com/itsmontoya/bag?status.svg)](https://godoc.org/github.com/itsmontoya/bag) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/itsmontoya/bag)](https://goreportcard.com/report/github.com/itsmontoya/bag) ![Go Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
Bag is a `bag of words` (`BoW`) implementation written in Go, utilizing a Naive Bayes classifier. Currently, it functions as a library that can be integrated into Go code. The goal is to offer a file format that provides bag of words functionality as code. In the future, it will be usable as a command line tool, allowing it to be called from any programming language.

## What is Bag of Words (BoW)?
The `bag of words` (`BoW`) model is a fundamental text representation technique in `natural language processing` (`NLP`). In this model, a text (such as a sentence or a document) is represented as an unordered collection of words, disregarding grammar and word order but keeping multiplicity. The key idea is to create a vocabulary of all the unique words in the text corpus and then represent each text by a vector of word frequencies or binary indicators. This vector indicates the presence or absence, or frequency, of each word from the vocabulary within the text. The `BoW` model is widely used for text classification tasks, including `sentiment analysis`, due to its simplicity and effectiveness in capturing word occurrences.

## Examples
### New
```go
func ExampleNew() {
	var cfg Config
	// Initialize with default values
	exampleBag = New(cfg)
}
```

### NewFromTrainingSet
```go
func ExampleNewFromTrainingSet() {
	var t TrainingSet
	t.Samples = SamplesByLabel{
		"positive": {
			"I love this product, it is amazing!",
			"I am very happy with this.",
			"Very good",
		},

		"negative": {
			"This is the worst thing ever.",
			"I hate this so much.",
			"Not good",
		},
	}

	// Initialize with default values
	exampleBag = NewFromTrainingSet(t)
}
```

### Bag.Train
```go
func ExampleBag_Train() {
	exampleBag.Train("I love this product, it is amazing!", "positive")
	exampleBag.Train("This is the worst thing ever.", "negative")
	exampleBag.Train("I am very happy with this.", "positive")
	exampleBag.Train("I hate this so much.", "negative")
	exampleBag.Train("Not good", "negative")
	exampleBag.Train("Very good", "negative")
}
```

### Bag.GetResults
```go
func ExampleBag_GetResults() {
	exampleResults = exampleBag.GetResults("I am very happy with this product.")
	fmt.Println("Collection of results", exampleResults)
}
```

### Results.GetHighestProbability
```go
func ExampleResults_GetHighestProbability() {
	match := exampleResults.GetHighestProbability()
	fmt.Println("Highest probability", match)
}
```

### TrainingSet File
```yaml
config:
  ngram-size: 1
samples:
  yes:
    - "yes"
    - "Yeah"
    - "Yep"

  no:
    - "No"
    - "Nope"
    - "Nah"

# Note: This training set is short for the sake of README filesize,
# please look in the examples directory for more complete examples
```

## Road to v1.0.0

- [X] Working implementation as Go library
- [X] Training sets
- [X] Support Character NGrams
- [ ] Text normalization added to inbound text processing
- [X] CLI utility

## Long term goals
- [ ] Generated model as MMAP file