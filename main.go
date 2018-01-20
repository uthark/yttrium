package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"bitbucket.org/uthark/yttrium/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	// Configure GOMAXPROCS.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Set seed for rand.
	rand.Seed(time.Now().Unix())

	// Start command.
	cmd.Execute()
}
