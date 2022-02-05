package dummy

import (
	"context"
	"encoding/json"
	"net/netip"

	"awesome-dragon.science/go/gopm/scanner"
)

// Scanner is an always true
type Scanner struct {
	ReturnValue   bool
	DefaultReturn string
}

func New(conf scanner.Config) (*Scanner, error) {
	s := &Scanner{}
	if err := json.Unmarshal(conf.Args, s); err != nil {
		return nil, err
	}

	return s, nil
}

var _ scanner.Scanner = &Scanner{}

// Scan implements the scanner.Scanner interface.
func (s *Scanner) Scan(ctx context.Context, ip netip.Addr) (scanner.Match, error) {
	if !s.ReturnValue {
		return nil, nil
	}

	return &scanner.StringMatch{Message: s.DefaultReturn}, nil
}
