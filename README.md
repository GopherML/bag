# Bag
Bag is a `bag of words` implementation written in Go. This is currently a library that can be used within Go code. The goal is to provide a file format that offers bag of words as code. In this future, this will be usable as a command line tool which can be called from any language.

## What is Bag of Words (BOW)?
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