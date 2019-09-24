package pedantic

import (
	"regexp"
)

const (
	ipv4Address = `(?:` + numberByteDec + `\.` + numberByteDec + `\.` + numberByteDec + `\.` + numberByteDec + `)`
)

// IPv4 is a regexp that matches a valid ipv4 address.
var IPv4 = regexp.MustCompile(`^` + ipv4Address + `$`)
