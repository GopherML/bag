package bag

type TrainingSet struct {
	Config `yaml:"config"`

	Samples SamplesByLabel `yaml:"samples"`
}

type SamplesByLabel map[string]Samples

type Samples []string
