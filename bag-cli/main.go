package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/itsmontoya/bag"
)

type app struct {
	trainingSetLocation string
	interactive         bool

	b *bag.Bag

	lines <-chan string
	done  <-chan os.Signal
}

func (a *app) interactivePrint(vals ...any) {
	if a.interactive {
		fmt.Print(vals...)
	}
}

func (a *app) process() (done bool) {
	select {
	case <-a.done:
		return true
	case line := <-a.lines:
		results := a.b.GetResults(line)
		a.interactivePrint("> ")
		fmt.Print(results.GetHighestProbability())
	}

	a.interactivePrint("\n")
	return !a.interactive
}

func main() {
	var a app
	flag.StringVar(&a.trainingSetLocation, "training", "", "./path/to/training.toml")
	flag.BoolVar(&a.interactive, "interactive", false, "true for interactive mode, false to exit immediately after first result")
	flag.Parse()

	a.lines = onLine()
	a.done = onExit()

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

func onLine() <-chan string {
	s := bufio.NewScanner(os.Stdin)
	lines := make(chan string)
	go func() {
		for {
			if !s.Scan() {
				return
			}

			input := string(s.Bytes())
			lines <- input
		}
	}()

	return lines
}

func onExit() <-chan os.Signal {
	out := make(chan os.Signal, 1)
	signal.Notify(out,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	return out
}
