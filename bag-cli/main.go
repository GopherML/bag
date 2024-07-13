package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/itsmontoya/bag"
)

func main() {
	var a app
	a.lines = onLine()
	a.done = onExit()
	flag.StringVar(&a.trainingSetLocation, "training", "", "./path/to/training.yaml")
	flag.BoolVar(&a.interactive, "interactive", false, "true for interactive mode, false to exit immediately after first result")
	flag.Parse()

	var (
		f   *os.File
		err error
	)

	if f, err = os.Open(a.trainingSetLocation); err != nil {
		log.Fatalf("error opening training set: %v\n", err)
		return
	}

	var t bag.TrainingSet
	if err = yaml.NewDecoder(f).Decode(&t); err != nil {
		log.Fatalf("error loading training set: %v\n", err)
		return
	}

	fmt.Printf("TS: %+v\n", t)

	a.interactivePrint("Training set loaded\n")
	if a.b, err = bag.NewFromTrainingSet(t); err != nil {
		log.Fatalf("error initializing from training set: %v\n", err)
		return
	}

	a.interactivePrint("Model generated\n")
	a.interactivePrint("Interactive mode is active. Type your input and press Enter:\n")

	var done bool
	for !done {
		done = a.process()
	}
}
