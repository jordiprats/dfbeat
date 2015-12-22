package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/logp"

  "database/sql"
  _ "github.com/ziutek/mymysql/godrv"
)

const selector = "dfbeat"

type DFbeat struct {
	MbConfig ConfigSettings

	done chan uint

	urls   []*url.URL
	period time.Duration
}

// Config
func (mb *DFbeat) Config(b *beat.Beat) error {
	err := cfgfile.Read(&mb.MbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}

	var urlConfig []string
	if mb.MbConfig.Input.URLs != nil {
		urlConfig = mb.MbConfig.Input.URLs
	} else {
		urlConfig = []string{"mysql://nagios:secret@127.0.0.1"}
	}

	mb.urls = make([]*url.URL, len(urlConfig))
	for i := 0; i < len(urlConfig); i++ {
		u, err := url.Parse(urlConfig[i])
		if err != nil {
			logp.Err("Invalid mysql URL: %v", err)
			return err
		}
		mb.urls[i] = u
	}

	if mb.MbConfig.Input.Period != nil {
		mb.period = time.Duration(*mb.MbConfig.Input.Period) * time.Second
	} else {
		mb.period = 10 * time.Second
	}

	logp.Debug(selector, "Init DFbeat")
	logp.Debug(selector, "Watch %v", mb.urls)
	logp.Debug(selector, "Period %v", mb.period)

	return nil
}

// Setup
func (mb *DFbeat) Setup(b *beat.Beat) error {
	mb.done = make(chan uint)

  for _,u := range mb.urls {
    fmt.Println("Run!")
  }


	return nil
}

// Run
func (mb *DFbeat) Run(b *beat.Beat) error {
	logp.Debug(selector, "Run DFbeat")

	return nil
}

// Cleanup
func (mb *DFbeat) Cleanup(b *beat.Beat) error {
	return nil
}

// Stop
func (mb *DFbeat) Stop() {
	logp.Debug(selector, "Stop DFbeat")
	close(mb.done)
}
