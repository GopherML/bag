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
	positiveNegative := SamplesByLabel{
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

	yesNo := SamplesByLabel{
		"yes": {
			"Yes",
			"Yeah",
			"Yep",
			"Yup",
			"Yea",
			"Sure",
			"Absolutely",
			"Definitely",
			"Of course",
			"For sure",
			"Indeed",
			"Affirmative",
			"Roger",
			"Totally",
			"Certainly",
			"Without a doubt",
			"You bet",
			"Uh-huh",
			"Right on",
			"Cool",
			"Okie dokie",
			"Aye",
			"Yass",
			"Fo sho",
			"Bet",
			"10-4",
		},
		"no": {
			"No",
			"Nope",
			"Nah",
			"Nuh-uh",
			"No way",
			"Not at all",
			"no",
			"Not really",
			"I don't think so",
			"Absolutely not",
			"No chance",
			"No way, José",
			"Out of the question",
			"By no means",
			"Under no circumstances",
			"Never",
			"Not in a million years",
			"Not happening",
			"No can do",
			"Not on your life",
			"Hell no",
			"Nah fam",
			"Pass",
			"Hard pass",
			"Nopey dopey",
			"Nix",
		},
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
