package bag

func New() *Bag {
	var b Bag
	b.ngramSize = 3
	b.labels = map[string]Vocabulary{}
	return &b
}

type Bag struct {
	ngramSize int
	labels    map[string]Vocabulary
}

func (b *Bag) GetSeniment(in string) (labels map[string]int) {
	ns := toNGrams(in, b.ngramSize)
	labels = make(map[string]int, len(b.labels))
	for _, n := range ns {
		for label, vocab := range b.labels {
			labels[label] += vocab[n.String()]
		}
	}

	return
}

func (b *Bag) Train(in, label string) {
	var target Vocabulary
	ns := toNGrams(in, b.ngramSize)
	target, ok := b.labels[label]
	if !ok {
		target = make(Vocabulary)
		b.labels[label] = target
	}

	for _, n := range ns {
		target[n.String()]++
	}
}
