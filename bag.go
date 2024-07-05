package bag

func New() *Bag {
	var b Bag
	b.labels = map[string]Vocabulary{}
	return &b
}

type Bag struct {
	labels map[string]Vocabulary
}

func (b *Bag) GetSeniment(in string) (labels map[string]int) {
	ts := toTrigrams(in)
	labels = make(map[string]int, len(b.labels))
	for _, t := range ts {
		for label, vocab := range b.labels {
			labels[label] += vocab[t]
		}
	}

	return
}

func (b *Bag) Train(in, label string) {
	var target Vocabulary
	ts := toTrigrams(in)
	target, ok := b.labels[label]
	if !ok {
		target = make(Vocabulary)
		b.labels[label] = target
	}

	for _, t := range ts {
		target[t]++
	}
}
