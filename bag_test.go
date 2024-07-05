package bag

import (
	"fmt"
	"testing"
)

type training struct {
	Input string
	Label string
}

func TestBag(t *testing.T) {
	trainings := []training{
		{Input: "I love this product, it is amazing!", Label: "positive"},
		{Input: "This is the worst thing ever.", Label: "negative"},
		{Input: "I am very happy with this.", Label: "positive"},
		{Input: "I hate this so much.", Label: "negative"},
		{Input: "Not good", Label: "negative"},
		{Input: "Very good", Label: "negative"},
	}

	// Test documents
	tests := []training{
		{Input: "I am very happy with this product.", Label: "positive"},
		{Input: "This is the worst purchase I have ever made.", Label: "negative"},
	}

	testBag(t, trainings, tests, Config{})
}

func TestBag_yesno(t *testing.T) {
	trainings := []training{
		{Input: "Yes", Label: "yes"},
		{Input: "Yeah", Label: "yes"},
		{Input: "Yep", Label: "yes"},
		{Input: "Yup", Label: "yes"},
		{Input: "Yea", Label: "yes"},
		{Input: "Sure", Label: "yes"},
		{Input: "Absolutely", Label: "yes"},
		{Input: "Definitely", Label: "yes"},
		{Input: "Of course", Label: "yes"},
		{Input: "For sure", Label: "yes"},
		{Input: "Indeed", Label: "yes"},
		{Input: "Affirmative", Label: "yes"},
		{Input: "Roger", Label: "yes"},
		{Input: "Totally", Label: "yes"},
		{Input: "Certainly", Label: "yes"},
		{Input: "Without a doubt", Label: "yes"},
		{Input: "You bet", Label: "yes"},
		{Input: "Uh-huh", Label: "yes"},
		{Input: "Right on", Label: "yes"},
		{Input: "Cool", Label: "yes"},
		{Input: "Okie dokie", Label: "yes"},
		{Input: "Aye", Label: "yes"},
		{Input: "Yass", Label: "yes"},
		{Input: "Fo sho", Label: "yes"},
		{Input: "Bet", Label: "yes"},
		{Input: "10-4", Label: "yes"},
		{Input: "No", Label: "no"},
		{Input: "Nope", Label: "no"},
		{Input: "Nah", Label: "no"},
		{Input: "Nuh-uh", Label: "no"},
		{Input: "No way", Label: "no"},
		{Input: "Not at all", Label: "no"},
		{Input: "no", Label: "no"},
		{Input: "Not really", Label: "no"},
		{Input: "I don't think so", Label: "no"},
		{Input: "Absolutely not", Label: "no"},
		{Input: "No chance", Label: "no"},
		{Input: "No way, José", Label: "no"},
		{Input: "Out of the question", Label: "no"},
		{Input: "By no means", Label: "no"},
		{Input: "Under no circumstances", Label: "no"},
		{Input: "Never", Label: "no"},
		{Input: "Not in a million years", Label: "no"},
		{Input: "Not happening", Label: "no"},
		{Input: "No can do", Label: "no"},
		{Input: "Not on your life", Label: "no"},
		{Input: "Hell no", Label: "no"},
		{Input: "Nah fam", Label: "no"},
		{Input: "Pass", Label: "no"},
		{Input: "Hard pass", Label: "no"},
		{Input: "Nopey dopey", Label: "no"},
		{Input: "Nix", Label: "no"},
	}

	// Test documents
	tests := []training{
		{Input: "Yep", Label: "yes"},
		{Input: "Oh yes", Label: "yes"},
		{Input: "Oh no", Label: "no"},
	}

	testBag(t, trainings, tests, Config{NGramSize: 1})
}

func testBag(t *testing.T, trainings, tests []training, cfg Config) {
	b := New(cfg)
	for _, training := range trainings {
		b.Train(training.Input, training.Label)
	}

	// Analyze sentiment
	for _, test := range tests {
		results := b.GetSentiment(test.Input)
		got := results.GetHighestProbability()
		if got != test.Label {
			t.Fatalf("invalid label, expected <%s> and received <%s>", test.Label, got)
		}

		fmt.Printf("Document: %s\nMatch: %v\nResults: %v\n\n", test.Input, results.GetHighestProbability(), results)
	}
}
