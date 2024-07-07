package bag

import (
	"fmt"
	"testing"
)

var (
	exampleBag     *Bag
	exampleResults Results
)

func TestBag_GetResults(t *testing.T) {
	positiveNegative := []Sample{
		{Input: "I love this product, it is amazing!", Label: "positive"},
		{Input: "This is the worst thing ever.", Label: "negative"},
		{Input: "I am very happy with this.", Label: "positive"},
		{Input: "I hate this so much.", Label: "negative"},
		{Input: "Not good", Label: "negative"},
		{Input: "Very good", Label: "negative"},
	}

	yesNo := []Sample{
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

	type fields struct {
		t TrainingSet
	}

	type args struct {
		in string
	}

	type testcase struct {
		name      string
		fields    fields
		args      args
		wantMatch string
	}

	tests := []testcase{
		{
			name: "positive",
			fields: fields{
				t: TrainingSet{
					Samples: positiveNegative,
				},
			},
			args: args{
				in: "I am very happy with this product.",
			},
			wantMatch: "positive",
		},
		{
			name: "negative",
			fields: fields{
				t: TrainingSet{
					Samples: positiveNegative,
				},
			},
			args: args{
				in: "This is the worst purchase I have ever made.",
			},
			wantMatch: "negative",
		},
		{
			name: "yes",
			fields: fields{
				t: TrainingSet{
					Samples: yesNo,
				},
			},
			args: args{
				in: "Oh yes.",
			},
			wantMatch: "yes",
		},
		{
			name: "no",
			fields: fields{
				t: TrainingSet{
					Config: Config{
						NGramSize: 1,
					},
					Samples: yesNo,
				},
			},
			args: args{
				in: "Oh no.",
			},
			wantMatch: "no",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewFromTrainingSet(tt.fields.t)
			gotR := b.GetResults(tt.args.in).GetHighestProbability()
			if gotR != tt.wantMatch {
				t.Errorf("Bag.GetResults() = %v, want %v", gotR, tt.wantMatch)
			}
		})
	}
}

func ExampleNew() {
	var cfg Config
	// Initialize with default values
	exampleBag = New(cfg)
}

func ExampleNewFromTrainingSet() {
	var t TrainingSet
	t.Samples = []Sample{
		{Input: "I love this product, it is amazing!", Label: "positive"},
		{Input: "This is the worst thing ever.", Label: "negative"},
		{Input: "I am very happy with this.", Label: "positive"},
		{Input: "I hate this so much.", Label: "negative"},
		{Input: "Not good", Label: "negative"},
		{Input: "Very good", Label: "negative"},
	}
	// Initialize with default values
	exampleBag = NewFromTrainingSet(t)
}

func ExampleBag_Train() {
	exampleBag.Train("I love this product, it is amazing!", "positive")
	exampleBag.Train("This is the worst thing ever.", "negative")
	exampleBag.Train("I am very happy with this.", "positive")
	exampleBag.Train("I hate this so much.", "negative")
	exampleBag.Train("Not good", "negative")
	exampleBag.Train("Very good", "negative")
}

func ExampleBag_GetResults() {
	exampleResults = exampleBag.GetResults("I am very happy with this product.")
	fmt.Println("Collection of results", exampleResults)
}

func ExampleResults_GetHighestProbability() {
	match := exampleResults.GetHighestProbability()
	fmt.Println("Highest probability", match)
}
