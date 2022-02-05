// Package scanner defines the Scanner interface, and some reference implementations
package scanner

import (
	"context"
	"encoding/json"
	"net/netip"
	"time"
)

// Match is a result from a Scanner.Scan call.
type Match interface {
	Reason() string
}

// StringMatch is a simple implementation of Match
type StringMatch struct{ Message string }

// Reason implements Match
func (s *StringMatch) Reason() string { return s.Message }

// Scanner is any struct that can scan an IP and return information about it
type Scanner interface {
	Scan(context.Context, netip.Addr) (Match, error)
}

// Config is the most basic representation of a config for a scanner.
// It is expected that Scanner implementations unmarshal this into something a bit more sane.
type Config struct {
	Type    string
	Timeout time.Duration
	Args    json.RawMessage
}
