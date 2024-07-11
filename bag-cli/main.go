package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/itsmontoya/bag"
)

func main() {
	var a app
	a.lines = onLine()
	a.done = onExit()
	flag.StringVar(&a.trainingSetLocation, "training", "", "./path/to/training.toml")
	flag.BoolVar(&a.interactive, "interactive", false, "true for interactive mode, false to exit immediately after first result")
	flag.Parse()

	var t bag.TrainingSet
	if _, err := toml.DecodeFile(a.trainingSetLocation, &t); err != nil {
		log.Fatalf("error loading training set: %v\n", err)
		return
	}

	a.interactivePrint("Training set loaded\n")
	a.b = bag.NewFromTrainingSet(t)
	a.interactivePrint("Model generated\n")
	a.interactivePrint("Interactive mode is active. Type your input and press Enter:\n")

	var done bool
	for !done {
		done = a.process()
	}
}
