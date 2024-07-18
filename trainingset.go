package bag

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

func makeTrainingSetFromFile(filepath string) (t TrainingSet, err error) {
	var f *os.File
	if f, err = os.Open(filepath); err != nil {
		err = fmt.Errorf("error opening training set: %v", err)
		return
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&t)
	return
}

type TrainingSet struct {
	Config `yaml:"config"`

	Samples SamplesByLabel `yaml:"samples"`
}
