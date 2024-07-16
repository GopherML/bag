package main

import (
	"fmt"
	"os"

	"github.com/GopherML/bag"
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
