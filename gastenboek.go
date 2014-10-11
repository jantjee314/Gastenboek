package main

import (
	"os"
	"os/signal"
)

// IT'S A TRAP! (This catches all signals so the program can't be interrupted)
func signalCatcher() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for _ = range c {
		//fmt.Printf("Signal received: %v", sig)
	}
}

// This is where it all starts
func main() {
	go signalCatcher()
	showMenu()
}
