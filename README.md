# Bag
Bag is a bag of words implementation written in Go. This is currently a library that can be used within Go code. The goal is to provide a file format that offers bag of words as code. In this future, this will be usable as a command line tool which can be called from any language.

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