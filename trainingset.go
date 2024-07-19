package bag

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

// makeTrainingSetFromFile will initialize a training set from a filepath
func makeTrainingSetFromFile(filepath string) (t TrainingSet, err error) {
	var f *os.File
	// Attempt to open file at given filepath
	if f, err = os.Open(filepath); err != nil {
		err = fmt.Errorf("error opening training set: %v", err)
		return
	}
	// Close file when function exits
	defer f.Close()

	// Initialize new YAML decoder and decode file as a training set
	err = yaml.NewDecoder(f).Decode(&t)
	return
}

// TrainingSet is used to train a bag of words (BoW) model
type TrainingSet struct {
	Config `yaml:"config"`

	Samples SamplesByLabel `yaml:"samples"`
}
