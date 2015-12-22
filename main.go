
package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
)

var Version = "0.1.0"
var Name = "dfbeat"

func main() {
	mb := &DFbeat{}

	b := beat.NewBeat(Name, Version, mb)

	b.CommandLineSetup()

	b.LoadConfig()
	err := mb.Config(b)
	if err != nil {
		logp.Critical("Config error: %v", err)
		os.Exit(1)
	}

	b.Run()
}
