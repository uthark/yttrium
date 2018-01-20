package main

import (
	"math/rand"
	"runtime"
	"time"

	"bitbucket.org/uthark/yttrium/cmd"
)

func main() {
	// Configure GOMAXPROCS.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Set seed for rand.
	rand.Seed(time.Now().Unix())

	// Start command.
	cmd.Execute()
}
