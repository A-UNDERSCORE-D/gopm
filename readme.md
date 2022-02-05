# Go Open Proxy Monitor (GOPM)

## Design

GOPM is split into two parts. The IRC head and the scanning head. The two are connected very loosely, and thus
the scanning tools can be used without IRC, if you so choose.

### Scanner

Scanner is defined in `scanner/` and provides an interface to implement
