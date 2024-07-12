package main

import (
	"bufio"
	"os"
	"os/signal"
	"syscall"
)

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
