# Bag [![GoDoc](https://godoc.org/github.com/GopherML/bag?status.svg)](https://godoc.org/github.com/GopherML/bag) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/GopherML/bag)](https://goreportcard.com/report/github.com/GopherML/bag) ![Go Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
**Bag** is a powerful yet user-friendly `bag of words` (`BoW`) implementation written in Go, leveraging a Naive Bayes classifier for efficient text analysis. It functions both as a library that can be seamlessly integrated into Go code and as an accessible command line tool. This dual functionality allows users to leverage bag of words capabilities directly from the command line, making it accessible from any programming language. The implementation supports a file format that facilitates using bag of words functionality as code, designed for ease of use and flexible integration in various environments.

![billboard](https://github.com/GopherML/bag/blob/main/bag-billboard.png?raw=true "Bag billboard")

## What is Bag of Words (BoW)?
The `bag of words` (`BoW`) model is a fundamental text representation technique in `natural language processing` (`NLP`). In this model, a text (such as a sentence or a document) is represented as an unordered collection of words, disregarding grammar and word order but keeping multiplicity. The key idea is to create a vocabulary of all the unique words in the text corpus and then represent each text by a vector of word frequencies or binary indicators. This vector indicates the presence or absence, or frequency, of each word from the vocabulary within the text. The `BoW` model is widely used for text classification tasks, including `sentiment analysis`, due to its simplicity and effectiveness in capturing word occurrences.

## Demo
<img src="https://github.com/GopherML/bag/blob/main/demo/demo.gif?raw=true" />

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
- [X] Text normalization added to inbound text processing
- [X] CLI utility

## Long term goals
- [ ] Generated model as MMAP file
## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/itsmontoya"><img src="https://avatars.githubusercontent.com/u/928954?v=4?s=100" width="100px;" alt="Josh Montoya"/><br /><sub><b>Josh Montoya</b></sub></a><br /><a href="https://github.com/GopherML/bag/commits?author=itsmontoya" title="Code">💻</a> <a href="https://github.com/GopherML/bag/commits?author=itsmontoya" title="Documentation">📖</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://mattstay.com"><img src="https://avatars.githubusercontent.com/u/414740?v=4?s=100" width="100px;" alt="Matt Stay"/><br /><sub><b>Matt Stay</b></sub></a><br /><a href="#design-matthew-stay" title="Design">🎨</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://twitter.com/chewxy"><img src="https://avatars.githubusercontent.com/u/471890?v=4?s=100" width="100px;" alt="Chewxy"/><br /><sub><b>Chewxy</b></sub></a><br /><a href="https://github.com/GopherML/bag/commits?author=chewxy" title="Tests">⚠️</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://jbmuir.github.io"><img src="https://avatars.githubusercontent.com/u/18120453?v=4?s=100" width="100px;" alt="Jack Muir"/><br /><sub><b>Jack Muir</b></sub></a><br /><a href="https://github.com/GopherML/bag/commits?author=jbmuir" title="Tests">⚠️</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!