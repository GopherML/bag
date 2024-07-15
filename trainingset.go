package bag

type TrainingSet struct {
	Config `yaml:"config"`

	Samples SamplesByLabel `yaml:"samples"`
}
