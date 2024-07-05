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
	}

	// Test documents
	tests := []training{
		{Input: "I am very happy with this product.", Label: "positive"},
		{Input: "This is the worst purchase I have ever made.", Label: "negative"},
	}

	testBag(t, trainings, tests)
}

func TestBag_yesno(t *testing.T) {
	trainings := []training{
		{Input: "Yes", Label: "positive"},
		{Input: "Yeah", Label: "positive"},
		{Input: "Yep", Label: "positive"},
		{Input: "Yup", Label: "positive"},
		{Input: "Yea", Label: "positive"},
		{Input: "Sure", Label: "positive"},
		{Input: "Absolutely", Label: "positive"},
		{Input: "Definitely", Label: "positive"},
		{Input: "Of course", Label: "positive"},
		{Input: "For sure", Label: "positive"},
		{Input: "Indeed", Label: "positive"},
		{Input: "Affirmative", Label: "positive"},
		{Input: "Roger", Label: "positive"},
		{Input: "Totally", Label: "positive"},
		{Input: "Certainly", Label: "positive"},
		{Input: "Without a doubt", Label: "positive"},
		{Input: "You bet", Label: "positive"},
		{Input: "Uh-huh", Label: "positive"},
		{Input: "Right on", Label: "positive"},
		{Input: "Cool", Label: "positive"},
		{Input: "Okie dokie", Label: "positive"},
		{Input: "Aye", Label: "positive"},
		{Input: "Yass", Label: "positive"},
		{Input: "Fo sho", Label: "positive"},
		{Input: "Bet", Label: "positive"},
		{Input: "10-4", Label: "positive"},
		{Input: "No", Label: "negative"},
		{Input: "Nope", Label: "negative"},
		{Input: "Nah", Label: "negative"},
		{Input: "Nuh-uh", Label: "negative"},
		{Input: "No way", Label: "negative"},
		{Input: "Not at all", Label: "negative"},
		{Input: "Negative", Label: "negative"},
		{Input: "Not really", Label: "negative"},
		{Input: "I don't think so", Label: "negative"},
		{Input: "Absolutely not", Label: "negative"},
		{Input: "No chance", Label: "negative"},
		{Input: "No way, José", Label: "negative"},
		{Input: "Out of the question", Label: "negative"},
		{Input: "By no means", Label: "negative"},
		{Input: "Under no circumstances", Label: "negative"},
		{Input: "Never", Label: "negative"},
		{Input: "Not in a million years", Label: "negative"},
		{Input: "Not happening", Label: "negative"},
		{Input: "No can do", Label: "negative"},
		{Input: "Not on your life", Label: "negative"},
		{Input: "Hell no", Label: "negative"},
		{Input: "Nah fam", Label: "negative"},
		{Input: "Pass", Label: "negative"},
		{Input: "Hard pass", Label: "negative"},
		{Input: "Nopey dopey", Label: "negative"},
		{Input: "Nix", Label: "negative"},
	}

	// Test documents
	tests := []training{
		{Input: "Yep", Label: "positive"},
		{Input: "Oh yes", Label: "positive"},
		{Input: "Oh no", Label: "negative"},
	}

	testBag(t, trainings, tests)
}

func testBag(_ *testing.T, trainings, tests []training) {
	b := New()
	for _, training := range trainings {
		b.Train(training.Input, training.Label)
	}

	// Analyze sentiment
	for _, test := range tests {
		labels := b.GetSeniment(test.Input)
		fmt.Printf("Document: %s\nSentiment: %v\n\n", test.Input, labels)
	}
}
