package gopm

import (
	"fmt"
	"sync"

	"awesome-dragon.science/go/gopm/scanner"
	"awesome-dragon.science/go/gopm/scanner/dummy"
)

/*
Scanners can do whatever, they return an error condition and a match object

Scanners accept a context, we will use that to enforce timeouts. Note that we will *not* be enforcing our own.
Its expected that a scanner deals with that
*/

type GOPM struct {
	mu       *sync.RWMutex // Protects Scanners and Pools
	scanners map[string]scanner.Scanner
	pools    map[string][]string
}

type GOPMConfig struct {
	DefaultPool string
	Scanners    map[string]scanner.Config
}

func New(config *GOPMConfig) (*GOPM, error) {
	g := &GOPM{
		scanners: make(map[string]scanner.Scanner),
		pools:    make(map[string][]string),
	}

	for name, scanner := range config.Scanners {
		if err := g.AddScanner(name, scanner); err != nil {
			return nil, fmt.Errorf("could not create scanner %s (%s): %w", name, scanner.Type, err)
		}
	}

	return g, nil
}

func (g *GOPM) AddScanner(name string, config scanner.Config) error {
	var s scanner.Scanner
	var err error
	switch config.Type {
	case "dummy":
		s, err = dummy.New(config)
	}

	if err != nil {
		return err
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.scanners[name] = s

	return nil
}
