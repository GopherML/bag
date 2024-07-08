package bag

type TrainingSet struct {
	Config

	Samples []Sample `toml:"samples"`
}
