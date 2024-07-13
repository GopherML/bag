package bag

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-yaml/yaml"
)

var (
	exampleBag     *Bag
	exampleResults Results
)

var (
	testTrainingYesNo TrainingSet
)

func TestMain(m *testing.M) {
	var (
		f   *os.File
		err error
	)
	if f, err = os.Open("./examples/yes-no-training.yaml"); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err = yaml.NewDecoder(f).Decode(&testTrainingYesNo); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	type args struct {
		c Config
	}

	type teststruct struct {
		name    string
		args    args
		wantErr bool
	}

	tests := []teststruct{
		{
			name: "empty",
			args: args{
				c: Config{},
			},
			wantErr: false,
		},
		{
			name: "invalid ngram type",
			args: args{
				c: Config{
					NGramType: "foobar",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewFromTrainingSet(t *testing.T) {
	type args struct {
		t TrainingSet
	}

	type teststruct struct {
		name    string
		args    args
		wantErr bool
	}

	tests := []teststruct{
		{
			name: "invalid ngram type",
			args: args{
				t: TrainingSet{
					Config: Config{
						NGramType: "foobar",
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFromTrainingSet(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

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
				t: testTrainingYesNo,
			},
			args: args{
				in: "Oh yes.",
			},
			wantMatch: "yes",
		},
		{
			name: "no",
			fields: fields{
				t: testTrainingYesNo,
			},
			args: args{
				in: "Oh no.",
			},
			wantMatch: "no",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := NewFromTrainingSet(tt.fields.t)
			if err != nil {
				t.Fatal(err)
			}

			gotR := b.GetResults(tt.args.in).GetHighestProbability()
			if gotR != tt.wantMatch {
				t.Errorf("Bag.GetResults() = wrong value for <%v>: %v, want %v", tt.args.in, gotR, tt.wantMatch)
				fmt.Printf("%+v\n", b.vocabByLabel)
				fmt.Println(b.GetResults(tt.args.in))
			}
		})
	}
}

func ExampleNew() {
	var (
		cfg Config
		err error
	)

	// Initialize with default values
	if exampleBag, err = New(cfg); err != nil {
		log.Fatal(err)
	}
}

func ExampleNewFromTrainingSet() {
	var (
		t   TrainingSet
		err error
	)

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
	if exampleBag, err = NewFromTrainingSet(t); err != nil {
		log.Fatal(err)
	}
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
