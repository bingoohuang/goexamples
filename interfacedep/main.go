package main

import (
	"interfacedep/async"
	"interfacedep/runner"
	"interfacedep/textworker"
	"time"
)

func main() {
	worker := textworker.Worker{}
	asyncRunner := async.Runner{}

	runner.Run(asyncRunner, worker) // compiler error
	time.Sleep(1 * time.Second)
}
